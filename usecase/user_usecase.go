package usecase

import (
	"errors"
	"final-project-booking-room/model"
	"final-project-booking-room/repository"
	"final-project-booking-room/utils/common"
	"final-project-booking-room/utils/modelutil"

	"fmt"
)

type UserUseCase interface {
	FindById(id string) (model.User, error)
	RegisterNewUser(payload model.User) (model.User, error)
	DeleteUser(id string) (model.User, error)
	ViewAllUser() ([]model.User, error)
	UpdateUserById(userId string, payload model.User) (model.User, error)
	FindByEmailPassword(email string, password string) (model.User, error)
}

type userUseCase struct {
	repo         repository.UserRepository
	emailService common.EmailService
}

func (u *userUseCase) FindByEmailPassword(email string, password string) (model.User, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("user with email %s not found", email)
	}

	if err := common.ComparePasswordHash(user.Password, password); err != nil {
		return model.User{}, fmt.Errorf("compare password %s ", err)
	}

	user.Password = ""

	return user, nil
}

// UpdateUserById implements UserUseCase.
func (u *userUseCase) UpdateUserById(userId string, payload model.User) (model.User, error) {
	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.User{}, err
	}

	payload.Password = newPassword

	user, err := u.repo.UpdateUserById(userId, payload)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get all user : %s", err)
	}
	return user, nil
}

// ViewAllUser implements UserUseCase.
func (u *userUseCase) ViewAllUser() ([]model.User, error) {
	user, err := u.repo.GetAllUser()
	if err != nil {
		return nil, fmt.Errorf("failed to get all user : %s", err)
	}
	return user, nil
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	user, err := u.repo.GetById(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (u *userUseCase) DeleteUser(id string) (model.User, error) {
	_, err := u.repo.DeleteUserById(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return model.User{}, nil
}

func (u *userUseCase) RegisterNewUser(payload model.User) (model.User, error) {
	if !payload.IsValidRole() {
		return model.User{}, errors.New("invalid role, role must be admin or employee")
	}

	if payload.IsEmpty() {
		return model.User{}, errors.New("all fields must be filled in")

		// return model.User{}, errors.New("invalid role, role must admin or employee")
	}

	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.User{}, err
	}

	if payload.Email != "" && payload.Password != "" {
		bodySender := modelutil.BodySender{
			To:      []string{payload.Email},
			Subject: "Registrasi Akun",
			Body:    "Selamat ! Akun anda telah terdaftar. sekarang anda dapat melakukan Booking Room",
		}
		err := u.emailService.SendEmail(bodySender)
		if err != nil {
			return model.User{}, err
		}
	}

	payload.Password = newPassword
	return u.repo.Create(payload)

}

func NewUserUseCase(repo repository.UserRepository, email common.EmailService) UserUseCase {
	return &userUseCase{repo: repo, emailService: email}
}
