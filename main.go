package main

import (
	"log"

	"github.com/ja-howell/stashclone/database"
	"github.com/ja-howell/stashclone/models"
	"github.com/ja-howell/stashclone/server"
)

func main() {
	si := models.StashItem{
		Name: "foo",
		ID:   0,
	}
	db := database.New([]models.StashItem{si})
	s := server.New(db)
	err := s.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
