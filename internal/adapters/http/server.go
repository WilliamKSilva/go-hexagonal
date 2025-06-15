package http

import (
	"log"
	"net/http"
	"strings"
	"time"

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
		validate: s.validate,
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
