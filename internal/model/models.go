package model

import "time"

// DBModel hold the share field
type DBModel struct {
	CreateAt time.Time
	UpdateAt time.Time
}

// User is the user model
type User struct {
	DBModel
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PassWord    string
	AccessLevel int
}

// Room is the room model
type Room struct {
	DBModel
	ID       int
	RoomName string
}

// Restriction is the restriction model
type Restriction struct {
	DBModel
	ID              int
	RestrictionName string
}

// Reservation is the reservation model
type Reservation struct {
	DBModel
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	Room      Room
}

// RoomRestriction is the roomRestriction model
type RoomRestriction struct {
	DBModel
	ID            int
	ReservationID int
	Reservation   Reservation
	RestrictionID int
	Restriction   Restriction
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	Room          Room
}
