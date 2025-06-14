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

func (s *RoomService) Update(name *string, capacity *int32, isAvaiable *bool, maintenanceNote *string) error {
	err := s.RoomRepository.Update(
		name,
		capacity,
		isAvaiable,
		maintenanceNote,
	)
	if err != nil {
		return fmt.Errorf("error trying to update room on the database: %w", err)
	}

	return nil
}
