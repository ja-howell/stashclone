// Satisfy the Database Interface
package database

import (
	"database/sql"
	"fmt"

	"github.com/ja-howell/stashclone/models"
)

type MySQL struct {
	wrapped *sql.DB
}

func NewMySQL(mysqldb *sql.DB) *MySQL {
	db := &MySQL{
		wrapped: mysqldb,
	}

	return db
}

func (db *MySQL) GetStashItem(id int) (models.StashItem, error) {
	var si models.StashItem
	row := db.wrapped.QueryRow("SELECT * FROM stashitems WHERE id = ?", id)
	err := row.Scan(&si.ID, &si.Name, &si.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return si, fmt.Errorf("item not found: %w", err)
		}
		return si, fmt.Errorf("read failure: %w", err)
	}
	return si, nil
}

func (db *MySQL) GetAllStashItems() (map[int]models.StashItem, error) {
	return nil, nil
}

func (db *MySQL) DeleteStashItem(id int) error {
	return nil
}

func (db *MySQL) CreateStashItem(si models.StashItem) error {
	return nil
}

func (db *MySQL) UpdateStashItem(id int, si models.StashItem) error {
	return nil
}
