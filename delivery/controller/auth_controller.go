package controller

import (
	"final-project-booking-room/config"
	"final-project-booking-room/model"
	"final-project-booking-room/model/dto"
	"final-project-booking-room/usecase"
	"final-project-booking-room/utils/common"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	uc         usecase.AuthUseCase
	rg         *gin.RouterGroup
	jwtService common.JwtToken
}

func (a *AuthController) registerHandler(ctx *gin.Context) {
	var payload model.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newRsp, err := a.uc.Register(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", newRsp)
}

func (a *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rspPayload, err := a.uc.Login(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", rspPayload)
}

func (a *AuthController) refreshTokenHandler(ctx *gin.Context) {
	tokenString := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", -1)
	newToken, err := a.jwtService.RefreshToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	common.SendCreateResponse(ctx, "Ok", newToken)
}

func (a *AuthController) Route() {
	ug := a.rg.Group(config.AuthGroup)
	ug.POST(config.AuthRegister, a.registerHandler)
	ug.POST(config.AuthLogin, a.loginHandler)
	ug.GET(config.AuthRefreshToken, a.refreshTokenHandler)
}

func NewAuthController(uc usecase.AuthUseCase, rg *gin.RouterGroup, jwtService common.JwtToken) *AuthController {
	return &AuthController{uc: uc, rg: rg, jwtService: jwtService}
}
