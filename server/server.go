package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ja-howell/stashclone/database"
)

type Server struct {
	db  *database.Database
	mux *http.ServeMux
}

func New(db *database.Database) Server {
	s := Server{}
	s.db = db
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("GET /stashitems/{id}", s.getStashItem)
	return s
}

func (s *Server) Run() error {
	return http.ListenAndServe(":8080", s.mux)
}

func (s *Server) getStashItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		return
	}

	si, err := s.db.GetStashItem(id)
	if err != nil {
		log.Printf("Failed to get stash item: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(si)
	if err != nil {
		log.Printf("Failed to encode: %v", err)
	}
}

func (s *Server) createStashItem(w http.ResponseWriter, r *http.Request) {

}
