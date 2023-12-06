package usecase

import (
	"final-project-booking-room/model"
	"final-project-booking-room/model/dto"
	"final-project-booking-room/utils/common"
)

type AuthUseCase interface {
	Register(payload model.User) (model.User, error)
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
}

type authUseCase struct {
	uc       UserUseCase
	jwtToken common.JwtToken
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.uc.FindByEmailPassword(payload.Email, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	token, err := a.jwtToken.GenerateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

// Register implements AuthUseCase.
func (a *authUseCase) Register(payload model.User) (model.User, error) {
	return a.uc.RegisterNewUser(payload)

}

func NewAuthUseCase(uc UserUseCase, jwtToken common.JwtToken) AuthUseCase {
	return &authUseCase{uc: uc, jwtToken: jwtToken}
}
