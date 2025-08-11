package model

import "time"

type BookingDTO struct {
	ID           int
	CheckinDate  time.Time
	CheckoutDate time.Time
	Status       string
	TotalPrice   float64
	GuestID      int
}