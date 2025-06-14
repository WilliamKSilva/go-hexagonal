package domain

type UserRepository interface {
	Save(uuid string, name string, password string, email string, role UserRole) (*User, error)
}
