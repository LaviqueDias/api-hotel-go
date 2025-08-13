package model

import "time"

type BookingRequest struct {
	CheckinDate  time.Time `json:"checkin_date"`
	CheckoutDate time.Time `json:"checkout_date"`
	Status       string    `json:"status" binding:"required,oneof=booked cancelled completed"`
	TotalPrice   float64   `json:"total_price"`
	GuestID      int       `json:"guest_id"`
	RoomIDs      []int     `json:"room_ids" binding:"required,min=1"`
}
