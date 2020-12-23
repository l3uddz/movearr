package sonarr

import (
	"fmt"
	"github.com/l3uddz/movearr"
)

func (c *Client) GetItemsWithIncorrectIds() ([]movearr.MediaItem, error) {
	items, err := c.store.GetItemsWithIncorrectIds()
	if err != nil {
		return nil, fmt.Errorf("retrieve items with incorrect ids: %w", err)
	}

	return items, nil
}

func (c *Client) GetItemsWithIncorrectYears() ([]movearr.MediaItem, error) {
	items, err := c.store.GetItemsWithIncorrectYears()
	if err != nil {
		return nil, fmt.Errorf("retrieve items with incorrect years: %w", err)
	}

	return items, nil
}

func (c *Client) GetItemsWithMissingIds() ([]movearr.MediaItem, error) {
	items, err := c.store.GetItemsMissingIds()
	if err != nil {
		return nil, fmt.Errorf("retrieve items with missing ids: %w", err)
	}

	return items, nil
}
