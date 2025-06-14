package domain

type Crypt interface {
	Encrypt(password string) (string, error)
}
