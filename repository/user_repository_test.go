package repository

import (
	"database/sql"
	"errors"
	"final-project-booking-room/model"
	"final-project-booking-room/utils/common"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    UserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(s.T(), err)
	s.mockDB = db
	s.mockSql = mock
	s.repo = NewUserRepository(s.mockDB)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) TestCreateUser_Success() {
	mockUser := model.User{
		Id:        "1",
		Name:      "koko",
		Divisi:    "HR",
		Jabatan:   "Senior",
		Email:     "koko@gmail.com",
		Password:  "12345",
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "divisi", "jabatan", "email", "role", "createdat", "updatedat"}).
		AddRow(mockUser.Id, mockUser.Name, mockUser.Divisi, mockUser.Jabatan, mockUser.Email, mockUser.Role, time.Now(), time.Now())
	s.mockSql.ExpectQuery("INSERT INTO users").WillReturnRows(rows)

	result, err := s.repo.Create(mockUser)

	require.NoError(s.T(), s.mockSql.ExpectationsWereMet())

	require.NoError(s.T(), err, "Unexpected error in Create method")
	require.Equal(s.T(), mockUser.Id, result.Id)
}
func (s *UserRepositoryTestSuite) TestGetAllUser_Success() {

	mockUsers := []model.User{
		{
			Id:        "1",
			Name:      "John Doe",
			Divisi:    "Engineering",
			Jabatan:   "Software Engineer",
			Email:     "john.doe@example.com",
			Role:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	s.mockSql.ExpectQuery(common.GetAllUser).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "divisi", "jabatan", "email", "role", "createdat", "updatedat"}).
			AddRow(mockUsers[0].Id, mockUsers[0].Name, mockUsers[0].Divisi, mockUsers[0].Jabatan, mockUsers[0].Email, mockUsers[0].Role, mockUsers[0].CreatedAt, mockUsers[0].UpdatedAt))

	actualUsers, err := s.repo.GetAllUser()

	assert.NoError(s.T(), err, "Unexpected error in GetAllUser method")
	assert.Equal(s.T(), mockUsers, actualUsers, "Returned users do not match expected users")

	assert.Nil(s.T(), s.mockSql.ExpectationsWereMet(), "Not all SQL expectations were met")
}

func (s *UserRepositoryTestSuite) TestGetAllUser_Fail() {

	s.mockSql.ExpectQuery(common.GetAllUser).
		WillReturnError(errors.New("db error"))

	actualUsers, err := s.repo.GetAllUser()

	expectedErrorMessage := "Expected error in GetAllUser method with 'db error' message"
	assert.Error(s.T(), err, expectedErrorMessage)

	assert.Empty(s.T(), actualUsers, "Expected empty user list for error scenario")

	assert.Nil(s.T(), s.mockSql.ExpectationsWereMet(), "Not all SQL expectations were met")
}

func (s *UserRepositoryTestSuite) TestDeleteUserByID_Fail() {
	userId := "1"
	s.mockSql.ExpectExec("DELETE FROM users WHERE id = ?").WithArgs(userId).
		WillReturnError(errors.New("failed to delete user"))

	_, err := s.repo.DeleteUserById(userId)

	expectedErrorMessage := "failed to delete user"
	assert.EqualError(s.T(), err, expectedErrorMessage, "Unexpected error in DeleteUserById method")

	assert.Nil(s.T(), s.mockSql.ExpectationsWereMet(), "Not all SQL expectations were met")
}

func (s *UserRepositoryTestSuite) TestDeleteUserByID_Success() {
	userID := "1"

	s.mockSql.ExpectExec("DELETE FROM users WHERE id = ?").WithArgs(userID).WillReturnResult(sqlmock.NewResult(0, 1))

	result, err := s.repo.DeleteUserById(userID)

	require.NoError(s.T(), s.mockSql.ExpectationsWereMet())

	require.NoError(s.T(), err, "Unexpected error in DeleteUserById method")
	require.Equal(s.T(), model.User{}, result)
}

func (s *UserRepositoryTestSuite) TestGetUserById_Success() {

	mockUser := model.User{
		Id:        "1",
		Name:      "John Doe",
		Divisi:    "Engineering",
		Jabatan:   "Software Engineer",
		Email:     "john.doe@example.com",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "divisi", "jabatan", "email", "role", "createdat", "updatedat"}).
		AddRow(mockUser.Id, mockUser.Name, mockUser.Divisi, mockUser.Jabatan, mockUser.Email, mockUser.Role, mockUser.CreatedAt, mockUser.UpdatedAt)
	s.mockSql.ExpectQuery("SELECT id,name,divisi,jabatan,email,role,createdat,updatedat FROM users WHERE id = ?").WithArgs(mockUser.Id).WillReturnRows(rows)

	result, err := s.repo.GetById(mockUser.Id)

	require.NoError(s.T(), s.mockSql.ExpectationsWereMet())

	require.NoError(s.T(), err, "Unexpected error in GetById method")
	require.Equal(s.T(), mockUser, result)
}

func (s *UserRepositoryTestSuite) TestGetUserById_Fail() {
	userId := "1"
	s.mockSql.ExpectQuery("SELECT id,name,divisi,jabatan,email,role,createdat,updatedat FROM users WHERE id = ?").
		WithArgs(userId).
		WillReturnError(errors.New("id not found"))

	_, err := s.repo.GetById(userId)

	expectedErrorMessage := "id not found"
	assert.EqualError(s.T(), err, expectedErrorMessage, "Unexpected error in GetById method")

	assert.Nil(s.T(), s.mockSql.ExpectationsWereMet(), "Not all SQL expectations were met")

}
func (s *UserRepositoryTestSuite) TestGetUserByEmail_Success() {

	email := "john.doe@example.com"
	mockUser := model.User{
		Id:       "1",
		Name:     "John Doe",
		Divisi:   "Engineering",
		Jabatan:  "Software Engineer",
		Email:    email,
		Password: "hashed_password",
		Role:     "user",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "divisi", "jabatan", "email", "password", "role"}).
		AddRow(mockUser.Id, mockUser.Name, mockUser.Divisi, mockUser.Jabatan, mockUser.Email, mockUser.Password, mockUser.Role)
	s.mockSql.ExpectQuery("SELECT id, name, divisi, jabatan, email, password, role FROM users WHERE email = ?").WithArgs("john.doe@example.com").WillReturnRows(rows)

	result, err := s.repo.GetByEmail(email)

	require.NoError(s.T(), s.mockSql.ExpectationsWereMet())

	require.NoError(s.T(), err, "Unexpected error in GetByEmail method")
	require.Equal(s.T(), mockUser, result)
}
