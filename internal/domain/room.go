package domain

type Room struct {
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	Capacity        int32  `json:"capacity"`
	IsAvaiable      bool   `json:"is_avaiable"`
	MaintenanceNote string `json:"maintenance_note"`
}

func NewRoom(uuid string, name string, capacity int32) Room {
	return Room{
		UUID:            uuid,
		Name:            name,
		Capacity:        capacity,
		IsAvaiable:      true,
		MaintenanceNote: "",
	}
}

type RoomService interface {
	Create(name string, capacity int32) (Room, error)
	Update(uuid string, name string, capacity int32) error
	SetMaintenance(uuid string) error
	Delete(uuid string) error
	ListAvaiable() ([]Room, error)
}
