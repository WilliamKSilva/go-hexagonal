package app

import (
	"fmt"

	"github.com/WilliamKSilva/go-hexagonal/internal/domain"
	"github.com/WilliamKSilva/go-hexagonal/internal/domain/repositories"
)

type UserService struct {
	domain.UserService

	UserRepo repositories.UserRepository
	UuidGen  domain.UUIDGenerator
	Crypt    domain.Crypt
}

func (s *UserService) Create(name string, password string, email string, role domain.UserRole) (*domain.User, error) {
	uuid, err := s.UuidGen.Generate()
	if err != nil {
		return nil, fmt.Errorf("error trying to generate a new UUID for the user: %w", err)
	}

	hash, err := s.Crypt.Encrypt(password)
	if err != nil {
		return nil, fmt.Errorf("error trying to encrypt user password: %w", err)
	}

	user := domain.NewUser(uuid, name, email, hash, role)

	userSaved, err := s.UserRepo.Save(*user)
	if err != nil {
		return nil, fmt.Errorf("error trying to save user on the database: %w", err)
	}

	return userSaved, nil
}
