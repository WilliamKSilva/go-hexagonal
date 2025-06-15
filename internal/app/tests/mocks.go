package tests

import "github.com/WilliamKSilva/go-hexagonal/internal/domain"

// === Mock UUID Generator ===

type MockUUIDGen struct {
	UUID string
	Err  error
}

func (m *MockUUIDGen) Generate() (string, error) {
	return m.UUID, m.Err
}

// === Mock Crypt implementation ===

type MockCrypt struct {
	Hash string
	Err  error
}

func (m *MockCrypt) Encrypt(pw string) (string, error) {
	return m.Hash, m.Err
}

// === Mock User repo ===

type MockUserRepo struct {
	SavedUser *domain.User
	Err       error
}

func (m *MockUserRepo) Save(u domain.User) (*domain.User, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	m.SavedUser = &u
	return &u, nil
}

// === Mock Room repo ===

type MockRoomRepo struct {
	SavedRoom *domain.Room
	Err       error
}

func (m *MockRoomRepo) Save(r domain.Room) (*domain.Room, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	m.SavedRoom = &r
	return &r, nil
}

func (m *MockRoomRepo) Update(name *string, capacity *int32, status *domain.RoomStatus, maintenanceNote *string) error {
	if m.Err != nil {
		return m.Err
	}

	return nil
}

func (m *MockRoomRepo) SearchByUUID(uuid string) (*domain.Room, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	return m.SavedRoom, nil
}
