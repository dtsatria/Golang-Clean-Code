package repositorymock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) GetById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepositoryMock) Create(payload model.User) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepositoryMock) UpdateUserById(userId string, payload model.User) (model.User, error) {
	args := u.Called(userId, payload)
	return args.Get(0).(model.User), args.Error(1)
}
func (u *UserRepositoryMock) DeleteUserById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}
func (u *UserRepositoryMock) GetAllUser() ([]model.User, error) {
	args := u.Called()
	return args.Get(0).([]model.User), args.Error(1)
}
func (u *UserRepositoryMock) GetByEmail(email string) (model.User, error) {
	args := u.Called(email)
	return args.Get(0).(model.User), args.Error(1)
}
