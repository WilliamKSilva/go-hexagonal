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
