package model

import "github.com/LaviqueDias/api-hotel-go/internal/room/model"

type HotelDTO struct {
	ID    int
	Name  string
	Rooms []model.Room
}

func HotelRequestToHotelDTO(hotelRequest HotelRequest) *HotelDTO {
	return &HotelDTO{
		Name: hotelRequest.Name,
	}
}

func HotelToHotelDTO(hotel *Hotel) *HotelDTO {
	return &HotelDTO{
		ID:    hotel.ID,
		Name:  hotel.Name,
		Rooms: hotel.Rooms,
	}
}

func HotelDTOToHotel(hotelDTO *HotelDTO) *Hotel {
	return &Hotel{
		ID:   hotelDTO.ID,
		Name: hotelDTO.Name,
		Rooms: hotelDTO.Rooms,
	}
}

func HotelDTOToHotelResponse(hotelDTO *HotelDTO) HotelResponse {
	return HotelResponse{
		ID:   hotelDTO.ID,
		Name: hotelDTO.Name,
		Rooms: hotelDTO.Rooms,
	}
}
