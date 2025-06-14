package domain

type UserRepository interface {
	Save(user User) (*User, error)
}
