package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/ja-howell/stashclone/database"
	"github.com/ja-howell/stashclone/server"
)

func main() {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "stash"

	if cfg.User == "" || cfg.Passwd == "" {
		log.Fatal("DBUSER and DBPASS environment variables must be set")
	}

	// Get a database handle.
	mysqldb, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}

	pingErr := mysqldb.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	s := server.New(database.NewMySQL(mysqldb))
	err = s.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
