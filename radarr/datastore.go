package radarr

import (
	"database/sql"
	"fmt"
	"github.com/l3uddz/movearr"
	"net/url"
	// database driver
	_ "github.com/mattn/go-sqlite3"
)

func newDatastore(path string) (*datastore, error) {
	q := url.Values{}
	q.Set("mode", "ro")

	db, err := sql.Open("sqlite3", movearr.DSN(path, q))
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	return &datastore{db: db}, nil
}

type datastore struct {
	db *sql.DB
}

type MediaItem struct {
	Id     uint64
	Title  string
	Year   uint64
	ImdbId string
	TmdbId uint64
	Path   string
}

func (d *datastore) GetItemsWithIncorrectIds() ([]MediaItem, error) {
	rows, err := d.db.Query(sqlSelectFixIds)
	if err != nil {
		return nil, fmt.Errorf("select media items: %v", err)
	}

	defer rows.Close()

	mediaItems := make([]MediaItem, 0)
	for rows.Next() {
		m := new(struct {
			Id     *uint64
			Title  *string
			Year   *uint64
			ImdbId *string
			TmdbId *uint64
			Path   *string
		})

		if err := rows.Scan(&m.Id, &m.Title, &m.Year, &m.ImdbId, &m.TmdbId, &m.Path); err != nil {
			return nil, fmt.Errorf("scan media item row: %w", err)
		}

		if m.Id == nil || m.Title == nil || m.TmdbId == nil || &m.Path == nil {
			return nil, fmt.Errorf("invalid media item row: %w: %v", err, m)
		}

		mediaItems = append(mediaItems, MediaItem{
			Id:     *m.Id,
			Title:  *m.Title,
			Year:   movearr.Uint64OrDefault(m.Year, 0),
			ImdbId: movearr.StringOrDefault(m.ImdbId, ""),
			TmdbId: *m.TmdbId,
			Path:   *m.Path,
		})
	}

	return mediaItems, nil
}

func (d *datastore) GetItemsWithIncorrectYears() ([]MediaItem, error) {
	rows, err := d.db.Query(sqlSelectFixYears)
	if err != nil {
		return nil, fmt.Errorf("select media items: %v", err)
	}

	defer rows.Close()

	mediaItems := make([]MediaItem, 0)
	for rows.Next() {
		m := new(struct {
			Id     *uint64
			Title  *string
			Year   *uint64
			ImdbId *string
			TmdbId *uint64
			Path   *string
		})

		if err := rows.Scan(&m.Id, &m.Title, &m.Year, &m.ImdbId, &m.TmdbId, &m.Path); err != nil {
			return nil, fmt.Errorf("scan media item row: %w", err)
		}

		if m.Id == nil || m.Title == nil || m.TmdbId == nil || &m.Path == nil {
			return nil, fmt.Errorf("invalid media item row: %w: %v", err, m)
		}

		mediaItems = append(mediaItems, MediaItem{
			Id:     *m.Id,
			Title:  *m.Title,
			Year:   movearr.Uint64OrDefault(m.Year, 0),
			ImdbId: movearr.StringOrDefault(m.ImdbId, ""),
			TmdbId: *m.TmdbId,
			Path:   *m.Path,
		})
	}

	return mediaItems, nil
}

func (d *datastore) GetItemsMissingIds() ([]MediaItem, error) {
	rows, err := d.db.Query(sqlSelectMissingIds)
	if err != nil {
		return nil, fmt.Errorf("select media items: %v", err)
	}

	defer rows.Close()

	mediaItems := make([]MediaItem, 0)
	for rows.Next() {
		m := new(struct {
			Id     *uint64
			Title  *string
			Year   *uint64
			ImdbId *string
			TmdbId *uint64
			Path   *string
		})

		if err := rows.Scan(&m.Id, &m.Title, &m.Year, &m.ImdbId, &m.TmdbId, &m.Path); err != nil {
			return nil, fmt.Errorf("scan media item row: %w", err)
		}

		if m.Id == nil || m.Title == nil || m.TmdbId == nil || &m.Path == nil {
			return nil, fmt.Errorf("invalid media item row: %w: %v", err, m)
		}

		mediaItems = append(mediaItems, MediaItem{
			Id:     *m.Id,
			Title:  *m.Title,
			Year:   movearr.Uint64OrDefault(m.Year, 0),
			ImdbId: movearr.StringOrDefault(m.ImdbId, ""),
			TmdbId: *m.TmdbId,
			Path:   *m.Path,
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
              , M.ImdbId
              , M.TmdbId
              , M.Path
FROM Movies M
         JOIN MovieFiles MF ON MF.MovieId = M.Id

WHERE M.Path IS NOT NULL
  AND (
        (M.ImdbId IS NOT NULL AND M.Path LIKE '%imdb:%' AND M.Path NOT LIKE '%imdb:' || M.ImdbId || '%')
        OR
        (M.Path LIKE '%tmdb:%' AND M.Path NOT LIKE '%tmdb:' || M.TmdbId || '%')
    )
ORDER BY M.Id ASC
`
	sqlSelectMissingIds = `
SELECT DISTINCT M.Id
              , M.Title
              , M.Year
              , M.ImdbId
              , M.TmdbId
              , M.Path
FROM Movies M
         JOIN MovieFiles MF ON MF.MovieId = M.Id
WHERE M.Path IS NOT NULL
  AND (
        (M.ImdbId IS NOT NULL AND M.Path NOT LIKE '%imdb:' || M.ImdbId || '%')
        OR
        M.Path NOT LIKE '%tmdb:' || M.TmdbId || '%'
    )
ORDER BY M.Id ASC
`
	sqlSelectFixYears = `
SELECT DISTINCT M.Id
              , M.Title
              , M.Year
              , M.ImdbId
              , M.TmdbId
              , M.Path
FROM Movies M
         JOIN MovieFiles MF ON MF.MovieId = M.Id
WHERE M.Path IS NOT NULL
  AND M.Year > 0
  AND M.Path NOT LIKE '%(' || M.Year || ')%'
ORDER BY M.Id ASC
`
)
