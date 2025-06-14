package tests

import "github.com/WilliamKSilva/go-hexagonal/internal/domain"

// === Mock Interfaces ===

type MockUUIDGen struct {
	UUID string
	Err  error
}

func (m *MockUUIDGen) Generate() (string, error) {
	return m.UUID, m.Err
}

type MockCrypt struct {
	Hash string
	Err  error
}

func (m *MockCrypt) Encrypt(pw string) (string, error) {
	return m.Hash, m.Err
}

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
