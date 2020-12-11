package radarr

import (
	"fmt"
)

func (c *Client) GetItemsWithIncorrectIds() ([]MediaItem, error) {
	items, err := c.store.GetItemsWithIncorrectIds()
	if err != nil {
		return nil, fmt.Errorf("retrieve items with incorrect ids: %w", err)
	}

	return items, nil
}

func (c *Client) GetItemsWithMissingIds() ([]MediaItem, error) {
	items, err := c.store.GetItemsMissingIds()
	if err != nil {
		return nil, fmt.Errorf("retrieve items with missing ids: %w", err)
	}

	return items, nil
}
