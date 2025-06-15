package http

import (
	"net/http"
)

func (s *Server) NewRoomHandler() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
}
