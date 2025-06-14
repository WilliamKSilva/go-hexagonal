package domain

import "time"

type ReservationStatus int

const (
	Booked ReservationStatus = iota
	Cancelled
	Completed
)

type Reservation struct {
	UUID      string    `json:"uuid"`
	UserUUID  string    `json:"user_uuid"`
	RoomUUID  string    `json:"room_uuid"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func NewReservation(uuid string, userUUID string, roomUUID string, startTime time.Time, endTime time.Time) Reservation {
	return Reservation{
		UUID:      uuid,
		UserUUID:  userUUID,
		RoomUUID:  roomUUID,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

type ReservationService interface {
	Create(userUUID string, roomUUID string, startTime time.Time, endTime time.Time) (Reservation, error)
	ListByUserUUID(userUUID string) ([]Reservation, error)
}
