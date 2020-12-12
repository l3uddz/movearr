package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/l3uddz/movearr/radarr"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"path/filepath"
)

type config struct {
	Radarr radarr.Config `yaml:"radarr"`
}

var (
	// Release variables
	Version   string
	Timestamp string
	GitCommit string

	// CLI
	cli struct {
		globals

		// flags
		Config    string `type:"path" default:"${config_file}" env:"movearr_CONFIG" help:"Config file path"`
		Log       string `type:"path" default:"${log_file}" env:"movearr_LOG" help:"Log file path"`
		Verbosity int    `type:"counter" default:"0" short:"v" env:"movearr_VERBOSITY" help:"Log level verbosity"`

		// commands
		FixIds struct {
			DryRun bool `type:"bool" default:"0" help:"Dry run mode"`
			Limit  int  `required:"0" type:"int" default:"0" help:"Maximum number of items"`
		} `cmd help:"Move items with incorrect ids."`
		FixYears struct {
			DryRun bool `type:"bool" default:"0" help:"Dry run mode"`
			Limit  int  `required:"0" type:"int" default:"0" help:"Maximum number of items"`
		} `cmd help:"Move items with incorrect years."`
		MissingIds struct {
			DryRun bool `type:"bool" default:"0" help:"Dry run mode"`
			Limit  int  `required:"0" type:"int" default:"0" help:"Maximum number of items"`
		} `cmd help:"Move items missing ids."`
	}
)

type globals struct {
	Version versionFlag `name:"version" help:"Print version information and quit"`
}

type versionFlag string

func (v versionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v versionFlag) IsBool() bool                         { return true }
func (v versionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

func main() {
	// parse cli
	ctx := kong.Parse(&cli,
		kong.Name("movearr"),
		kong.Description("Various plex tools"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Summary: true,
			Compact: true,
		}),
		kong.Vars{
			"version":     fmt.Sprintf("%s (%s@%s)", Version, GitCommit, Timestamp),
			"config_file": filepath.Join(defaultConfigPath(), "config.yml"),
			"log_file":    filepath.Join(defaultConfigPath(), "activity.log"),
		},
	)

	if err := ctx.Validate(); err != nil {
		fmt.Println("Failed parsing cli:", err)
		return
	}

	// logger
	logger := log.Output(io.MultiWriter(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}, zerolog.ConsoleWriter{
		Out: &lumberjack.Logger{
			Filename:   cli.Log,
			MaxSize:    5,
			MaxAge:     14,
			MaxBackups: 5,
		},
		NoColor: true,
	}))

	switch {
	case cli.Verbosity == 1:
		log.Logger = logger.Level(zerolog.DebugLevel)
	case cli.Verbosity > 1:
		log.Logger = logger.Level(zerolog.TraceLevel)
	default:
		log.Logger = logger.Level(zerolog.InfoLevel)
	}

	// config
	file, err := os.Open(cli.Config)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed opening config")
	}
	defer file.Close()

	cfg := config{}
	decoder := yaml.NewDecoder(file)
	decoder.SetStrict(true)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed decoding config")
	}

	switch {
	case cfg.Radarr.Database == "":
		log.Fatal().Msg("You must set a radarr database path in your configuration")
	}

	// plex
	p, err := radarr.New(cfg.Radarr)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed initialising radarr")
	}

	if err := p.Available(); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed validating radarr availability")
	}

	// switch to appropriate command
	switch ctx.Command() {
	case "fix-ids":
		fixIds(p, cli.FixIds.DryRun, cli.FixIds.Limit)
		return
	case "fix-years":
		fixYears(p, cli.FixYears.DryRun, cli.FixYears.Limit)
		return
	case "missing-ids":
		missingIds(p, cli.MissingIds.DryRun, cli.MissingIds.Limit)
		return
	}
}
