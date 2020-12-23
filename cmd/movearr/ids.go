package main

import (
	"github.com/rs/zerolog/log"
)

func fixIds(pc PVR, dryRun bool, limit int) {
	// set logger
	l := log.With().
		Str("pvr", pc.Type()).
		Logger()

	// retrieve items with incorrect ids
	items, err := pc.GetItemsWithIncorrectIds()
	if err != nil {
		l.Error().
			Err(err).
			Msg("Failed retrieving items")
		return
	}

	count := len(items)
	if count == 0 {
		l.Info().Msg("There are no items with incorrect ids")
		return
	}

	l.Info().
		Int("count", count).
		Msg("Found items with incorrect ids")

	// process items
	idsToMove := make([]uint64, 0)
	for pos, item := range items {
		l.Info().
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
		l.Warn().
			Int("count", len(idsToMove)).
			Msg("Dry run enabled, not sending move request")
		return
	}

	// move items
	if err := pc.Move(idsToMove); err != nil {
		l.Error().
			Err(err).
			Int("count", len(idsToMove)).
			Msg("Failed moving items...")
		return
	}

	l.Info().
		Int("count", len(idsToMove)).
		Msg("Move request sent")
	return
}

func missingIds(pc PVR, dryRun bool, limit int) {
	// set logger
	l := log.With().
		Str("pvr", pc.Type()).
		Logger()

	// retrieve items with missing ids
	items, err := pc.GetItemsWithMissingIds()
	if err != nil {
		l.Error().
			Err(err).
			Msg("Failed retrieving items")
		return
	}

	count := len(items)
	if count == 0 {
		l.Info().Msg("There are no items with missing ids")
		return
	}

	l.Info().
		Int("count", count).
		Msg("Found items with missing ids")

	// process items
	idsToMove := make([]uint64, 0)
	for pos, item := range items {
		l.Info().
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
		l.Warn().
			Int("count", len(idsToMove)).
			Msg("Dry run enabled, not sending move request")
		return
	}

	// move items
	if err := pc.Move(idsToMove); err != nil {
		l.Error().
			Err(err).
			Int("count", len(idsToMove)).
			Msg("Failed moving items...")
		return
	}

	l.Info().
		Int("count", len(idsToMove)).
		Msg("Move request sent")
	return
}
