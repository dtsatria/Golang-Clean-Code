package dto

import "final-project-booking-room/model"

type BookingRequestDto struct {
	Id              string                `json:"id"`
	BoookingDetails []model.BookingDetail `json:"bookingDetails" binding:"required"`
	Description     string                `json:"description"`
}
