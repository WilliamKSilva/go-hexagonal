package domain

type RoomStatus int

const (
	FREE RoomStatus = iota
	BOOKED
	MAINTENANCE
)

type Room struct {
	UUID            string     `json:"uuid"`
	Name            string     `json:"name"`
	Capacity        int32      `json:"capacity"`
	Status          RoomStatus `json:"status"`
	MaintenanceNote string     `json:"maintenance_note"`
}

func NewRoom(uuid string, name string, capacity int32) Room {
	return Room{
		UUID:            uuid,
		Name:            name,
		Capacity:        capacity,
		Status:          FREE,
		MaintenanceNote: "",
	}
}

type RoomService interface {
	Create(name string, capacity int32) (Room, error)
	Update(name *string, capacity *int32, status *int32, maintenanceNote *string) error
	Delete(uuid string) error
	ListAvaiable() ([]Room, error)
}
