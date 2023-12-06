package usecase

import (
	"errors"
	"final-project-booking-room/model"
	"final-project-booking-room/model/dto"
	repositorymock "final-project-booking-room/unit-test/mock-test/repository-mock"
	usecasemock "final-project-booking-room/unit-test/mock-test/usecase-mock"
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookingUseCaseTestSuite struct {
	suite.Suite
	brm *repositorymock.BookingRepoMock
	uum *usecasemock.UserUseCaseMock
	rum *usecasemock.RoomUseCaseMock
	ues *usecasemock.EmailServiceMock
	bu  BookingUseCase
}

func (suite *BookingUseCaseTestSuite) SetupTest() {
	suite.brm = new(repositorymock.BookingRepoMock)
	suite.uum = new(usecasemock.UserUseCaseMock)
	suite.rum = new(usecasemock.RoomUseCaseMock)
	suite.ues = new(usecasemock.EmailServiceMock)
	suite.bu = NewBookingUseCase(suite.brm, suite.uum, suite.rum, suite.ues)
}

func TestBookingUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BookingUseCaseTestSuite))
}

var mockBooking = model.Booking{
	Id: "1",
	Users: model.User{
		Id:   "1",
		Name: "Saya",
		Role: "admin",
	},
	BookingDetails: []model.BookingDetail{
		{
			Id:        "1",
			BookingId: "1",
			Rooms: model.Room{
				Id:          "1",
				RoomType:    "kamar",
				MaxCapacity: 5,
				Facility: model.RoomFacility{
					Id:              "1",
					RoomDescription: "mantap",
				},
				Status: "",
			},
			Description: "ok",
			Status:      "pending",
		},
	},
}
var mockBookingSlice = []model.Booking{{
	Id: "1",
	Users: model.User{
		Id:   "1",
		Name: "Saya",
		Role: "admin",
	},
	BookingDetails: []model.BookingDetail{
		{
			Id:        "1",
			BookingId: "1",
			Rooms: model.Room{
				Id:          "1",
				RoomType:    "kamar",
				MaxCapacity: 5,
				Facility: model.RoomFacility{
					Id:              "1",
					RoomDescription: "mantap",
				},
				Status: "",
			},
			Description: "ok",
			Status:      "pending",
		},
	},
},
}
var mockPayload = dto.BookingRequestDto{
	Id: "1",
	BoookingDetails: []model.BookingDetail{
		{
			Id:        "1",
			BookingId: "1",
			Rooms: model.Room{
				Id:          "5",
				RoomType:    "kolam",
				MaxCapacity: 20,
				Facility: model.RoomFacility{
					Id:              "1",
					RoomDescription: "mantap",
				},
				Status: "",
			},
			Description: "ok",
			Status:      "pending",
		},
	},
	Description: "ok",
}
var mockUser1 = model.User{
	Id: "1",
}

var mockRoom1 = model.Room{
	Id: "1",
}
var userId = "1"
var id = "1"
var roleUser = "1"

func (suite *BookingUseCaseTestSuite) TestDownloadReport() {
	suite.brm.On("GetAll").Return(mockBookingSlice, nil)

	suite.brm.On("GetBookingDetailsByBookingID", "1").Return(mockBooking.BookingDetails, nil)

	actualBookings, err := suite.bu.DownloadReport()

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), mockBookingSlice, actualBookings)
}

func (suite *BookingUseCaseTestSuite) TestViewBookingAllByStatus_Failed() {
	expectedError := errors.New("some error message")
	suite.brm.On("GetAllByStatus", "someStatus").Return(mockBookingSlice, expectedError)

	actualBookings, err := suite.bu.ViewAllBookingByStatus("someStatus")

	assert.Error(suite.T(), err)

	assert.EqualError(suite.T(), err, fmt.Sprintf("failed to get data error: %v", expectedError))

	assert.Nil(suite.T(), actualBookings)
}
func (suite *BookingUseCaseTestSuite) TestViewBookingAllByStatus_Success() {
	suite.brm.On("GetAllByStatus", mockRoom.Status).Return(mockBookingSlice, nil)

	actualBookings, err := suite.bu.ViewAllBookingByStatus(mockRoom.Status)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), mockBooking, actualBookings[0])
}
func (suite *BookingUseCaseTestSuite) TestViewBookingAll_Failed() {
	expectedError := errors.New("some error message")
	suite.brm.On("GetAll").Return(mockBookingSlice, expectedError)

	actualBookings, err := suite.bu.ViewAllBooking()

	assert.Error(suite.T(), err)

	assert.EqualError(suite.T(), err, fmt.Sprintf("failed to get all bookings: %v", expectedError))

	assert.Nil(suite.T(), actualBookings)
}

func (suite *BookingUseCaseTestSuite) TestViewBookingAll_Success() {
	suite.brm.On("GetAll").Return(mockBookingSlice, nil)

	actualBookings, err := suite.bu.ViewAllBooking()

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), mockBooking, actualBookings[0])
}

func (suite *BookingUseCaseTestSuite) TestFindById_Success() {
	suite.brm.On("Get", id, userId, roleUser).Return(mockBooking, nil)
	actualBooking, err := suite.bu.FindById(id, userId, roleUser)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), mockBooking, actualBooking)
}

func (suite *BookingUseCaseTestSuite) TestFindById_Failed() {
	expectedError := errors.New("some error message")
	suite.brm.On("Get", id, userId, roleUser).Return(mockBooking, expectedError)

	actualBooking, err := suite.bu.FindById(id, userId, roleUser)

	assert.Error(suite.T(), err)

	assert.EqualError(suite.T(), err, fmt.Sprintf("booking with id %s not found", id))

	assert.NotNil(suite.T(), actualBooking)
}

func (suite *BookingUseCaseTestSuite) TestRegisterNewBooking_Succes() {
	suite.uum.On("FindById", userId).Return(mockUser, nil)

	var mockBookingDetails []model.BookingDetail
	for _, v := range mockPayload.BoookingDetails {
		suite.rum.On("FindById", v.Rooms.Id).Return(mockRoom1, nil)

		suite.rum.On("GetRoomStatus", v.Rooms.Id).Return(mockRoom1.Status, nil)

		mockBookingDetails = append(mockBookingDetails, model.BookingDetail{
			Rooms:       mockRoom1,
			Description: v.Description,
			Status:      mockRoom.Status,
		})
	}

	mockNewBookingPayload := model.Booking{
		Users:          mockBooking.Users,
		BookingDetails: mockBookingDetails,
	}

	suite.brm.On("Create", mockNewBookingPayload, userId).Return(mockBooking, nil)
	actual, err := suite.bu.RegisterNewBooking(mockPayload, userId)
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "room status with id 5 is not available")
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), mockBooking.Users.Id, userId)
}
