package requests

type CreateRoom struct {
	Name     string `json:"name" validate:"required"`
	Capacity int32  `json:"capacity" validate:"required"`
}
