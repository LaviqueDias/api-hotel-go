package model

import "github.com/LaviqueDias/api-hotel-go/internal/room/model"

type HotelResponse struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Rooms []model.Room `json:"rooms"`
}
