package radarr

import (
	"github.com/l3uddz/movearr"
	"github.com/rs/zerolog"
)

type Config struct {
	Database   string `yaml:"database"`
	URL        string `yaml:"url"`
	ApiKey     string `yaml:"api_key"`
	RootFolder string `yaml:"root_folder"`

	Verbosity string `yaml:"verbosity"`
}

type Client struct {
	url        string
	apiKey     string
	rootFolder string

	log   zerolog.Logger
	store *datastore
}

func New(c Config) (*Client, error) {
	store, err := newDatastore(c.Database)
	if err != nil {
		return nil, err
	}

	l := movearr.GetLogger(c.Verbosity).With().Logger()

	return &Client{
		url:        c.URL,
		apiKey:     c.ApiKey,
		rootFolder: c.RootFolder,

		log:   l,
		store: store,
	}, nil
}
