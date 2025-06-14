package tests

import (
	"errors"
	"testing"

	"github.com/WilliamKSilva/go-hexagonal/internal/app"
	"github.com/WilliamKSilva/go-hexagonal/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// === Tests ===

func TestUserService_Create(t *testing.T) {
	tests := []struct {
		name          string
		uuidGen       *MockUUIDGen
		crypt         *MockCrypt
		repo          *MockUserRepo
		expectErr     bool
		expectedError string
	}{
		{
			name: "successfully creates user",
			uuidGen: &MockUUIDGen{
				UUID: "uuid-123",
			},
			crypt: &MockCrypt{
				Hash: "hashed-password",
			},
			repo:      &MockUserRepo{},
			expectErr: false,
		},
		{
			name: "uuid generator fails",
			uuidGen: &MockUUIDGen{
				Err: errors.New("uuid error"),
			},
			crypt:         &MockCrypt{},
			repo:          &MockUserRepo{},
			expectErr:     true,
			expectedError: "uuid error",
		},
		{
			name: "crypt fails",
			uuidGen: &MockUUIDGen{
				UUID: "uuid-123",
			},
			crypt: &MockCrypt{
				Err: errors.New("crypt error"),
			},
			repo:          &MockUserRepo{},
			expectErr:     true,
			expectedError: "crypt error",
		},
		{
			name: "repository save fails",
			uuidGen: &MockUUIDGen{
				UUID: "uuid-123",
			},
			crypt: &MockCrypt{
				Hash: "hashed-password",
			},
			repo: &MockUserRepo{
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
