package model

import "time"

type BookingResponse struct {
	ID           int       `json:"id"`
	CheckinDate  time.Time `json:"checkin_date"`
	CheckoutDate time.Time `json:"checkout_date"`
	Status       string    `json:"status"`       
	TotalPrice   float64   `json:"total_price"`  
	GuestID      int       `json:"guest_id"`
}
