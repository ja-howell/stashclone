package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/ja-howell/stashclone/server"
)

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = "user"
	cfg.Passwd = "test"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "stash"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	// si := models.StashItem{
	// 	Name: "foo",
	// 	ID:   0,
	// }
	// db := database.New(map[int]models.StashItem{
	// 	0: si,
	// })
	s := server.New(db)
	err = s.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
