package controller

import (
	"final-project-booking-room/model"
	middlerwaremock "final-project-booking-room/unit-test/mock-test/middlerware-mock"
	usecasemock "final-project-booking-room/unit-test/mock-test/usecase-mock"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	ucm *usecasemock.UserUseCaseMock
	rg  *gin.RouterGroup
	amm *middlerwaremock.AuthMiddlewareMock
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.ucm = new(usecasemock.UserUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.amm = new(middlerwaremock.AuthMiddlewareMock)
}

func TestUSerControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

var mockUser = model.User{
	Id:        "1",
	Name:      "dika",
	Divisi:    "HR",
	Jabatan:   "Senior",
	Email:     "dika@gmail.com",
	Password:  "12345",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockTokenJwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJlbmlnbWFjYW1wIiwiZXhwIjoxNjk5NTAxMjA0LCJpYXQiOjE2OTk0OTc2MDQsInVzZXJJZCI6IjdlODA5N2Q4LWZlZjEtNDllZC05ZjdiLWNlY2M5NTIwY2NhOCIsInJvbGUiOiJlbXBsb3llZSIsInNlcnZpY2VzIjpudWxsfQ.qG38l5-E9P-nEzNg3GIiHNELmuz7BW-cF8g48ooMv98"
