package http

import (
	"net/http"

	"github.com/WilliamKSilva/go-hexagonal/internal/adapters/http/requests"
	"github.com/WilliamKSilva/go-hexagonal/internal/adapters/http/responses"
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
	var res responses.HTTPResponse
	payload, err := ParseBody[requests.CreateRoom](w, r, "create room")
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		MakeErrorResponse(w, res, "create room")
		return
	}

	room, err := h.roomService.Create(payload.Name, payload.Capacity)
	if err != nil {
		// TODO: Add error handling to check type of error returned
		// by the service to return proper status code
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		MakeErrorResponse(w, res, "create room")
		return
	}

	err = ValidateFields(h.validate, "create room", payload)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		MakeErrorResponse(w, res, "create room")
		return
	}

	MakeSuccessResponse(w, room, http.StatusCreated, "create room")
}
