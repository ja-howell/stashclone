package server

import "github.com/ja-howell/stashclone/models"

type Database interface {
	// GetStashItem retrieves a single stash item by its ID.
	GetStashItem(id int) (models.StashItem, error)
	// ListStashItems retrieves all stash items.
	// TODO: Add optional parameters for filtering
	ListStashItems() ([]models.StashItem, error)
	// DeleteStashItem deletes a stash item by its ID.
	DeleteStashItem(id int) error
	// CreateStashItem creates a new stash item.
	CreateStashItem(si models.StashItem) error
	// UpdateStashItem updates an existing stash item by its ID.
	UpdateStashItem(id int, si models.StashItem) error
}
