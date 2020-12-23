package movearr

import (
	"github.com/pkg/errors"
)

var (
	// ErrRadarrUnavailable may occur when a radarr api cannot be validated
	ErrRadarrUnavailable = errors.New("radarr unavailable")

	// ErrFatal indicates a severe problem related to development.
	ErrFatal = errors.New("fatal development related error")
)

type MediaItem struct {
	Id     uint64
	Title  string
	Year   uint64
	ImdbId string
	TmdbId uint64
	TvdbId uint64
	Path   string
}

func Uint64OrDefault(currentValue *uint64, defaultValue uint64) uint64 {
	if currentValue == nil {
		return defaultValue
	}

	return *currentValue
}

func StringOrDefault(currentValue *string, defaultValue string) string {
	if currentValue == nil || *currentValue == "" {
		return defaultValue
	}

	return *currentValue
}
