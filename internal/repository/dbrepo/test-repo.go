package dbrepo

import (
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"github.com/vitaLemoTea/secondstepweb/internal/repository"
	"time"
)

func (p *testDBRepo) GetUserByID(id int) bool {
	return true
}

// InsertReservation insert a reservation into the database and return the new reservation id
func (p *testDBRepo) InsertReservation(res model.Reservation) (int, error) {

	return 1, nil
}

func (p *testDBRepo) InsertRoomRestriction(res model.RoomRestriction) error {

	return nil
}

func (p *testDBRepo) SearchAvailabilityByDates(roomID int, startDate, endDate time.Time) (bool, error) {

	return true, nil
}

func (p *testDBRepo) SerachAvailabilityAllRooms(startDate, endDate time.Time) ([]model.Room, error) {

	var rooms []model.Room

	return rooms, nil
}

func (p *testDBRepo) GetRoomByID(roomID int) (model.Room, error) {

	var room model.Room

	return room, nil
}

var _ repository.DatabaseRepo = (*testDBRepo)(nil)
