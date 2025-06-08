package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ja-howell/stashclone/database"
)

type Server struct {
	db  *database.Database
	mux *http.ServeMux
}

func New() Server {
	s := Server{}
	s.db = &database.Database{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("GET /stashitems/{id}", s.getStashItem)
	return s
}

func (s *Server) Start() {
	log.Fatal(http.ListenAndServe(":8080", s.mux))
}

func (s *Server) getStashItem(w http.ResponseWriter, r *http.Request) {
	si, err := s.db.GetStashItem(0)
	if err != nil {
		log.Printf("Failed to get stash item: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(si)
	if err != nil {
		log.Printf("Failed to encode: %v", err)
	}
}
