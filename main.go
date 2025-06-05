package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ja-howell/stashclone/models"
)

func main() {
	http.HandleFunc("GET /stashitems/{id}", getStashItem)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getStashItem(w http.ResponseWriter, r *http.Request) {
	si := models.StashItem{
		Name: "foo",
		ID:   1,
	}
	err := json.NewEncoder(w).Encode(si)
	if err != nil {
		log.Printf("Failed to encode: %v", err)
	}
}
