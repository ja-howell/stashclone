package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ja-howell/stashclone/database"
)

func main() {
	http.HandleFunc("GET /stashitems/{id}", getStashItem)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getStashItem(w http.ResponseWriter, r *http.Request) {
	db := database.Database{}
	si, err := db.GetStashItem(0)
	if err != nil {
		log.Printf("Failed to get stash item: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(si)
	if err != nil {
		log.Printf("Failed to encode: %v", err)
	}
}
