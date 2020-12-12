package main

import (
	"github.com/l3uddz/movearr/radarr"
	"github.com/rs/zerolog/log"
)

func fixYears(pc *radarr.Client, dryRun bool, limit int) {
	// retrieve items with incorrect years
	items, err := pc.GetItemsWithIncorrectYears()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed retrieving items")
		return
	}

	count := len(items)
	if count == 0 {
		log.Info().Msg("There are no items with incorrect years")
		return
	}

	log.Info().
		Int("count", count).
		Msg("Found items with incorrect years")

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
		log.Warn().
			Int("count", len(idsToMove)).
			Msg("Dry run enabled, not sending move request")
		return
	}

	// move items
	if err := pc.Move(idsToMove); err != nil {
		log.Error().
			Err(err).
			Int("count", len(idsToMove)).
			Msg("Failed moving items...")
		return
	}

	log.Info().
		Int("count", len(idsToMove)).
		Msg("Move request sent")
	return
}