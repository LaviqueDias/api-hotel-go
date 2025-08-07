package model

type RoomRequest struct {
	RoomNumber int    `json:"room_number" binding:"required"`
	Status     string `json:"status" binding:"required,oneof='available maintenance occupied"`
	DailyPrice float64 `json:"daily_price" binding:"required"`
	HotelID    int    `json:"hotel_id" binding:"required"`
}


