package repository

import (
	"database/sql"
	"errors"
	"final-project-booking-room/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookingRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    BookingRepository
}

func (suite *BookingRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = NewBookingRepository(suite.mockDB)
}

func TestBillRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookingRepositoryTestSuite))
}

func (suite *BookingRepositoryTestSuite) TestGetBookStatus_Success() {
	mockBookingDetail := model.BookingDetail{
		Id:        "1",
		BookingId: "1",
		Rooms: model.Room{
			Id:       "1",
			RoomType: "ruang_santai",
			Status:   "available",
		},
		Status: "pending",
	}

	rows := sqlmock.NewRows([]string{"status"}).AddRow(mockBookingDetail.Status)
	suite.mockSql.ExpectQuery("SELECT status FROM booking_details").WillReturnRows(rows)

	actual, err := suite.repo.GetBookStatus(mockBookingDetail.BookingId)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockBookingDetail.Status, actual)
}

func (suite *BookingRepositoryTestSuite) TestGetBookStatus_Fail() {
	mockBookingDetail := model.BookingDetail{
		Id:        "1",
		BookingId: "1",
		Rooms: model.Room{
			Id:       "1",
			RoomType: "ruang_santai",
			Status:   "available",
		},
		Status: "pending",
	}

	suite.mockSql.ExpectQuery("SELECT status FROM booking_details").WillReturnError(errors.New("failed to get status"))

	_, err := suite.repo.GetBookStatus(mockBookingDetail.BookingId)
	assert.Error(suite.T(), err)
}

func (suite *BookingRepositoryTestSuite) TestCreateBooking_Success() {
	mockBooking := model.Booking{
		Id: "1",
		Users: model.User{
			Id:   "1",
			Name: "Siapa",
			Role: "admin",
		},
		BookingDetails: []model.BookingDetail{
			{
				Id:        "1",
				BookingId: "1",
				Rooms: model.Room{
					Id:       "1",
					RoomType: "ruang_santai",
					Status:   "available",
				},
				Status: "pending",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.mockSql.ExpectBegin()

	rows := sqlmock.NewRows([]string{"id", "userId", "createdat", "updatedat"}).AddRow(mockBooking.Id, mockBooking.Users.Id, mockBooking.CreatedAt, mockBooking.UpdatedAt)
	suite.mockSql.ExpectQuery("INSERT INTO booking").WillReturnRows(rows)

	for _, v := range mockBooking.BookingDetails {
		rows := sqlmock.NewRows([]string{"id", "bookingid", "roomid", "bookingdate", "boookingdateend", "status", "description", "created_at", "updated_at"}).AddRow(v.Id, v.BookingId, v.Rooms.Id, v.BookingDate, v.BookingDateEnd, v.Status, v.Description, v.CreatedAt, v.UpdatedAt)

		suite.mockSql.ExpectQuery("INSERT INTO booking_details").WillReturnRows(rows)

		suite.mockSql.ExpectCommit()
		actual, err := suite.repo.Create(mockBooking, mockBooking.Users.Id)
		assert.Nil(suite.T(), err)
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), mockBooking.Id, actual.Id)
	}
}

func (suite *BookingRepositoryTestSuite) TestGetBookingDetailsByBookingID_Success() {

	expectedBookingDetails := model.BookingDetail{
		Id:        "1",
		BookingId: "",
		Rooms: model.Room{
			Id:          "1",
			RoomType:    "kolam",
			MaxCapacity: 5,
			Facility: model.RoomFacility{
				Id:               "1",
				RoomDescription:  "",
				Fwifi:            "",
				FsoundSystem:     "",
				Fprojector:       "",
				FscreenProjector: "",
				Fchairs:          "",
				Ftables:          "",
				FsoundProof:      "",
				FsmonkingArea:    "",
				Ftelevison:       "",
				FAc:              "",
				Fbathroom:        "",
				FcoffeMaker:      "",
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
			Status:    "available",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Description:    "",
		Status:         "pending",
		BookingDate:    time.Now(),
		BookingDateEnd: time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	suite.mockSql.ExpectQuery("^SELECT .*").WithArgs("booking_id_value").WillReturnRows(sqlmock.NewRows(
		[]string{"id", "bookingdate", "bookingdateend", "status", "description", "createdat", "updatedat", "rooms.id", "rooms.roomtype", "rooms.capacity", "rooms.status", "rooms.createdat", "rooms.updatedat", "rooms.facility.id", "rooms.facility.roomdescription", "rooms.facility.fwifi", "rooms.facility.fsoundsystem", "rooms.facility.fprojector", "rooms.facility.fchairs", "rooms.facility.ftables", "rooms.facility.fsoundproof", "rooms.facility.fsmonkingarea", "rooms.facility.ftelevison", "rooms.facility.fac", "rooms.facility.fbathroom", "rooms.facility.fcoffemaker", "rooms.facility.createdat", "rooms.facility.updatedat"},
	).AddRow(
		expectedBookingDetails.Id,
		expectedBookingDetails.BookingDate,
		expectedBookingDetails.BookingDateEnd,
		expectedBookingDetails.Status,
		expectedBookingDetails.Description,
		expectedBookingDetails.CreatedAt,
		expectedBookingDetails.UpdatedAt,
		expectedBookingDetails.Rooms.Id,
		expectedBookingDetails.Rooms.RoomType,
		expectedBookingDetails.Rooms.MaxCapacity,
		expectedBookingDetails.Rooms.Status,
		expectedBookingDetails.Rooms.CreatedAt,
		expectedBookingDetails.Rooms.UpdatedAt,
		expectedBookingDetails.Rooms.Facility.Id,
		expectedBookingDetails.Rooms.Facility.RoomDescription,
		expectedBookingDetails.Rooms.Facility.Fwifi,
		expectedBookingDetails.Rooms.Facility.FsoundSystem,
		expectedBookingDetails.Rooms.Facility.Fprojector,
		expectedBookingDetails.Rooms.Facility.Fchairs,
		expectedBookingDetails.Rooms.Facility.Ftables,
		expectedBookingDetails.Rooms.Facility.FsoundProof,
		expectedBookingDetails.Rooms.Facility.FsmonkingArea,
		expectedBookingDetails.Rooms.Facility.Ftelevison,
		expectedBookingDetails.Rooms.Facility.FAc,
		expectedBookingDetails.Rooms.Facility.Fbathroom,
		expectedBookingDetails.Rooms.Facility.FcoffeMaker,
		expectedBookingDetails.Rooms.Facility.UpdatedAt,
		expectedBookingDetails.Rooms.Facility.CreatedAt,
	))

	results, err := suite.repo.GetBookingDetailsByBookingID("booking_id_value")

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), results, 1)

}

func (suite *BookingRepositoryTestSuite) TestGet_Success() {
	mockBooking := model.Booking{
		Id: "1",
		Users: model.User{
			Id:   "1",
			Name: "Siapa",
			Role: "admin",
		},
		BookingDetails: []model.BookingDetail{
			{
				Id:        "1",
				BookingId: "1",
				Rooms: model.Room{
					Id:       "1",
					RoomType: "ruang_santai",
					Status:   "available",
				},
				Status: "pending",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.mockSql.ExpectQuery("^SELECT .*").WithArgs("booking_id_value", "userId").WillReturnRows(sqlmock.NewRows(
		[]string{"id", "users.id", "users.name", "users.divisi", "users.jabatan", "users.email", "users.role", "users.createdat", "users.updatedat", "createdat", "updatedat"},
	).AddRow(
		mockBooking.Id,
		mockBooking.Users.Id,
		mockBooking.Users.Name,
		mockBooking.Users.Divisi,
		mockBooking.Users.Jabatan,
		mockBooking.Users.Email,
		mockBooking.Users.Role,
		mockBooking.Users.CreatedAt,
		mockBooking.Users.UpdatedAt,
		mockBooking.CreatedAt,
		mockBooking.UpdatedAt,
	))

	for _, v := range mockBooking.BookingDetails {
		suite.mockSql.ExpectQuery("^SELECT .*").WithArgs("booking_id_value").WillReturnRows(sqlmock.NewRows(
			[]string{"id", "bookingdate", "bookingdateend", "status", "description", "createdat", "updatedat", "rooms.id", "rooms.roomtype", "rooms.capacity", "rooms.status", "rooms.createdat", "rooms.updatedat", "rooms.facility.id", "rooms.facility.roomdescription", "rooms.facility.fwifi", "rooms.facility.fsoundsystem", "rooms.facility.fprojector", "rooms.facility.fchairs", "rooms.facility.ftables", "rooms.facility.fsoundproof", "rooms.facility.fsmonkingarea", "rooms.facility.ftelevison", "rooms.facility.fac", "rooms.facility.fbathroom", "rooms.facility.fcoffemaker", "rooms.facility.createdat", "rooms.facility.updatedat"},
		).AddRow(
			v.Id,
			v.BookingDate,
			v.BookingDateEnd,
			v.Status,
			v.Description,
			v.CreatedAt,
			v.UpdatedAt,
			v.Rooms.Id,
			v.Rooms.RoomType,
			v.Rooms.MaxCapacity,
			v.Rooms.Status,
			v.Rooms.CreatedAt,
			v.Rooms.UpdatedAt,
			v.Rooms.Facility.Id,
			v.Rooms.Facility.RoomDescription,
			v.Rooms.Facility.Fwifi,
			v.Rooms.Facility.FsoundSystem,
			v.Rooms.Facility.Fprojector,
			v.Rooms.Facility.Fchairs,
			v.Rooms.Facility.Ftables,
			v.Rooms.Facility.FsoundProof,
			v.Rooms.Facility.FsmonkingArea,
			v.Rooms.Facility.Ftelevison,
			v.Rooms.Facility.FAc,
			v.Rooms.Facility.Fbathroom,
			v.Rooms.Facility.FcoffeMaker,
			v.Rooms.Facility.UpdatedAt,
			v.Rooms.Facility.CreatedAt,
		))
	}
	result, err := suite.repo.Get("booking_id_value", "userId", "admin")

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockBooking.Id, result.Id)
}

func (suite *BookingRepositoryTestSuite) TestGetAll_Success() {
	mockBooking := []model.Booking{
		{Id: "1",
			Users: model.User{
				Id:   "1",
				Name: "Siapa",
				Role: "admin",
			},
			BookingDetails: []model.BookingDetail{
				{
					Id:        "1",
					BookingId: "1",
					Rooms: model.Room{
						Id:       "1",
						RoomType: "ruang_santai",
						Status:   "available",
					},
					Status: "pending",
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, x := range mockBooking {
		// Konfigurasi mock query
		suite.mockSql.ExpectQuery("^SELECT .*").WillReturnRows(sqlmock.NewRows(
			[]string{"id", "users.id", "users.name", "users.divisi", "users.jabatan", "users.email", "users.role", "users.createdat", "users.updatedat", "createdat", "updatedat"},
		).AddRow(
			x.Id,
			x.Users.Id,
			x.Users.Name,
			x.Users.Divisi,
			x.Users.Jabatan,
			x.Users.Email,
			x.Users.Role,
			x.Users.CreatedAt,
			x.Users.UpdatedAt,
			x.CreatedAt,
			x.UpdatedAt,
		))

		for _, v := range x.BookingDetails {
			suite.mockSql.ExpectQuery("^SELECT .*").WithArgs(v.BookingId).WillReturnRows(sqlmock.NewRows(
				[]string{"id", "bookingdate", "bookingdateend", "status", "description", "createdat", "updatedat", "rooms.id", "rooms.roomtype", "rooms.capacity", "rooms.status", "rooms.createdat", "rooms.updatedat", "rooms.facility.id", "rooms.facility.roomdescription", "rooms.facility.fwifi", "rooms.facility.fsoundsystem", "rooms.facility.fprojector", "rooms.facility.fchairs", "rooms.facility.ftables", "rooms.facility.fsoundproof", "rooms.facility.fsmonkingarea", "rooms.facility.ftelevison", "rooms.facility.fac", "rooms.facility.fbathroom", "rooms.facility.fcoffemaker", "rooms.facility.createdat", "rooms.facility.updatedat"},
			).AddRow(
				v.Id,
				v.BookingDate,
				v.BookingDateEnd,
				v.Status,
				v.Description,
				v.CreatedAt,
				v.UpdatedAt,
				v.Rooms.Id,
				v.Rooms.RoomType,
				v.Rooms.MaxCapacity,
				v.Rooms.Status,
				v.Rooms.CreatedAt,
				v.Rooms.UpdatedAt,
				v.Rooms.Facility.Id,
				v.Rooms.Facility.RoomDescription,
				v.Rooms.Facility.Fwifi,
				v.Rooms.Facility.FsoundSystem,
				v.Rooms.Facility.Fprojector,
				v.Rooms.Facility.Fchairs,
				v.Rooms.Facility.Ftables,
				v.Rooms.Facility.FsoundProof,
				v.Rooms.Facility.FsmonkingArea,
				v.Rooms.Facility.Ftelevison,
				v.Rooms.Facility.FAc,
				v.Rooms.Facility.Fbathroom,
				v.Rooms.Facility.FcoffeMaker,
				v.Rooms.Facility.UpdatedAt,
				v.Rooms.Facility.CreatedAt,
			))
		}
	}
	result, err := suite.repo.GetAll()

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), mockBooking[0].Id, result[0].Id)
}

func (suite *BookingRepositoryTestSuite) TestGetAllByStatus_Success() {
	mockBooking := []model.Booking{
		{Id: "1",
			Users: model.User{
				Id:   "1",
				Name: "Siapa",
				Role: "admin",
			},
			BookingDetails: []model.BookingDetail{
				{
					Id:        "1",
					BookingId: "1",
					Rooms: model.Room{
						Id:       "1",
						RoomType: "ruang_santai",
						Status:   "available",
					},
					Status: "pending",
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, x := range mockBooking {
		suite.mockSql.ExpectQuery("^SELECT .*").WithArgs("status").WillReturnRows(sqlmock.NewRows(
			[]string{"id", "users.id", "users.name", "users.divisi", "users.jabatan", "users.email", "users.role", "users.createdat", "users.updatedat", "createdat", "updatedat"},
		).AddRow(
			x.Id,
			x.Users.Id,
			x.Users.Name,
			x.Users.Divisi,
			x.Users.Jabatan,
			x.Users.Email,
			x.Users.Role,
			x.Users.CreatedAt,
			x.Users.UpdatedAt,
			x.CreatedAt,
			x.UpdatedAt,
		))

		for _, v := range x.BookingDetails {
			suite.mockSql.ExpectQuery("^SELECT .*").WithArgs(v.BookingId).WillReturnRows(sqlmock.NewRows(
				[]string{"id", "bookingdate", "bookingdateend", "status", "description", "createdat", "updatedat", "rooms.id", "rooms.roomtype", "rooms.capacity", "rooms.status", "rooms.createdat", "rooms.updatedat", "rooms.facility.id", "rooms.facility.roomdescription", "rooms.facility.fwifi", "rooms.facility.fsoundsystem", "rooms.facility.fprojector", "rooms.facility.fchairs", "rooms.facility.ftables", "rooms.facility.fsoundproof", "rooms.facility.fsmonkingarea", "rooms.facility.ftelevison", "rooms.facility.fac", "rooms.facility.fbathroom", "rooms.facility.fcoffemaker", "rooms.facility.createdat", "rooms.facility.updatedat"},
			).AddRow(
				v.Id,
				v.BookingDate,
				v.BookingDateEnd,
				v.Status,
				v.Description,
				v.CreatedAt,
				v.UpdatedAt,
				v.Rooms.Id,
				v.Rooms.RoomType,
				v.Rooms.MaxCapacity,
				v.Rooms.Status,
				v.Rooms.CreatedAt,
				v.Rooms.UpdatedAt,
				v.Rooms.Facility.Id,
				v.Rooms.Facility.RoomDescription,
				v.Rooms.Facility.Fwifi,
				v.Rooms.Facility.FsoundSystem,
				v.Rooms.Facility.Fprojector,
				v.Rooms.Facility.Fchairs,
				v.Rooms.Facility.Ftables,
				v.Rooms.Facility.FsoundProof,
				v.Rooms.Facility.FsmonkingArea,
				v.Rooms.Facility.Ftelevison,
				v.Rooms.Facility.FAc,
				v.Rooms.Facility.Fbathroom,
				v.Rooms.Facility.FcoffeMaker,
				v.Rooms.Facility.UpdatedAt,
				v.Rooms.Facility.CreatedAt,
			))
		}
	}
	result, err := suite.repo.GetAllByStatus("status")

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), mockBooking[0].Id, result[0].Id)
}
