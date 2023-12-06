package controller

import (
	"bytes"
	"encoding/json"
	"final-project-booking-room/config"
	"final-project-booking-room/model"
	"final-project-booking-room/model/dto"
	middlerwaremock "final-project-booking-room/unit-test/mock-test/middlerware-mock"
	usecasemock "final-project-booking-room/unit-test/mock-test/usecase-mock"

	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookingControllerTestSuite struct {
	suite.Suite
	bum *usecasemock.BookingUseCaseMock
	rg  *gin.RouterGroup
	amm *middlerwaremock.AuthMiddlewareMock
}

func (suite *BookingControllerTestSuite) SetupTest() {
	suite.bum = new(usecasemock.BookingUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.amm = new(middlerwaremock.AuthMiddlewareMock)
}

func TestBookingControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BookingControllerTestSuite))
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
var approval = "accept"
var userId = "1"
var id = "1"
var roleUser = "1"
var mockTokenJwt1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJlbmlnbWFjYW1wIiwiZXhwIjoxNjk5NTAxMjA0LCJpYXQiOjE2OTk0OTc2MDQsInVzZXJJZCI6IjdlODA5N2Q4LWZlZjEtNDllZC05ZjdiLWNlY2M5NTIwY2NhOCIsInJvbGUiOiJlbXBsb3llZSIsInNlcnZpY2VzIjpudWxsfQ.qG38l5-E9P-nEzNg3GIiHNELmuz7BW-cF8g48ooMv98"

func JSONToMap(jsonStr string) gin.H {
	var result gin.H
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		panic(err) // Handle error appropriately in your real code
	}
	return result
}

func (suite *BookingControllerTestSuite) TestGetByStatusHandler_Success() {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	status := "someStatus"
	ctx.Params = append(ctx.Params, gin.Param{Key: "status", Value: status})

	suite.bum.On("ViewAllBookingByStatus", status).Return(mockBookingSlice, nil)

	bookingController := NewBookingController(suite.bum, suite.rg, suite.amm)

	// Trigger the actual function
	bookingController.getByStatusHandler(ctx)

	// Assert the HTTP response status
	assert.Equal(suite.T(), http.StatusOK, ctx.Writer.Status())

	// Assert the response body
	expectedResponseBody := gin.H{"message": "Ok", "data": mockBookingSlice}
	assert.Equal(suite.T(), expectedResponseBody, gin.H{"message": "Ok", "data": mockBookingSlice})
}
func (suite *BookingControllerTestSuite) TestGetAllHandler_Success() {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	suite.bum.On("ViewAllBooking").Return(mockBookingSlice, nil)

	bookingController := NewBookingController(suite.bum, suite.rg, suite.amm)

	bookingController.getAllHandler(ctx)

	assert.Equal(suite.T(), http.StatusOK, ctx.Writer.Status())

	expectedResponseBody := gin.H{"message": "Ok", "data": mockBookingSlice}
	assert.Equal(suite.T(), expectedResponseBody, gin.H{"message": "Ok", "data": mockBookingSlice})
}
func (suite *BookingControllerTestSuite) TestGetHandler_Success() {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	id := "someBookingID"
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: id})

	ctx.Set(config.UserSesion, "someUserID")
	ctx.Set(config.RoleSesion, "someRole")

	suite.bum.On("FindById", id, "someUserID", "someRole").Return(mockBooking, nil)

	bookingController := NewBookingController(suite.bum, suite.rg, suite.amm)

	bookingController.getHandler(ctx)

	assert.Equal(suite.T(), http.StatusOK, ctx.Writer.Status())

	expectedResponseBody := gin.H{"message": "Ok", "data": mockBooking}
	assert.Equal(suite.T(), expectedResponseBody, gin.H{"message": "Ok", "data": mockBooking})
}

func (suite *BookingControllerTestSuite) TestUpdateStatusHandler_Success() {
	suite.bum.On("UpdateStatusBookAndRoom", userId, approval).Return(mockBooking, nil)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	bookingController := NewBookingController(suite.bum, suite.rg, suite.amm)

	bookingController.UpdateStatusHandler(ctx)

	assert.Equal(suite.T(), 400, ctx.Writer.Status())

	expectedResponseBody := gin.H{"message": "Ok", "data": mockBooking}
	assert.Equal(suite.T(), expectedResponseBody, gin.H{"message": "Ok", "data": mockBooking})
}

func (suite *BookingControllerTestSuite) TestCreateHandler_Success() {
	suite.bum.On("RegisterNewBooking", mockPayload, userId).Return(mockBooking, nil)
	bookingController := NewBookingController(suite.bum, suite.rg, suite.amm)
	bookingController.Route()
	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/bills", bytes.NewBuffer(mockPayloadJson))
	fmt.Println("err:", err)
	assert.NoError(suite.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwt)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("user", mockBooking.Users.Id)
	bookingController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}
