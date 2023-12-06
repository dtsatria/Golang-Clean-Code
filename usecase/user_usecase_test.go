package usecase

import (
	"final-project-booking-room/model"
	repositorymock "final-project-booking-room/unit-test/mock-test/repository-mock"
	usecasemock "final-project-booking-room/unit-test/mock-test/usecase-mock"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	uc  UserUseCase
	urm *repositorymock.UserRepositoryMock
	ues *usecasemock.EmailServiceMock
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.ues = new(usecasemock.EmailServiceMock)
	suite.urm = new(repositorymock.UserRepositoryMock)
	suite.uc = NewUserUseCase(suite.urm, suite.ues)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

var sampleMockUser []model.User

var mockUser = model.User{
	Id:        "1",
	Name:      "test",
	Divisi:    "HR",
	Jabatan:   "Senior",
	Email:     "dika@gmail.com",
	Password:  "12345",
	Role:      "admin",
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
}

func (suite *UserUseCaseTestSuite) TestGetAllUser_Success() {
	mockError := fmt.Errorf("an error occurred")

	suite.urm.On("GetAllUser").Return(sampleMockUser, mockError)

	result, err := suite.uc.ViewAllUser()

	suite.urm.AssertCalled(suite.T(), "GetAllUser")

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to get all user : an error occurred")

	assert.Nil(suite.T(), result)
}

func (s *UserUseCaseTestSuite) TestGetById_Success() {
	s.urm.On("GetById", "1").Return(mockUser, nil)

	resultUser, err := s.uc.FindById("1")

	assert.NoError(s.T(), err, "Unexpected error")
	assert.Equal(s.T(), mockUser, resultUser, "User not as expected")

	s.urm.On("GetById", "nonexistent").Return(model.User{}, fmt.Errorf("user not found"))

	_, err = s.uc.FindById("nonexistent")

	assert.Error(s.T(), err, "Expected error for user not found")
	assert.EqualError(s.T(), err, "user with ID nonexistent not found", "Error message not as expected")

	s.urm.AssertExpectations(s.T())
}

// func (s *UserUseCaseTestSuite) TestDeleteById_Success() {

// 	s.urm.On("DeleteUserById", "1").Return(mockUser, nil)

// 	resultUser, err := s.uc.DeleteUser("1")

// 	assert.NoError(s.T(), err, "Unexpected error")
// 	assert.Equal(s.T(), mockUser, resultUser, "User not as expected")

// 	s.urm.On("DeleteUserById", "nonexistent").Return(model.User{}, fmt.Errorf("user not found"))

// 	_, err = s.uc.DeleteUser("nonexistent")

// 	assert.Error(s.T(), err, "Expected error for user not found")
// 	assert.EqualError(s.T(), err, "user with ID nonexistent not found", "Error message not as expected")

// 	s.urm.AssertExpectations(s.T())
// }
