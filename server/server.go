package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ja-howell/stashclone/models"
)

type Server struct {
	db  Database
	mux *http.ServeMux
}

func New(db Database) Server {
	s := Server{}
	s.db = db
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("GET /stashitems/{id}", s.getStashItem)
	s.mux.HandleFunc("GET /stashitems", s.listStashItems)
	s.mux.HandleFunc("POST /stashitems", s.createStashItem)
	s.mux.HandleFunc("PUT /stashitems/{id}", s.updateStashItem)
	s.mux.HandleFunc("DELETE /stashitems/{id}", s.deleteStashItem)

	return s
}

func (s *Server) Run() error {
	return http.ListenAndServe(":8080", s.mux)
}

func (s *Server) getStashItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	si, err := s.db.GetStashItem(id)
	if err != nil {
		log.Printf("Failed to get stash item: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(si)
	if err != nil {
		log.Printf("Failed to encode: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (s *Server) listStashItems(w http.ResponseWriter, r *http.Request) {

	sis, err := s.db.ListStashItems()
	if err != nil {
		log.Printf("Failed to get stash items: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(sis)
	if err != nil {
		log.Printf("Failed to encode: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (s *Server) createStashItem(w http.ResponseWriter, r *http.Request) {
	si := models.StashItem{}
	err := json.NewDecoder(r.Body).Decode(&si)
	if err != nil {
		log.Printf("Failed to create stash item: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = s.db.CreateStashItem(si)
	if err != nil {
		log.Printf("Failed to save stash item to database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("Created stash item: %v", http.StatusAccepted)
	w.WriteHeader(http.StatusAccepted)
}

//TODO: look up http.Error()

func (s *Server) updateStashItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	si := models.StashItem{}
	err = json.NewDecoder(r.Body).Decode(&si)
	if err != nil {
		log.Printf("Failed to update stash item: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.db.UpdateStashItem(id, si)
	if err != nil {
		log.Printf("Failed to save updated stash item to database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Updated stash item: %v", http.StatusAccepted)
	w.WriteHeader(http.StatusAccepted)

}

func (s *Server) deleteStashItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.db.DeleteStashItem(id)
	if err != nil {
		log.Printf("Failed to delete stash item: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Deleted stash item %d: %v", id, http.StatusAccepted)
}
