package domain

type UserRole int

const (
	Guest UserRole = iota
	Admin
)

type User struct {
	UUID     string   `json:"uuid"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

func NewUser(uuid string, name string, email string, password string, role UserRole) *User {
	return &User{
		UUID:     uuid,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

type UserService interface {
	Create(name string, password string, email string, role UserRole) (*User, error)
}
