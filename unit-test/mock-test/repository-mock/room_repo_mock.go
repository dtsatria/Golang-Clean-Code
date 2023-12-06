package repositorymock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type RoomRepositoryMock struct {
	mock.Mock
}

// GetByRoomType implements repository.RoomRepository.
func (r *RoomRepositoryMock) GetByRoomType(roomType string) (model.Room, error) {
	args := r.Called(roomType)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) Create(payload model.Room) (model.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) Get(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) GetAllRoom() ([]model.Room, error) {
	args := r.Called()
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) Delete(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) Update(id string, payload model.Room) (model.Room, error) {
	args := r.Called(id, payload)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) GetStatus(id string) (string, error) {
	args := r.Called(id)
	return args.Get(0).(string), args.Error(1)
}

func (r *RoomRepositoryMock) GetStatusByBd(id string) (string, error) {
	args := r.Called(id)
	return args.Get(0).(string), args.Error(1)
}

func (r *RoomRepositoryMock) ChangeStatus(id string) error {
	args := r.Called(id)
	return args.Error(1)
}

func (r *RoomRepositoryMock) GetAllRoomByStatus(status string) ([]model.Room, error) {
	args := r.Called(status)
	return args.Get(0).([]model.Room), args.Error(1)
}
