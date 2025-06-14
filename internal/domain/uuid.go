package domain

type UUIDGenerator interface {
	Generate() (string, error)
}
