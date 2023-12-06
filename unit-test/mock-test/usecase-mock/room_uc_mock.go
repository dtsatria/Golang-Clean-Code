package usecasemock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type RoomUseCaseMock struct {
	mock.Mock
}

// DeleteById implements usecase.RoomUseCase.
func (r *RoomUseCaseMock) DeleteById(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

// FindByRoomType implements usecase.RoomUseCase.
func (r *RoomUseCaseMock) FindByRoomType(roomType string) (model.Room, error) {
	args := r.Called(roomType)
	return args.Get(0).(model.Room), args.Error(1)
}

// GetAllRoomByStatus implements usecase.RoomUseCase.
func (r *RoomUseCaseMock) GetAllRoomByStatus(status string) ([]model.Room, error) {
	args := r.Called(status)
	return args.Get(0).([]model.Room), args.Error(1)
}

// GetRoomStatus implements usecase.RoomUseCase.
func (r *RoomUseCaseMock) GetRoomStatus(id string) (string, error) {
	args := r.Called(id)
	return args.String(0), args.Error(1)
}

// GetRoomStatusByBdId implements usecase.RoomUseCase.
func (r *RoomUseCaseMock) GetRoomStatusByBdId(id string) (string, error) {
	args := r.Called(id)
	return args.String(0), args.Error(1)
}

// RegisterNewRoom implements usecase.RoomUseCase.
func (r *RoomUseCaseMock) RegisterNewRoom(payload model.Room) (model.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(model.Room), args.Error(1)
}

// UpdateById implements usecase.RoomUseCase.
func (r *RoomUseCaseMock) UpdateById(id string, payload model.Room) (model.Room, error) {
	args := r.Called(id, payload)
	return args.Get(0).(model.Room), args.Error(1)
}

// ViewAllRooms implements usecase.RoomUseCase.
func (r *RoomUseCaseMock) ViewAllRooms() ([]model.Room, error) {
	args := r.Called()
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomUseCaseMock) FindById(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomUseCaseMock) ChangeRoomStatus(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
