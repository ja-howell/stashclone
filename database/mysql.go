// Satisfy the Database Interface
package database

import (
	"context"
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

func (db *MySQL) ListStashItems() ([]models.StashItem, error) {
	sis := []models.StashItem{}
	rows, err := db.wrapped.QueryContext(context.TODO(), "SELECT * FROM stashitems")
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	for rows.Next() {
		var si models.StashItem
		err := rows.Scan(&si.ID, &si.Name, &si.Type)
		if err != nil {
			return nil, fmt.Errorf("read failure: %w", err)
		}
		sis = append(sis, si)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to get next row: %w", err)
	}
	return sis, nil
}

// Return just the name and id for the landing page.
func (db *MySQL) GetStashItemsFrontPage() ([]models.StashItem, error) {
	// This method is not implemented yet.
	return nil, fmt.Errorf("GetStashItemsFrontPage not implemented")
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
