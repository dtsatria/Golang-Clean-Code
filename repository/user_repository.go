package repository

import (
	"database/sql"
	"errors"
	"final-project-booking-room/model"
	"final-project-booking-room/utils/common"
	"fmt"

	"time"
)

type UserRepository interface {
	GetById(id string) (model.User, error)
	Create(payload model.User) (model.User, error)
	UpdateUserById(userId string, payload model.User) (model.User, error)
	DeleteUserById(id string) (model.User, error)
	GetAllUser() ([]model.User, error)
	GetByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

// MENCARI USER BERDASARKAN EMAIL => UNTUK LOGIN
func (u *userRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.GetByEmail, email).Scan(
		&user.Id,
		&user.Name,
		&user.Divisi,
		&user.Jabatan,
		&user.Email,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return model.User{}, errors.New("email not found")
	}

	return user, err
}

// !MENCARI USER BERDASARKAN ID
func (u *userRepository) GetById(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.GetUserById, id).
		Scan(&user.Id, &user.Name, &user.Divisi, &user.Jabatan, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.User{}, errors.New("id not found")
	}

	return user, nil
}

// !MEMBUAT USER BARU
func (u *userRepository) Create(payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.CreateUser, payload.Name, payload.Divisi, payload.Jabatan,
		payload.Email, payload.Password, payload.Role, time.Now()).
		Scan(&user.Id, &user.Name, &user.Divisi, &user.Jabatan, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.User{}, errors.New("failed to create user")
	}

	return user, nil
}

// !MENGUPDATE USER BERDASARKAN ID
func (u *userRepository) UpdateUserById(userId string, payload model.User) (model.User, error) {

	var user model.User
	fmt.Println(userId)
	err := u.db.QueryRow(common.UpdateUser, payload.Name, payload.Divisi, payload.Jabatan,
		payload.Email, payload.Password, payload.Role, time.Now(), userId).
		Scan(&user.Id, &user.Name, &user.Divisi, &user.Jabatan, &user.Email, &user.Role, &user.UpdatedAt)

	if err != nil {
		return model.User{}, errors.New("failed to update user")
	}

	return user, nil
}

// !MENGHAPUS USER BERDASARKAN ID
func (u *userRepository) DeleteUserById(id string) (model.User, error) {
	_, err := u.db.Exec(common.DeleteUser, id)
	if err != nil {
		return model.User{}, errors.New("failed to delete user")
	}

	return model.User{}, nil
}

// !MENAMPILKAN SEMUA USER
func (u *userRepository) GetAllUser() ([]model.User, error) {
	var users []model.User

	rows, err := u.db.Query(common.GetAllUser)
	if err != nil {
		return nil, errors.New("failed to get all user")
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Divisi, &user.Jabatan, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, errors.New("one of the rows is missing")
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
