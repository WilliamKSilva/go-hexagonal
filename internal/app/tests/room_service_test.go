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
			Status:          domain.FREE,
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
			assert.Equal(t, room.Status, expect.Status)
			assert.Equal(t, room.MaintenanceNote, expect.MaintenanceNote)
		}
	}
}

func TestRoomService_Update(t *testing.T) {
	tests := []struct {
		name          string
		repo          *MockRoomRepo
		expectErr     bool
		expectedError string
	}{
		{
			name:      "successfully updates room",
			repo:      &MockRoomRepo{},
			expectErr: false,
		},
		{
			name: "repository save fails",
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
		}

		name := "room-2"
		capacity := int32(10)
		status := domain.FREE
		maintenanceNote := "Fixing air conditioner"

		err := svc.Update(&name, &capacity, &status, &maintenanceNote)
		if tt.expectErr {
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedError)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRoomService_StartMaintenance(t *testing.T) {
	tests := []struct {
		name          string
		repo          *MockRoomRepo
		expectErr     bool
		expectedError string
	}{
		{
			name: "successfully start room maintenance",
			repo: &MockRoomRepo{
				SavedRoom: &domain.Room{
					UUID:            "uuid-123",
					Name:            "room-123",
					Capacity:        4,
					Status:          domain.FREE,
					MaintenanceNote: "",
				},
			},
			expectErr: false,
		},
		{
			name: "repository search by UUID fails",
			repo: &MockRoomRepo{
				Err: errors.New("repository fails"),
			},
			expectErr:     true,
			expectedError: "repository fails",
		},
		{
			name: "repository user not found",
			repo: &MockRoomRepo{
				SavedRoom: nil,
			},
			expectErr:     true,
			expectedError: "room not found",
		},
		{
			name: "room is already on maintenace",
			repo: &MockRoomRepo{
				SavedRoom: &domain.Room{
					UUID:            "uuid-123",
					Name:            "room-123",
					Capacity:        4,
					Status:          domain.MAINTENANCE,
					MaintenanceNote: "fixing bed",
				},
			},
			expectErr:     true,
			expectedError: "room is already on maintenance",
		},
		{
			name: "room is currently booked",
			repo: &MockRoomRepo{
				SavedRoom: &domain.Room{
					UUID:            "uuid-123",
					Name:            "room-123",
					Capacity:        4,
					Status:          domain.BOOKED,
					MaintenanceNote: "",
				},
			},
			expectErr:     true,
			expectedError: "room is currently booked",
		},
	}

	for _, tt := range tests {
		svc := app.RoomService{
			RoomRepository: tt.repo,
		}

		err := svc.StartMaintenance("uuid-123", "fixing air conditioner")
		if tt.expectErr {
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedError)
		} else {
			require.NoError(t, err)
		}
	}
}
