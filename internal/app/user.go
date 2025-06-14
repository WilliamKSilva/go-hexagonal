package app

import "github.com/WilliamKSilva/go-hexagonal/internal/domain"

type UserService struct {
	domain.UserService
}

func (userService *UserService) Create(name string, email string, role domain.UserRole) (*domain.User, error) {
	return nil, nil
}
