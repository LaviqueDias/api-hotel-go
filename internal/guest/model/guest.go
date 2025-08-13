package model

type Guest struct {
	ID        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
}