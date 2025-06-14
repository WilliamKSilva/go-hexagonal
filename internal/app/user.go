package app

import (
	"fmt"

	"github.com/WilliamKSilva/go-hexagonal/internal/domain"
)

type UserService struct {
	domain.UserService

	userRepo domain.UserRepository
	uuidGen  domain.UUIDGenerator
	crypt    domain.Crypt
}

func (s *UserService) Create(name string, password string, email string, role domain.UserRole) (*domain.User, error) {
	uuid, err := s.uuidGen.Generate()
	if err != nil {
		return nil, fmt.Errorf("error trying to generate a new UUID for the user: %w", err)
	}

	hash, err := s.crypt.Encrypt(password)
	if err != nil {
		return nil, fmt.Errorf("error trying to encrypt user password: %w", err)
	}

	user, err := s.userRepo.Save(uuid, name, hash, email, role)
	if err != nil {
		return nil, fmt.Errorf("error trying to save user on the database: %w", err)
	}

	return user, nil
}
