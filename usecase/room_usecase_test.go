package usecase

import (
	"final-project-booking-room/model"
	repositorymock "final-project-booking-room/unit-test/mock-test/repository-mock"
	"fmt"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RoomUsecaseTestSuite struct {
	suite.Suite
	rrm *repositorymock.RoomRepositoryMock
	ru  RoomUseCase
}

func (suite *RoomUsecaseTestSuite) SetupTest() {
	suite.rrm = new(repositorymock.RoomRepositoryMock)
	suite.ru = NewRoomUseCase(suite.rrm)
}

func TestRoomUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoomUsecaseTestSuite))
}

var arrayMockRoom []model.Room

var mockRoom = model.Room{
	Id:          "1",
	RoomType:    "test",
	MaxCapacity: 10,
	Facility: model.RoomFacility{
		Id:               "1",
		RoomDescription:  "ruangan test",
		Fwifi:            "ada",
		FsoundSystem:     "ada",
		Fprojector:       "ada",
		FscreenProjector: "ada",
		Fchairs:          "ada",
		Ftables:          "ada",
		FsoundProof:      "ada",
		FsmonkingArea:    "ada",
		Ftelevison:       "ada",
		FAc:              "ada",
		Fbathroom:        "ada",
		FcoffeMaker:      "ada",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	},
	Status:    "available",
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
}

func (suite *RoomUsecaseTestSuite) TestRegisterNewRoom_Success() {
	suite.rrm.On("Create", mockRoom).Return(mockRoom, nil)
	_, err := suite.ru.RegisterNewRoom(mockRoom)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestFindById_Success() {
	suite.rrm.On("Get", mockRoom.Id).Return(mockRoom, nil)
	_, err := suite.ru.FindById(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestFindByRoomType_Success() {
	suite.rrm.On("GetByRoomType", mockRoom.RoomType).Return(mockRoom, nil)
	_, err := suite.ru.FindByRoomType(mockRoom.RoomType)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestViewAllRooms_Success() {
	suite.rrm.On("GetAllRoom").Return(arrayMockRoom, nil)
	_, err := suite.ru.ViewAllRooms()
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestDeleteById_Success() {
	suite.rrm.On("Delete", mockRoom.Id).Return(mockRoom, nil)
	_, err := suite.ru.DeleteById(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

//update

func (suite *RoomUsecaseTestSuite) TestGetRoomStatus_Success() {
	suite.rrm.On("GetStatus", mockRoom.Id).Return(mockRoom.Status, nil)
	_, err := suite.ru.GetRoomStatus(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestGetRoomStatusByBdId_Success() {
	suite.rrm.On("GetStatusByBd", mockRoom.Id).Return(mockRoom.Id, nil)
	_, err := suite.ru.GetRoomStatusByBdId(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestChangeRoomStatus_Success() {
	suite.rrm.On("ChangeStatus", mockRoom.Id).Return(mockRoom.Id, nil)
	err := suite.ru.ChangeRoomStatus(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestGetAllRoomByStatus() {
	suite.rrm.On("GetAllRoomByStatus", mockRoom.Status).Return(arrayMockRoom, nil)
	_, err := suite.ru.GetAllRoomByStatus(mockRoom.Status)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestDeleteById() {
	expectedError := fmt.Errorf("room with id %s not found", mockRoom.Id)

	suite.rrm.On("Delete", mockRoom.Id).Return(model.Room{}, expectedError)

	deletedRoom, err := suite.ru.DeleteById(mockRoom.Id)

	suite.rrm.AssertExpectations(suite.T())

	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Equal(suite.T(), model.Room{}, deletedRoom)
}

func (suite *RoomUsecaseTestSuite) TestFindById() {
	expectedError := fmt.Errorf("room with ID %s not found", mockRoom.Id)
	suite.rrm.On("Get", mockRoom.Id).Return(model.Room{}, expectedError)

	foundRoom, err := suite.ru.FindById(mockRoom.Id)

	suite.rrm.AssertExpectations(suite.T())

	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Equal(suite.T(), model.Room{}, foundRoom)
}

func (suite *RoomUsecaseTestSuite) TestUpdateById_Success() {
	roomId := "1"
	var updatePayload = model.Room{
		Id:          "1",
		RoomType:    "test",
		MaxCapacity: 10,
		Facility: model.RoomFacility{
			Id:               "1",
			RoomDescription:  "ruangan test",
			Fwifi:            "ada",
			FsoundSystem:     "ada",
			Fprojector:       "ada",
			FscreenProjector: "ada",
			Fchairs:          "ada",
			Ftables:          "ada",
			FsoundProof:      "ada",
			FsmonkingArea:    "ada",
			Ftelevison:       "ada",
			FAc:              "ada",
			Fbathroom:        "ada",
			FcoffeMaker:      "ada",
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
		},
		Status:    "available",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	expectedError := fmt.Errorf("room with id %s not found", updatePayload.Id)

	suite.rrm.On("Get", mockRoom.Id).Return(mockRoom, nil)

	suite.rrm.On("Update", roomId, mock.AnythingOfType("model.Room")).Return(model.Room{}, expectedError)

	updatedRoom, err := suite.ru.UpdateById(roomId, updatePayload)

	suite.rrm.AssertExpectations(suite.T())

	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Equal(suite.T(), model.Room{}, updatedRoom)
}
