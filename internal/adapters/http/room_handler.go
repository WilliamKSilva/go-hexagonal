package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/WilliamKSilva/go-hexagonal/internal/adapters/http/requests"
	"github.com/WilliamKSilva/go-hexagonal/internal/adapters/http/responses"
	"github.com/go-playground/validator/v10"
)

type RoomHandlerWithoutPaths struct {
	validate *validator.Validate
}

func (h RoomHandlerWithoutPaths) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := ParseRoute(r.URL.String())

	withoutPathParameters := len(parts) == 1

	if withoutPathParameters {
		log.Println(withoutPathParameters, r.Method)
		if r.Method == http.MethodPost {
			h.Create(w, r)
		}
	}

	if len(parts) == 2 && r.Method == http.MethodGet {
		log.Println("get by uuid")
	}
}

func (h RoomHandlerWithoutPaths) Create(w http.ResponseWriter, r *http.Request) {
	var payload requests.CreateRoom
	var res responses.HTTPResponse

	b, err := io.ReadAll(r.Body)
	if err != nil {
		res = responses.NewInternalServerError("create room")
		b, err = json.Marshal(res)
		if err != nil {
			log.Println("create room: json encoding error %w", err)
			http.Error(w, res.Message, res.Code)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(b)
	}

	err = json.Unmarshal(b, &payload)
	if err != nil {
		res = responses.NewInternalServerError("create room")
		b, err = json.Marshal(res)
		if err != nil {
			log.Println("create room: json encoding error %w", err)
			http.Error(w, res.Message, res.Code)
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
	}

	w.WriteHeader(http.StatusCreated)
}
