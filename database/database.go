package database

import (
	"fmt"

	"github.com/ja-howell/stashclone/models"
)

type Database struct {
	nextIndex int
	rows      map[int]models.StashItem
}

func New(data map[int]models.StashItem) *Database {
	db := &Database{
		nextIndex: len(data),
		rows:      data,
	}

	return db
}

func (db *Database) GetStashItem(id int) (models.StashItem, error) {
	if _, ok := db.rows[id]; !ok {
		err := fmt.Errorf("invalid ID: %v", id)
		return models.StashItem{}, err
	}

	return db.rows[id], nil
}

func (db *Database) GetAllStashItems() (map[int]models.StashItem, error) {
	return db.rows, nil
}

// TODO: func DeleteStashItem
func (db *Database) DeleteStashItem(id int) error {
	if _, ok := db.rows[id]; !ok {
		err := fmt.Errorf("invalid ID: %v", id)
		return err
	}
	delete(db.rows, id)
	return nil
}

func (db *Database) CreateStashItem(si models.StashItem) error {
	si.ID = db.nextIndex
	db.nextIndex++
	db.rows[si.ID] = si
	return nil
}

func (db *Database) UpdateStashItem(id int, si models.StashItem) error {
	if _, ok := db.rows[id]; !ok {
		err := fmt.Errorf("invalid ID: %v", id)
		return err
	}

	db.rows[id] = si
	return nil
}
