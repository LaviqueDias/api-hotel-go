package model

import "github.com/LaviqueDias/api-hotel-go/internal/room/model"

type Hotel struct {
	ID    int
	Name  string
	Rooms []model.Room
}

func (h *Hotel) AddRoom(room model.Room) {
	h.Rooms = append(h.Rooms, room)
}