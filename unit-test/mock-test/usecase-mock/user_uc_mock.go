package usecasemock

import (
	"final-project-booking-room/model"
	"final-project-booking-room/utils/modelutil"

	"github.com/stretchr/testify/mock"
)

type EmailServiceMock struct {
	SendEmailFunc func(payload modelutil.BodySender) error
}
type UserUseCaseMock struct {
	mock.Mock
}

// DeleteUser implements usecase.UserUseCase.
func (u *UserUseCaseMock) DeleteUser(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

// FindByEmailPassword implements usecase.UserUseCase.
func (u *UserUseCaseMock) FindByEmailPassword(email string, password string) (model.User, error) {
	args := u.Called(email, password)
	return args.Get(0).(model.User), args.Error(1)
}

// RegisterNewUser implements usecase.UserUseCase.
func (u *UserUseCaseMock) RegisterNewUser(payload model.User) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

// UpdateUserById implements usecase.UserUseCase.
func (u *UserUseCaseMock) UpdateUserById(userId string, payload model.User) (model.User, error) {
	args := u.Called(userId, payload)
	return args.Get(0).(model.User), args.Error(1)
}

// ViewAllUser implements usecase.UserUseCase.
func (u *UserUseCaseMock) ViewAllUser() ([]model.User, error) {
	args := u.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (u *UserUseCaseMock) FindById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *EmailServiceMock) SendEmail(payload modelutil.BodySender) error {
	if m.SendEmailFunc != nil {
		return m.SendEmailFunc(payload)
	}
	return nil
}

func (m *EmailServiceMock) SendEmailFile(payload modelutil.BodySender) error {
	if m.SendEmailFunc != nil {
		return m.SendEmailFunc(payload)
	}
	return nil
}
