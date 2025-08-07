package model

type RoomResponse struct {
	ID         int    `json:"id"`
	RoomNumber int    `json:"room_number"`
	Status     string `json:"status"`
	HotelID    int    `json:"hotel_id"`
}