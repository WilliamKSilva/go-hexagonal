package app

import (
	"fmt"

	"github.com/WilliamKSilva/go-hexagonal/internal/domain"
	"github.com/WilliamKSilva/go-hexagonal/internal/domain/repositories"
)

type RoomService struct {
	RoomRepository repositories.RoomRepository
	UUIDGenerator  domain.UUIDGenerator
}

func (s *RoomService) Create(name string, capacity int32) (*domain.Room, error) {
	uuid, err := s.UUIDGenerator.Generate()
	if err != nil {
		return nil, fmt.Errorf("error trying to generate UUID for new Room: %w", err)
	}

	room := domain.NewRoom(uuid, name, capacity)
	roomSaved, err := s.RoomRepository.Save(room)
	if err != nil {
		return nil, fmt.Errorf("error trying to save room on the database: %w", err)
	}

	return roomSaved, nil
}
