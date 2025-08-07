package model

type RoomResponse struct {
	ID         int    `json:"id"`
	RoomNumber int    `json:"room_number"`
	Status     string `json:"status"`
	DailyPrice float64 `json:"daily_price"`
	HotelID    int    `json:"hotel_id"`
}