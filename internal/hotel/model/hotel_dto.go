package model

type HotelDTO struct {
	ID   int
	Name string
}

func HotelRequestToHotelDTO(hotelRequest HotelRequest) *HotelDTO {
	return &HotelDTO{
		Name: hotelRequest.Name,
	}
}

func HotelToHotelDTO(hotel *Hotel) *HotelDTO {
	return &HotelDTO{
		ID: hotel.ID,
		Name: hotel.Name,
	}
}

func HotelDTOToHotel(hotelDTO *HotelDTO) *Hotel {
	return &Hotel{
		ID:   hotelDTO.ID,
		Name: hotelDTO.Name,
	}
}

func HotelDTOToHotelResponse(hotelDTO *HotelDTO) HotelResponse {
	return HotelResponse{
		ID: hotelDTO.ID,
		Name: hotelDTO.Name,
	}
}
