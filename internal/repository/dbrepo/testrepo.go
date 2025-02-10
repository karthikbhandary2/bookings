package dbrepo

import (
	"errors"
	"time"

	"github.com/karthikbhandary2/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	return 1,nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(res models.RoomRestriction) error {
	
	return nil
}

// SearchAvailabilityByDates returns true if availability exists for roomID, and false if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	
	return false, nil
}

func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

func (m *testDBRepo) GetRoomById(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}
	return room, nil
}	

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User
	return u, nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 0, "", nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}
func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {

	var reservations []models.Reservation
	return reservations, nil
}	