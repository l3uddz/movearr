package sonarr

import (
	"database/sql"
	"fmt"
	"github.com/l3uddz/movearr"
	"net/url"
	// database driver
	_ "github.com/mattn/go-sqlite3"
)

func newDatastore(path string, metadataSeparator string) (*datastore, error) {
	q := url.Values{}
	q.Set("mode", "ro")

	db, err := sql.Open("sqlite3", movearr.DSN(path, q))
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	return &datastore{db: db, metadataSeparator: metadataSeparator}, nil
}

type datastore struct {
	db                *sql.DB
	metadataSeparator string
}

func (d *datastore) GetItemsWithIncorrectIds() ([]movearr.MediaItem, error) {
	rows, err := d.db.Query(sqlSelectFixIds, d.metadataSeparator)
	if err != nil {
		return nil, fmt.Errorf("select media items: %v", err)
	}

	defer rows.Close()

	mediaItems := make([]movearr.MediaItem, 0)
	for rows.Next() {
		m := new(struct {
			Id    *uint64
			Title *string
			Year  *uint64
			Path  *string
		})

		if err := rows.Scan(&m.Id, &m.Title, &m.Year, &m.Path); err != nil {
			return nil, fmt.Errorf("scan media item row: %w", err)
		}

		if m.Id == nil || m.Title == nil || &m.Path == nil {
			return nil, fmt.Errorf("invalid media item row: %w: %v", err, m)
		}

		mediaItems = append(mediaItems, movearr.MediaItem{
			Id:    *m.Id,
			Title: *m.Title,
			Year:  movearr.Uint64OrDefault(m.Year, 0),
			Path:  *m.Path,
		})
	}

	return mediaItems, nil
}

func (d *datastore) GetItemsWithIncorrectYears() ([]movearr.MediaItem, error) {
	rows, err := d.db.Query(sqlSelectFixYears)
	if err != nil {
		return nil, fmt.Errorf("select media items: %v", err)
	}

	defer rows.Close()

	mediaItems := make([]movearr.MediaItem, 0)
	for rows.Next() {
		m := new(struct {
			Id    *uint64
			Title *string
			Year  *uint64
			Path  *string
		})

		if err := rows.Scan(&m.Id, &m.Title, &m.Year, &m.Path); err != nil {
			return nil, fmt.Errorf("scan media item row: %w", err)
		}

		if m.Id == nil || m.Title == nil || &m.Path == nil {
			return nil, fmt.Errorf("invalid media item row: %w: %v", err, m)
		}

		mediaItems = append(mediaItems, movearr.MediaItem{
			Id:    *m.Id,
			Title: *m.Title,
			Year:  movearr.Uint64OrDefault(m.Year, 0),
			Path:  *m.Path,
		})
	}

	return mediaItems, nil
}

func (d *datastore) GetItemsMissingIds() ([]movearr.MediaItem, error) {
	rows, err := d.db.Query(sqlSelectMissingIds, d.metadataSeparator)
	if err != nil {
		return nil, fmt.Errorf("select media items: %v", err)
	}

	defer rows.Close()

	mediaItems := make([]movearr.MediaItem, 0)
	for rows.Next() {
		m := new(struct {
			Id    *uint64
			Title *string
			Year  *uint64
			Path  *string
		})

		if err := rows.Scan(&m.Id, &m.Title, &m.Year, &m.Path); err != nil {
			return nil, fmt.Errorf("scan media item row: %w", err)
		}

		if m.Id == nil || m.Title == nil || &m.Path == nil {
			return nil, fmt.Errorf("invalid media item row: %w: %v", err, m)
		}

		mediaItems = append(mediaItems, movearr.MediaItem{
			Id:    *m.Id,
			Title: *m.Title,
			Year:  movearr.Uint64OrDefault(m.Year, 0),
			Path:  *m.Path,
		})
	}

	return mediaItems, nil
}

//goland:noinspection ALL
const (
	sqlSelectFixIds = `
SELECT DISTINCT M.Id
              , M.Title
              , M.Year
              , M.Path
FROM Series M
WHERE M.Path IS NOT NULL
  AND (M.TvdbId > 0 AND M.Path LIKE '%tvdb' || $1 || '%' AND M.Path NOT LIKE '%tvdb' || $1 || M.TvdbId || '%')
ORDER BY M.Id ASC
`
	sqlSelectMissingIds = `
SELECT DISTINCT M.Id
              , M.Title
              , M.Year
              , M.Path
FROM Series M
WHERE M.Path IS NOT NULL
  AND (M.TvdbId > 0 AND M.Path NOT LIKE '%tvdb' || $1 || M.TvdbId || '%')
ORDER BY M.Id ASC
`
	sqlSelectFixYears = `
SELECT DISTINCT M.Id
              , M.Title
              , M.Year
              , M.Path
FROM Series M
WHERE M.Path IS NOT NULL
  AND M.Year > 0
  AND M.Path NOT LIKE '%(' || M.Year || ')%'
ORDER BY M.Id ASC
`
)
