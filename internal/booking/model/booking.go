package model

import "time"

type Booking struct {
	ID           int       
	CheckinDate  time.Time 
	CheckoutDate time.Time 
	Status       string    
	TotalPrice   float64   
	GuestID      int       
}
