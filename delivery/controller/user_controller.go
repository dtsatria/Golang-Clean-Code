package controller

import (
	"final-project-booking-room/config"
	"final-project-booking-room/delivery/middleware"
	"final-project-booking-room/model"
	"final-project-booking-room/usecase"
	"final-project-booking-room/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc             usecase.UserUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (u *UserController) getAllHandler(ctx *gin.Context) {

	rspPayload, err := u.uc.ViewAllUser()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) createHandler(ctx *gin.Context) {
	var payload model.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rspPayload, err := u.uc.RegisterNewUser(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) UpdateUserHandler(ctx *gin.Context) {
	// id := ctx.Param("id")
	// if id == "" {
	// 	common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
	// 	return
	// }

	var payload model.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.MustGet(config.UserSesion).(string)
	rspPayload, err := u.uc.UpdateUserById(userId, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	rspPayload, err := u.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) DeleteByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	_, err := u.uc.DeleteUser(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", nil)
}

func (u *UserController) Route() {
	ur := u.rg.Group(config.UserGroup)
	ur.POST(config.UserPost, u.authMiddleware.RequireToken("admin"), u.createHandler)
	ur.PUT(config.UserUpdate, u.authMiddleware.RequireToken("admin", "employee"), u.UpdateUserHandler)
	ur.GET(config.UserGet, u.authMiddleware.RequireToken("admin"), u.getByIdHandler)
	ur.DELETE(config.UserDelete, u.authMiddleware.RequireToken("admin"), u.DeleteByIdHandler)
	ur.GET(config.UserGetAll, u.authMiddleware.RequireToken("admin"), u.getAllHandler)

}
func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup, authmiddleware middleware.AuthMiddleware) *UserController {
	return &UserController{uc: uc, rg: rg, authMiddleware: authmiddleware}
}
