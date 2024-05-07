package api

import (
	"encoding/json"
	"net/http"

	"github.com/Devon-ODell/testcode/storage"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store store) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserbyID)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUserbyID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	json.NewEncoder(w).Encode(user)

}

func ValidateUser(u *User) bool { return true }
