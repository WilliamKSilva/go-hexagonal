package http

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":8080",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run() error {
	// Setup HTTP routes
	s.NewRoomHandler()

	log.Println("server listening at port :8080")

	log.Fatal(s.server.ListenAndServe())

	return nil
}
