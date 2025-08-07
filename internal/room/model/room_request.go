package model

type RoomRequest struct {
	RoomNumber int    `json:"room_number"`
	Status     string `json:"status" binding:"required,oneof=disponivel manutencao ocupado"`
	HotelID    int    `json:"hotel_id"`
}


