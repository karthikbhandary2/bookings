package repository

import "github.com/karthikbhandary2/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}