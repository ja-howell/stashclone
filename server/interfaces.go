package server

import "github.com/ja-howell/stashclone/models"

type Database interface {
	GetStashItem(id int) (models.StashItem, error)
	ListStashItems() ([]models.StashItem, error)
	DeleteStashItem(id int) error
	CreateStashItem(si models.StashItem) error
	UpdateStashItem(id int, si models.StashItem) error
}
