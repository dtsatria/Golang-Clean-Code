package usecasemock

import (
	"final-project-booking-room/model"
	"final-project-booking-room/model/dto"

	"github.com/stretchr/testify/mock"
)

type BookingUseCaseMock struct {
	mock.Mock
}

// SendReport implements usecase.BookingUseCase.
func (b *BookingUseCaseMock) SendReport(requestJSON string) ([]model.Booking, error) {
	args := b.Called(requestJSON)
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BookingUseCaseMock) RegisterNewBooking(payload dto.BookingRequestDto, userId string) (model.Booking, error) {
	args := b.Called(payload, userId)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BookingUseCaseMock) FindById(id string, userId string, roleUser string) (model.Booking, error) {
	args := b.Called(id, userId, roleUser)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BookingUseCaseMock) ViewAllBooking() ([]model.Booking, error) {
	args := b.Called()
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BookingUseCaseMock) ViewAllBookingByStatus(status string) ([]model.Booking, error) {
	args := b.Called(status)
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BookingUseCaseMock) UpdateStatusBookAndRoom(id string, approval string) (model.Booking, error) {
	args := b.Called(id, approval)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BookingUseCaseMock) DownloadReport() ([]model.Booking, error) {
	args := b.Called()
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BookingUseCaseMock) GetBookStatus(id string) (string, error) {
	args := b.Called(id)
	return args.String(0), args.Error(1)
}

func (b *BookingUseCaseMock) GetBookingDetailsByBookingID(bookingID string) ([]model.BookingDetail, error) {
	args := b.Called(bookingID)
	return args.Get(0).([]model.BookingDetail), args.Error(1)
}
