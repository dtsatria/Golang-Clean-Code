package repository

import (
	"database/sql"
	"final-project-booking-room/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoomRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    RoomRepository
}

func (suite *RoomRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = NewRoomRepository(suite.mockDB)
}

func TestRoomRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RoomRepositoryTestSuite))
}

func (suite *RoomRepositoryTestSuite) TestCreateRoom_Success() {
	mockRoom := model.Room{
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

	rows := sqlmock.NewRows([]string{"id", "roomdescription", "fwifi", "fsoundsystem", "fprojector", "fscreenprojector", "fchairs", "ftables", "fsoundproof", "fsmonkingarea", "ftelevison", "fac", "fbathroom", "fcoffemaker", "createdat", "updatedat"}).AddRow(
		mockRoom.Facility.Id,
		mockRoom.Facility.RoomDescription,
		mockRoom.Facility.Fwifi,
		mockRoom.Facility.FsoundSystem,
		mockRoom.Facility.Fprojector,
		mockRoom.Facility.FscreenProjector,
		mockRoom.Facility.Fchairs,
		mockRoom.Facility.Ftables,
		mockRoom.Facility.FsoundProof,
		mockRoom.Facility.FsmonkingArea,
		mockRoom.Facility.Ftelevison,
		mockRoom.Facility.FAc,
		mockRoom.Facility.Fbathroom,
		mockRoom.Facility.FcoffeMaker,
		mockRoom.Facility.CreatedAt,
		mockRoom.Facility.UpdatedAt)
	suite.mockSql.ExpectQuery("INSERT INTO facilities").WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"id", "roomtype", "capacity", "status", "createdat", "updatedat"}).AddRow(
		mockRoom.Id,
		mockRoom.RoomType,
		mockRoom.MaxCapacity,
		mockRoom.Status,
		mockRoom.CreatedAt,
		mockRoom.UpdatedAt)
	suite.mockSql.ExpectQuery("INSERT INTO rooms").WillReturnRows(rows)

	suite.mockSql.ExpectCommit()
	actual, err := suite.repo.Create(mockRoom)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockRoom.Id, actual.Id)

}

func (suite *RoomRepositoryTestSuite) TestGetByRoomType() {
	mockRoom := model.Room{
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

	rows := sqlmock.NewRows([]string{
		"id", "roomtype", "capacity",
		"facilityid", "roomdescription", "fwifi", "fsoundsystem", "fprojector", "fscreenprojector", "fchairs", "ftables", "fsoundproof", "fsmonkingarea", "ftelevison", "fac", "fbathroom", "fcoffemaker", "updatedat", "createdat",
		"status", "createdat", "updatedat"}).AddRow(
		mockRoom.Id,
		mockRoom.RoomType,
		mockRoom.MaxCapacity,
		mockRoom.Facility.Id,
		mockRoom.Facility.RoomDescription,
		mockRoom.Facility.Fwifi,
		mockRoom.Facility.FsoundSystem,
		mockRoom.Facility.Fprojector,
		mockRoom.Facility.FscreenProjector,
		mockRoom.Facility.Fchairs,
		mockRoom.Facility.Ftables,
		mockRoom.Facility.FsoundProof,
		mockRoom.Facility.FsmonkingArea,
		mockRoom.Facility.Ftelevison,
		mockRoom.Facility.FAc,
		mockRoom.Facility.Fbathroom,
		mockRoom.Facility.FcoffeMaker,
		mockRoom.Facility.UpdatedAt,
		mockRoom.Facility.CreatedAt,
		mockRoom.Status,
		mockRoom.CreatedAt,
		mockRoom.UpdatedAt,
	)

	suite.mockSql.ExpectQuery(`SELECT r.id, r.roomtype, r.capacity, f.id, f.roomdescription, f.fwifi, f.fsoundsystem, f.fprojector, f.fscreenprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat, r.status, r.createdat, r.updatedat FROM rooms AS r JOIN facilities AS f ON f.id = r.facilities WHERE r.roomtype = \$1;`).
		WithArgs("room test").
		WillReturnRows(rows)

	result, err := suite.repo.GetByRoomType("room test")

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockRoom, result)

}

func (suite *RoomRepositoryTestSuite) TestGet_Success() {
	mockRoom := model.Room{
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

	rows := sqlmock.NewRows([]string{
		"id", "roomtype", "capacity",
		"facilityid", "roomdescription", "fwifi", "fsoundsystem", "fprojector", "fscreenprojector", "fchairs", "ftables", "fsoundproof", "fsmonkingarea", "ftelevison", "fac", "fbathroom", "fcoffemaker", "updatedat", "createdat",
		"status", "createdat", "updatedat",
	}).
		AddRow(
			mockRoom.Id,
			mockRoom.RoomType,
			mockRoom.MaxCapacity,
			mockRoom.Facility.Id,
			mockRoom.Facility.RoomDescription,
			mockRoom.Facility.Fwifi,
			mockRoom.Facility.FsoundSystem,
			mockRoom.Facility.Fprojector,
			mockRoom.Facility.FscreenProjector,
			mockRoom.Facility.Fchairs,
			mockRoom.Facility.Ftables,
			mockRoom.Facility.FsoundProof,
			mockRoom.Facility.FsmonkingArea,
			mockRoom.Facility.Ftelevison,
			mockRoom.Facility.FAc,
			mockRoom.Facility.Fbathroom,
			mockRoom.Facility.FcoffeMaker,
			mockRoom.Facility.UpdatedAt,
			mockRoom.Facility.CreatedAt,
			mockRoom.Status,
			mockRoom.CreatedAt,
			mockRoom.UpdatedAt,
		)

	suite.mockSql.ExpectQuery(`SELECT r.id, r.roomtype, r.capacity, f.id, f.roomdescription, f.fwifi, f.fsoundsystem, f.fprojector, f.fscreenprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat, r.status, r.createdat, r.updatedat FROM rooms AS r JOIN facilities AS f ON f.id = r.facilities WHERE r.id = \$1;`).
		WithArgs("1").
		WillReturnRows(rows)

	result, err := suite.repo.Get("1")

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockRoom, result)
}

func (suite *RoomRepositoryTestSuite) TestGetAllRoomByStatus_Success() {
	mockRoom := []model.Room{
		{
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
		},
	}
	rows := sqlmock.NewRows([]string{
		"id", "roomtype", "capacity",
		"facilityid", "roomdescription", "fwifi", "fsoundsystem", "fprojector", "fscreenprojector", "fchairs", "ftables", "fsoundproof", "fsmonkingarea", "ftelevison", "fac", "fbathroom", "fcoffemaker", "updatedat", "createdat",
		"status", "createdat", "updatedat",
	}).
		AddRow(
			mockRoom[0].Id,
			mockRoom[0].RoomType,
			mockRoom[0].MaxCapacity,
			mockRoom[0].Facility.Id,
			mockRoom[0].Facility.RoomDescription,
			mockRoom[0].Facility.Fwifi,
			mockRoom[0].Facility.FsoundSystem,
			mockRoom[0].Facility.Fprojector,
			mockRoom[0].Facility.FscreenProjector,
			mockRoom[0].Facility.Fchairs,
			mockRoom[0].Facility.Ftables,
			mockRoom[0].Facility.FsoundProof,
			mockRoom[0].Facility.FsmonkingArea,
			mockRoom[0].Facility.Ftelevison,
			mockRoom[0].Facility.FAc,
			mockRoom[0].Facility.Fbathroom,
			mockRoom[0].Facility.FcoffeMaker,
			mockRoom[0].Facility.UpdatedAt,
			mockRoom[0].Facility.CreatedAt,
			mockRoom[0].Status,
			mockRoom[0].CreatedAt,
			mockRoom[0].UpdatedAt,
		)

	suite.mockSql.ExpectQuery(`SELECT r.id, r.roomtype, r.capacity, f.id, f.roomdescription, f.fwifi, f.fsoundsystem, f.fprojector, f.fscreenprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat, r.status, r.createdat, r.updatedat FROM rooms AS r JOIN facilities AS f ON f.id = r.facilities WHERE r.status = \$1`).
		WithArgs("available").
		WillReturnRows(rows)

	result, err := suite.repo.GetAllRoomByStatus("available")

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockRoom, result)
}

func (suite *RoomRepositoryTestSuite) TestGetAllRoom() {
	mockRoom := []model.Room{
		{
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
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "roomtype", "capacity",
		"facilityid", "roomdescription", "fwifi", "fsoundsystem", "fprojector", "fscreenprojector", "fchairs", "ftables", "fsoundproof", "fsmonkingarea", "ftelevison", "fac", "fbathroom", "fcoffemaker", "updatedat", "createdat",
		"status", "createdat", "updatedat",
	}).
		AddRow(
			mockRoom[0].Id,
			mockRoom[0].RoomType,
			mockRoom[0].MaxCapacity,
			mockRoom[0].Facility.Id,
			mockRoom[0].Facility.RoomDescription,
			mockRoom[0].Facility.Fwifi,
			mockRoom[0].Facility.FsoundSystem,
			mockRoom[0].Facility.Fprojector,
			mockRoom[0].Facility.FscreenProjector,
			mockRoom[0].Facility.Fchairs,
			mockRoom[0].Facility.Ftables,
			mockRoom[0].Facility.FsoundProof,
			mockRoom[0].Facility.FsmonkingArea,
			mockRoom[0].Facility.Ftelevison,
			mockRoom[0].Facility.FAc,
			mockRoom[0].Facility.Fbathroom,
			mockRoom[0].Facility.FcoffeMaker,
			mockRoom[0].Facility.UpdatedAt,
			mockRoom[0].Facility.CreatedAt,
			mockRoom[0].Status,
			mockRoom[0].CreatedAt,
			mockRoom[0].UpdatedAt,
		)

	suite.mockSql.ExpectQuery(`SELECT r.id, r.roomtype, r.capacity, f.id, f.roomdescription, f.fwifi, f.fsoundsystem, f.fprojector, f.fscreenprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat, r.status, r.createdat, r.updatedat FROM rooms AS r JOIN facilities AS f ON f.id = r.facilities`).
		WillReturnRows(rows)

	result, err := suite.repo.GetAllRoom()

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockRoom, result)
}

func (suite *RoomRepositoryTestSuite) TestChangeStatus_Success() {
	roomId := "110"

	suite.mockSql.ExpectExec(`UPDATE rooms SET status = \$1 WHERE id = \$2`).
		WithArgs("available", roomId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.repo.ChangeStatus(roomId)

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), roomId, roomId)
}
