package usecasemock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type RoomUsecaseMock struct {
	mock.Mock
}

func (r *RoomUsecaseMock) RegisterNewRoom(payload model.Room) (model.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomUsecaseMock) FindById(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomUsecaseMock) FindByRoomType(roomType string) (model.Room, error) {
	args := r.Called(roomType)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomUsecaseMock) ViewAllRows() ([]model.Room, error) {
	args := r.Called()
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomUsecaseMock) DeleteById(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomUsecaseMock) UpdateById(id string, payload model.Room) (model.Room, error) {
	args := r.Called(id, payload)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomUsecaseMock) GetRoomStatus(id string) (string, error) {
	args := r.Called(id)
	return args.Get(0).(string), args.Error(1)
}

func (r *RoomUsecaseMock) GetRoomStatusByBdId(id string) (string, error) {
	args := r.Called(id)
	return args.Get(0).(string), args.Error(1)
}

func (r *RoomUsecaseMock) ChangeRoomStatus(id string) error {
	args := r.Called(id)
	return args.Error(1)
}

func (r *RoomUsecaseMock) GetAllRoomByStatus(status string) ([]model.Room, error) {
	args := r.Called(status)
	return args.Get(0).([]model.Room), args.Error(1)
}
