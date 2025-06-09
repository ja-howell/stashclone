package database

import (
	"fmt"

	"github.com/ja-howell/stashclone/models"
)

type Database struct {
	rows []models.StashItem
}

func New(data []models.StashItem) *Database {
	db := &Database{rows: data}

	return db
}

func (db *Database) GetStashItem(id int) (models.StashItem, error) {
	if id < 0 || id >= len(db.rows) {
		err := fmt.Errorf("invalid ID: %v", id)
		return models.StashItem{}, err
	}
	return db.rows[id], nil
}

func (db *Database) CreateStashItem(si models.StashItem) error {
	si.ID = len(db.rows)
	db.rows = append(db.rows, si)
	return nil
}

func (db *Database) UpdateStashItem(id int, si models.StashItem) error {
	if id < 0 || id >= len(db.rows) {
		err := fmt.Errorf("invalid ID: %v", si.ID)
		return err
	}
	db.rows[id].Name = si.Name
	return nil
}
