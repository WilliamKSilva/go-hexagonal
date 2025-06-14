package app

import (
	"errors"
	"testing"

	"github.com/WilliamKSilva/go-hexagonal/internal/app"
	"github.com/WilliamKSilva/go-hexagonal/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// === Mock Interfaces ===

type mockUUIDGen struct {
	UUID string
	Err  error
}

func (m *mockUUIDGen) Generate() (string, error) {
	return m.UUID, m.Err
}

type mockCrypt struct {
	Hash string
	Err  error
}

func (m *mockCrypt) Encrypt(pw string) (string, error) {
	return m.Hash, m.Err
}

type mockUserRepo struct {
	SavedUser *domain.User
	Err       error
}

func (m *mockUserRepo) Save(u domain.User) (*domain.User, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	m.SavedUser = &u
	return &u, nil
}

// === Tests ===

func TestUserService_Create(t *testing.T) {
	tests := []struct {
		name          string
		uuidGen       *mockUUIDGen
		crypt         *mockCrypt
		repo          *mockUserRepo
		expectErr     bool
		expectedError string
	}{
		{
			name: "successfully creates user",
			uuidGen: &mockUUIDGen{
				UUID: "uuid-123",
			},
			crypt: &mockCrypt{
				Hash: "hashed-password",
			},
			repo:      &mockUserRepo{},
			expectErr: false,
		},
		{
			name: "uuid generator fails",
			uuidGen: &mockUUIDGen{
				Err: errors.New("uuid error"),
			},
			crypt:         &mockCrypt{},
			repo:          &mockUserRepo{},
			expectErr:     true,
			expectedError: "uuid error",
		},
		{
			name: "crypt fails",
			uuidGen: &mockUUIDGen{
				UUID: "uuid-123",
			},
			crypt: &mockCrypt{
				Err: errors.New("crypt error"),
			},
			repo:          &mockUserRepo{},
			expectErr:     true,
			expectedError: "crypt error",
		},
		{
			name: "repository save fails",
			uuidGen: &mockUUIDGen{
				UUID: "uuid-123",
			},
			crypt: &mockCrypt{
				Hash: "hashed-password",
			},
			repo: &mockUserRepo{
				Err: errors.New("db error"),
			},
			expectErr:     true,
			expectedError: "db error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &app.UserService{
				UserRepo: tt.repo,
				UuidGen:  tt.uuidGen,
				Crypt:    tt.crypt,
			}

			user, err := service.Create("Alice", "super-secret", "alice@mail.com", domain.Guest)

			if tt.expectErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, user)
			} else {
				require.NoError(t, err)
				require.NotNil(t, user)
				assert.Equal(t, "uuid-123", user.UUID)
				assert.Equal(t, "Alice", user.Name)
				assert.Equal(t, "hashed-password", user.Password)
				assert.Equal(t, domain.Guest, user.Role)
			}
		})
	}
}
