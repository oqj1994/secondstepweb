package dbrepo

import (
	"context"
	"github.com/vitaLemoTea/secondstepweb/internal/model"
	"time"
)

func (p *postgresDBRepo) GetUserByID(id int) bool {
	return true
}

// InsertReservation insert a reservation into the database and return the new reservation id
func (p *postgresDBRepo) InsertReservation(res model.Reservation) (int, error) {
	//set a ctx with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	stmt := `insert into reservations (first_name,last_name ,phone,email,start_date,end_date,room_id
	,created_at,updated_at
)  values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`

	row := p.DB.QueryRowContext(ctx,
		stmt,
		res.FirstName,
		res.LastName,
		res.Phone,
		res.Email,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)
	var newReservationID int
	err := row.Scan(&newReservationID)
	if err != nil {
		return 0, err
	}
	return newReservationID, nil
}

func (p *postgresDBRepo) InsertRoomRestriction(res model.RoomRestriction) error {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	query := `insert into room_restrictions (
                               start_date,end_date,room_id,reservation_id,restriction_id,
                               created_at,updated_at)
				values ($1,$2,$3,$4,$5,$6,$7)
`
	_, err := p.DB.ExecContext(
		ctx,
		query,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		res.ReservationID,
		res.RestrictionID,
		res.CreateAt,
		res.UpdateAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresDBRepo) SearchAvailabilityByDates(roomID int, startDate, endDate time.Time) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var stmt string
	var numRow int
	if startDate == endDate {
		stmt = `select count(*) from room_restrictions where 
			(start_date= $1 or end_date=$2)and room_id=$3	`

	} else {
		stmt = `select count(*) from room_restrictions where 
			start_date< $1 and end_date > $2 and room_id=$3	`
	}

	err := p.DB.QueryRowContext(ctx, stmt, endDate, startDate, roomID).Scan(&numRow)
	if err != nil {
		return false, err
	}
	return numRow == 0, nil
}

func (p *postgresDBRepo) SerachAvailabilityAllRooms(startDate, endDate time.Time) ([]model.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var query string
	if startDate != endDate {
		query = `
		select r.id,r.room_name from rooms r
		where r.id not in (
			select rr.room_id from room_restrictions rr where $1 <rr.end_date and $2> rr.start_date
		)`
	} else {
		query = `
		select r.id,r.room_name from rooms r
		where r.id not in (
			select rr.room_id from room_restrictions rr where $1 =rr.end_date or $2= rr.start_date
		)`
	}
	var rooms []model.Room
	rows, err := p.DB.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (p *postgresDBRepo) GetRoomByID(roomID int) (model.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var room model.Room
	query := `select room_name ,created_at,updated_at from rooms where id=$1`
	row := p.DB.QueryRowContext(ctx, query, roomID)
	err := row.Scan(&room.RoomName, &room.CreateAt, &room.UpdateAt)
	if err != nil {
		return room, err
	}
	room.ID = roomID
	return room, nil
}
