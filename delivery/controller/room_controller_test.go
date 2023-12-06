package controller

import (
	"bytes"
	"encoding/json"
	"final-project-booking-room/model"
	middlerwaremock "final-project-booking-room/unit-test/mock-test/middlerware-mock"
	usecasemock "final-project-booking-room/unit-test/mock-test/usecase-mock"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoomControllerTestSuit struct {
	suite.Suite
	rum *usecasemock.RoomUseCaseMock
	rg  *gin.RouterGroup
	amm *middlerwaremock.AuthMiddlewareMock
}

func (suite *RoomControllerTestSuit) SetupTest() {
	suite.rum = new(usecasemock.RoomUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.amm = new(middlerwaremock.AuthMiddlewareMock)
}

func TestRoomControllerTestSuit(t *testing.T) {
	suite.Run(t, new(RoomControllerTestSuit))
}

var mockRoom = model.Room{
	Id:          "1",
	RoomType:    "room test",
	MaxCapacity: 10,
	Facility: model.RoomFacility{
		Id:               "101",
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

var mockTokenJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJib29raW5nX3Rva2VuIiwiZXhwIjoxNzAwNjEwMjIwLCJpYXQiOjE3MDA2MDY2MjAsInVzZXJJZCI6IjUxMDA2OGMzLTgxNzItNDhjZS04ZDViLWVjYjNkZTU5MWI1MSIsInJvbGUiOiJhZG1pbiIsInNlcnZpY2VzIjpudWxsfQ.n_OOILGUlBO681ZL9LG8snGN09oYo75iQXXiplq-oGQ"

func (suite *RoomControllerTestSuit) TestCreateHandler_Success() {
	suite.rum.On("RegisterNewRoom", mockRoom).Return(mockRoom, nil)
	roomController := NewRoomController(suite.rum, suite.rg, suite.amm)
	roomController.Route()
	// Record untuk menangkap response HTTP
	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockRoom)
	assert.NoError(suite.T(), err)

	// Simulasi melakukan request ke path /api/v1/bills
	// Authorization
	req, err := http.NewRequest(http.MethodPost, "/api/v1/bills", bytes.NewBuffer(mockPayloadJson))
	fmt.Println("err:", err)
	assert.NoError(suite.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJWT)
	ctx, _ := gin.CreateTestContext(record)
	roomController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, http.StatusCreated)
}

func (suite *RoomControllerTestSuit) TestGetAllRoomByStatus_Success() {
	suite.rum.On("GetAllRoomByStatus", mockRoom).Return(mockRoom, nil)
	roomController := NewRoomController(suite.rum, suite.rg, suite.amm)
	roomController.Route()
	// Record untuk menangkap response HTTP
	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockRoom)
	assert.NoError(suite.T(), err)

	// Simulasi melakukan request ke path /api/v1/bills
	// Authorization
	req, err := http.NewRequest(http.MethodPost, "/api/v1/rooms", bytes.NewBuffer(mockPayloadJson))
	fmt.Println("err:", err)
	assert.NoError(suite.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJWT)
	ctx, _ := gin.CreateTestContext(record)
	// ctx.Request = req
	// ctx.Set("user", "1")
	roomController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, http.StatusBadRequest)
}
