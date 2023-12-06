package dto

type Approval struct {
	Approval        string `json:"approval" binding:"required"`
	BookingDetailId string `json:"bookingDetailId" binding:"required"`
}
