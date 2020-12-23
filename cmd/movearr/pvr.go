package main

import (
	"github.com/l3uddz/movearr"
	"github.com/l3uddz/movearr/radarr"
	"github.com/l3uddz/movearr/sonarr"
	"github.com/pkg/errors"
	"strings"
)

type PVR interface {
	// client
	Type() string

	// api
	Available() error
	Move([]uint64) error

	// datastore
	GetItemsWithIncorrectIds() ([]movearr.MediaItem, error)
	GetItemsWithIncorrectYears() ([]movearr.MediaItem, error)
	GetItemsWithMissingIds() ([]movearr.MediaItem, error)
}

func NewPVR(c *config, pvr string) (PVR, error) {
	switch strings.ToLower(pvr) {
	case "sonarr":
		return sonarr.New(c.Sonarr)
	case "radarr":
		return radarr.New(c.Radarr)
	}

	return nil, errors.New("unknown pvr")
}
