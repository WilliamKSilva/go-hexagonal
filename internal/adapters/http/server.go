package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/WilliamKSilva/go-hexagonal/internal/adapters/http/responses"
	"github.com/WilliamKSilva/go-hexagonal/internal/app"
	"github.com/WilliamKSilva/go-hexagonal/internal/app/tests"
	"github.com/WilliamKSilva/go-hexagonal/internal/domain"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	// Server instance
	server   *http.Server
	validate *validator.Validate

	// Internal services
	roomService *app.RoomService
}

func NewServer() *Server {
	roomRepository := tests.MockRoomRepo{
		SavedRoom: &domain.Room{
			UUID:            "uuid-123",
			Name:            "room-1",
			Capacity:        4,
			Status:          domain.MAINTENANCE,
			MaintenanceNote: "",
		},
	}

	uuidGenerator := tests.MockUUIDGen{
		UUID: "uuid-123",
	}

	roomService := app.NewRoomService(
		&roomRepository,
		&uuidGenerator,
	)

	return &Server{
		server: &http.Server{
			Addr:           ":8080",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		validate:    validator.New(),
		roomService: roomService,
	}
}

func (s *Server) Run() error {
	// Is necessary to have two handlers, one for routes with paths and one for basic routes.
	// Example: if you try to call a handler registered for /rooms with a path you will get
	// not found errors.
	// Example: if you try to call a handler registered with /rooms/ and calls /rooms all the
	// HTTP methods are translated to GET.

	// Setup HTTP handlers
	roomHandlerWithoutPaths := RoomHandlerWithoutPaths{
		validate:    s.validate,
		roomService: s.roomService,
	}
	http.HandleFunc("/rooms", roomHandlerWithoutPaths.ServeHTTP)

	log.Println("server listening at port :8080")

	log.Fatal(s.server.ListenAndServe())

	return nil
}

// I choosed to make route path handling from scratch
// for learning purposes. I know that more fancier options
// exist on Golang ecosystem, like Gin, Gorilla Mux, etc.
func ParseRoute(url string) []string {
	path := strings.Trim(url, "/")
	parts := strings.Split(path, "/")

	return parts
}

func ParseBody[T any](w http.ResponseWriter, r *http.Request, route string) (T, error) {
	var payload T

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return payload, fmt.Errorf("internal server error")
	}

	err = json.Unmarshal(b, &payload)
	if err != nil {
		return payload, fmt.Errorf("internal server error")
	}

	return payload, nil
}

func MakeSuccessResponse(w http.ResponseWriter, payload any, code int, route string) {
	b, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// If can't parse JSON payload response for any reason
		// fallback to return body as raw bytes
		w.Write([]byte(fmt.Sprintf("%s: internal server error", route)))
		return
	}

	w.WriteHeader(code)
	w.Write(b)
}

func MakeErrorResponse(w http.ResponseWriter, res responses.HTTPResponse, route string) {
	res.Message = fmt.Sprintf("%s: %s", route, res.Message)
	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// If can't parse JSON payload response for any reason
		// fallback to return body as raw bytes
		w.Write([]byte(fmt.Sprintf("%s: internal server error", route)))
		return
	}

	w.WriteHeader(res.Code)
	w.Write(b)
}
