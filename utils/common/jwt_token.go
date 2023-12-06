package common

import (
	"errors"
	"final-project-booking-room/config"
	"final-project-booking-room/model"
	"final-project-booking-room/model/dto"
	"final-project-booking-room/utils/modelutil"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken interface {
	GenerateToken(payload model.User) (dto.AuthResponseDto, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
	RefreshToken(oldTokenString string) (dto.AuthResponseDto, error)
	GetUserIdToken(tokenString string) (string, error)
}

type jwtToken struct {
	cfg config.TokenConfig
}

// GenerateToken implements JwtToken.
func (j *jwtToken) GenerateToken(payload model.User) (dto.AuthResponseDto, error) {
	claims := modelutil.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtLifeTime)),
		},
		UserId: payload.Id,
		Role:   payload.Role,
	}

	jwtNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtNewClaims.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, errors.New("failed to generate token")
	}
	return dto.AuthResponseDto{Token: token}, nil
}

// RefreshToken implements JwtToken.
func (j *jwtToken) RefreshToken(oldTokenString string) (dto.AuthResponseDto, error) {
	token, err := jwt.Parse(oldTokenString, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return dto.AuthResponseDto{}, errors.New("invalid claim token")
	}

	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return dto.AuthResponseDto{Token: newTokenString}, nil
}

// VerifyToken implements JwtToken.
func (j *jwtToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, errors.New("failed to verify token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return nil, errors.New("invalid claim token")
	}

	return claims, nil
}

func (j *jwtToken) GetUserIdToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return "", errors.New("failed to verify token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return "", errors.New("invalid claim token")
	}

	userID, ok := claims["UserId"].(string)
	if !ok {
		return "", errors.New("userid not found in claims")
	}

	return userID, nil
}

func NewJwtToken(cfg config.TokenConfig) JwtToken {
	return &jwtToken{cfg: cfg}
}
