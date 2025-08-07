package model

import "github.com/LaviqueDias/api-hotel-go/internal/room/model"

type Hotel struct {
	ID    int
	Name  string
	Rooms []model.Room
}