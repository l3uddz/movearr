package radarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/l3uddz/movearr"
	"net/http"
)

func (c *Client) Available() error {
	// create request
	req, err := http.NewRequest("GET", movearr.JoinURL(c.url, "api", "v3", "system", "status"),
		nil)
	if err != nil {
		return fmt.Errorf("%v: %w", err, movearr.ErrFatal)
	}

	// set headers
	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Accept", "application/json")

	// send request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not check Radarr availability: %v: %w",
			err, movearr.ErrRadarrUnavailable)
	}

	defer res.Body.Close()

	// validate response
	if res.StatusCode != 200 {
		return fmt.Errorf("could not check Radarr availability: %v: %w",
			res.StatusCode, movearr.ErrRadarrUnavailable)
	}

	return nil
}

func (c *Client) Move(movieIds []uint64) error {
	// build payload
	payload := new(struct {
		MovieIds       []uint64 `json:"movieIds"`
		RootFolderPath string   `json:"rootFolderPath"`
		MoveFiles      bool     `json:"moveFiles"`
	})

	payload.MoveFiles = true
	payload.RootFolderPath = c.rootFolder
	payload.MovieIds = append(payload.MovieIds, movieIds...)

	// encode payload
	js, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("could not encode payload: %w", err)
	}

	// create request
	req, err := http.NewRequest("PUT", movearr.JoinURL(c.url, "api", "v3", "movie", "editor"),
		bytes.NewBuffer(js))
	if err != nil {
		return fmt.Errorf("%v: %w", err, movearr.ErrFatal)
	}

	// set headers
	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// send request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not move Radarr moviesv: %w", err)
	}

	defer res.Body.Close()

	// validate response
	if res.StatusCode != 202 {
		return fmt.Errorf("could not move Radarr movies %v: %w", res.StatusCode, movearr.ErrFatal)
	}

	return nil
}
