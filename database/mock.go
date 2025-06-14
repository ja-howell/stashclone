package database

import (
	"fmt"

	"github.com/ja-howell/stashclone/models"
)

type Mock struct {
	nextIndex int
	rows      map[int]models.StashItem
}

func NewMock(data map[int]models.StashItem) *Mock {
	db := &Mock{
		nextIndex: len(data),
		rows:      data,
	}

	return db
}

func (db *Mock) GetStashItem(id int) (models.StashItem, error) {
	if _, ok := db.rows[id]; !ok {
		err := fmt.Errorf("invalid ID: %v", id)
		return models.StashItem{}, err
	}

	return db.rows[id], nil
}

func (db *Mock) ListStashItems() ([]models.StashItem, error) {
	sis := []models.StashItem{}
	for _, si := range db.rows {
		sis = append(sis, si)
	}
	return sis, nil
}

// TODO: func DeleteStashItem
func (db *Mock) DeleteStashItem(id int) error {
	if _, ok := db.rows[id]; !ok {
		err := fmt.Errorf("invalid ID: %v", id)
		return err
	}
	delete(db.rows, id)
	return nil
}

func (db *Mock) CreateStashItem(si models.StashItem) error {
	si.ID = db.nextIndex
	db.nextIndex++
	db.rows[si.ID] = si
	return nil
}

func (db *Mock) UpdateStashItem(id int, si models.StashItem) error {
	if _, ok := db.rows[id]; !ok {
		err := fmt.Errorf("invalid ID: %v", id)
		return err
	}

	db.rows[id] = si
	return nil
}
