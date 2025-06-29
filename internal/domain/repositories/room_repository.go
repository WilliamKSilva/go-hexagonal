package repositories

import "github.com/WilliamKSilva/go-hexagonal/internal/domain"

type RoomRepository interface {
	Save(domain.Room) (*domain.Room, error)
	Update(name *string, capacity *int32, status *domain.RoomStatus, maintenanceNote *string) error
	Delete(uuid string) error
	SearchByUUID(uuid string) (*domain.Room, error)
}
