package database

import (
	"github.com/ja-howell/stashclone/models"
)

type Database struct {
	rows []models.StashItem
}

func (db *Database) GetStashItem(id int) (models.StashItem, error) {
	si := models.StashItem{
		Name: "foo",
		ID:   1,
	}

	return si, nil
}
