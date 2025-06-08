package main

import (
	"log"

	"github.com/ja-howell/stashclone/database"
	"github.com/ja-howell/stashclone/server"
)

func main() {
	s := server.New(&database.Database{})
	err := s.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
