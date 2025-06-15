package http

import (
	"log"
	"net/http"

	"github.com/WilliamKSilva/go-hexagonal/internal/adapters/http/requests"
	"github.com/WilliamKSilva/go-hexagonal/internal/app"
	"github.com/go-playground/validator/v10"
)

type RoomHandlerWithoutPaths struct {
	validate *validator.Validate

	roomService *app.RoomService
}

func (h RoomHandlerWithoutPaths) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := ParseRoute(r.URL.String())

	withoutPathParameters := len(parts) == 1

	if withoutPathParameters {
		if r.Method == http.MethodPost {
			h.Create(w, r)
		}
	}
}

func (h RoomHandlerWithoutPaths) Create(w http.ResponseWriter, r *http.Request) {
	payload := ParseBody[requests.CreateRoom](w, r)
	log.Println(payload)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("iuahsdiuashd"))
}
