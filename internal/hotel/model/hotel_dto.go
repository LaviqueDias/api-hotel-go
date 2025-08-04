package model

type HotelDTO struct {
	ID   int
	Name string
}

func FromHotelRequest(hotelRequest HotelRequest) *HotelDTO {
	return &HotelDTO{
		Name: hotelRequest.Name,
	}
}

func ToHotelDTO(hotel Hotel) *HotelDTO {
	return &HotelDTO{
		ID: hotel.ID,
		Name: hotel.Name,
	}
}

func FromHotelDTO(hotelDTO HotelDTO) *Hotel {
	return &Hotel{
		ID:   hotelDTO.ID,
		Name: hotelDTO.Name,
	}
}

func ToHotelResponse(hotelDTO *HotelDTO) HotelResponse {
	return HotelResponse{
		ID: hotelDTO.ID,
		Name: hotelDTO.Name,
	}
}
