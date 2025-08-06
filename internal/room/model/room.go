package model

type Room struct {
	ID         int    `json:"id"`
	RoomNumber int    `json:"room_number"`
	Status     string `json:"status" binding:"required,oneof=disponivel manutencao ocupado"`
	HotelID    int    `json:"hotel_id" binding:"required"`
}
