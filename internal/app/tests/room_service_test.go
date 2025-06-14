package tests

import (
	"errors"
	"testing"

	"github.com/WilliamKSilva/go-hexagonal/internal/app"
	"github.com/WilliamKSilva/go-hexagonal/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoomService_Create(t *testing.T) {
	tests := []struct {
		name          string
		uuidGen       *MockUUIDGen
		repo          *MockRoomRepo
		expectErr     bool
		expectedError string
	}{
		{
			name: "successfully creates room",
			uuidGen: &MockUUIDGen{
				UUID: "uuid-123",
			},
			repo:      &MockRoomRepo{},
			expectErr: false,
		},
		{
			name: "uuid generator fails",
			uuidGen: &MockUUIDGen{
				Err: errors.New("uuid fails"),
			},
			repo:          &MockRoomRepo{},
			expectErr:     true,
			expectedError: "uuid fails",
		},
		{
			name: "repository save fails",
			uuidGen: &MockUUIDGen{
				UUID: "uuid-123",
			},
			repo: &MockRoomRepo{
				Err: errors.New("repository fails"),
			},
			expectErr:     true,
			expectedError: "repository fails",
		},
	}

	for _, tt := range tests {
		svc := app.RoomService{
			RoomRepository: tt.repo,
			UUIDGenerator:  tt.uuidGen,
		}

		expect := domain.Room{
			UUID:            "uuid-123",
			Name:            "room-1",
			Capacity:        4,
			IsAvaiable:      true,
			MaintenanceNote: "",
		}

		room, err := svc.Create("room-1", 4)
		if tt.expectErr {
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedError)
		} else {
			require.NoError(t, err)
			require.NotNil(t, room)
			assert.Equal(t, room.UUID, expect.UUID)
			assert.Equal(t, room.Name, expect.Name)
			assert.Equal(t, room.Capacity, expect.Capacity)
			assert.Equal(t, room.IsAvaiable, expect.IsAvaiable)
			assert.Equal(t, room.MaintenanceNote, expect.MaintenanceNote)
		}
	}
}
