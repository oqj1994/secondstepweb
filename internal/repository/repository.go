package repository

import (
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"time"
)

type DatabaseRepo interface {
	GetUserByID(id int) bool
	InsertReservation(res model.Reservation) (int, error)
	InsertRoomRestriction(res model.RoomRestriction) error
	SearchAvailabilityByDates(roomID int, startDate, endDate time.Time) (bool, error)
	SerachAvailabilityAllRooms(startDate, endDate time.Time) ([]model.Room, error)
	GetRoomByID(roomID int) (model.Room, error)
}
