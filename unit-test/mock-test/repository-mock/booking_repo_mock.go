package repositorymock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type BookingRepoMock struct {
	mock.Mock
}

func (b *BookingRepoMock) Create(payload model.Booking, userId string) (model.Booking, error) {
	args := b.Called(payload, userId)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BookingRepoMock) Get(id string, userId string, roleUser string) (model.Booking, error) {
	args := b.Called(id, userId, roleUser)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BookingRepoMock) GetAll() ([]model.Booking, error) {
	args := b.Called()
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BookingRepoMock) GetAllByStatus(status string) ([]model.Booking, error) {
	args := b.Called(status)
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BookingRepoMock) UpdateStatus(id string, approval string) (model.Booking, error) {
	args := b.Called(id, approval)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BookingRepoMock) GetBookStatus(id string) (string, error) {
	args := b.Called(id)
	return args.String(0), args.Error(1)
}

func (b *BookingRepoMock) GetBookingDetailsByBookingID(bookingID string) ([]model.BookingDetail, error) {
	args := b.Called(bookingID)
	return args.Get(0).([]model.BookingDetail), args.Error(1)
}

func (b *BookingRepoMock) GetReport(requestJSON string) ([]model.Booking, error) {
	args := b.Called(requestJSON)
	return args.Get(0).([]model.Booking), args.Error(1)
}
