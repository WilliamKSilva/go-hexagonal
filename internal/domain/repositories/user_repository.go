package repositories

import "github.com/WilliamKSilva/go-hexagonal/internal/domain"

type UserRepository interface {
	Save(user domain.User) (*domain.User, error)
}
