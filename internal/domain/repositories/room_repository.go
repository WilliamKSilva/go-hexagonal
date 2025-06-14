package repositories

import "github.com/WilliamKSilva/go-hexagonal/internal/domain"

type RoomRepository interface {
	Save(domain.Room) (*domain.Room, error)
}
