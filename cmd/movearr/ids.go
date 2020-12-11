package main

import (
	"github.com/l3uddz/movearr/radarr"
	"github.com/rs/zerolog/log"
)

func fixIds(pc *radarr.Client, dryRun bool, limit int) {
	// retrieve items with incorrect ids
	items, err := pc.GetItemsWithIncorrectIds()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed retrieving items")
		return
	}

	count := len(items)
	if count == 0 {
		log.Info().Msg("There are no items with incorrect ids")
		return
	}

	log.Info().
		Int("count", count).
		Msg("Found items with incorrect ids")

	// process items
	idsToMove := make([]uint64, 0)
	for pos, item := range items {
		log.Info().
			Str("title", item.Title).
			Uint64("year", item.Year).
			Str("path", item.Path).
			Msg("Batching item")

		idsToMove = append(idsToMove, item.Id)
		if limit == 0 {
			continue
		} else if (pos + 1) >= limit {
			break
		}
	}

	if dryRun {
		return
	}

	// move items
	log.Info().
		Int("count", len(idsToMove)).
		Msg("Moving items")

	if err := pc.Move(idsToMove); err != nil {
		log.Error().
			Err(err).
			Msg("Failed moving items...")
		return
	}

	log.Info().Msg("Move request sent")
}

func missingIds(pc *radarr.Client, dryRun bool, limit int) {
	// retrieve items with missing ids
	items, err := pc.GetItemsWithMissingIds()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed retrieving items")
		return
	}

	count := len(items)
	if count == 0 {
		log.Info().Msg("There are no items with missing ids")
		return
	}

	log.Info().
		Int("count", count).
		Msg("Found items with missing ids")

	// process items
	idsToMove := make([]uint64, 0)
	for pos, item := range items {
		log.Info().
			Str("title", item.Title).
			Uint64("year", item.Year).
			Str("path", item.Path).
			Msg("Batching item")

		idsToMove = append(idsToMove, item.Id)
		if limit == 0 {
			continue
		} else if (pos + 1) >= limit {
			break
		}
	}

	if dryRun {
		return
	}

	// move items
	log.Info().
		Int("count", len(idsToMove)).
		Msg("Moving items")

	if err := pc.Move(idsToMove); err != nil {
		log.Error().
			Err(err).
			Msg("Failed moving items...")
		return
	}

	log.Info().Msg("Move request sent")
}
