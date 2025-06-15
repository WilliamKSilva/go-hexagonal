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

func (s *RoomService) Update(name *string, capacity *int32, status *domain.RoomStatus, maintenanceNote *string) error {
	err := s.RoomRepository.Update(
		name,
		capacity,
		status,
		maintenanceNote,
	)
	if err != nil {
		return fmt.Errorf("error trying to update room on the database: %w", err)
	}

	return nil
}

func (s *RoomService) StartMaintenance(roomUUID string, maintenanceNote string) error {
	room, err := s.RoomRepository.SearchByUUID(roomUUID)
	if err != nil {
		return fmt.Errorf("error trying to search room on the database: %w", err)
	}

	if room == nil {
		return fmt.Errorf("room not found")
	}

	if room.Status == domain.MAINTENANCE {
		return fmt.Errorf("room is already on maintenance")
	}

	if room.Status == domain.BOOKED {
		return fmt.Errorf("room is currently booked")
	}

	status := domain.MAINTENANCE

	err = s.RoomRepository.Update(
		nil,
		nil,
		&status,
		&maintenanceNote,
	)
	if err != nil {
		return fmt.Errorf("error trying to update room on the database: %w", err)
	}

	return nil
}

func (s *RoomService) Delete(roomUUID string) error {
	room, err := s.RoomRepository.SearchByUUID(roomUUID)
	if err != nil {
		return fmt.Errorf("error trying to search room on the database: %w", err)
	}

	if room == nil {
		return fmt.Errorf("room not found")
	}

	// Can't delete rooms that are being used in any way
	if room.Status == domain.MAINTENANCE {
		return fmt.Errorf("room is currently on maintenance")
	}

	if room.Status == domain.BOOKED {
		return fmt.Errorf("room is currently booked")
	}

	err = s.RoomRepository.Delete(roomUUID)
	if err != nil {
		return fmt.Errorf("error trying to delete room on the database: %w", err)
	}

	return nil
}
