package controller

import (
	"encoding/json"
	"final-project-booking-room/config"
	"final-project-booking-room/delivery/middleware"
	"final-project-booking-room/model/dto"
	"final-project-booking-room/usecase"
	"final-project-booking-room/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	uc             usecase.BookingUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (b *BookingController) createHandler(ctx *gin.Context) {
	var payload dto.BookingRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.MustGet(config.UserSesion).(string)
	rspPayload, err := b.uc.RegisterNewBooking(payload, userId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) UpdateStatusHandler(ctx *gin.Context) {
	var payload dto.Approval
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rspPayload, err := b.uc.UpdateStatusBookAndRoom(payload.BookingDetailId, payload.Approval)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Booking ID can't be empty")
		return
	}

	userId := ctx.MustGet(config.UserSesion).(string)
	roleUser := ctx.MustGet(config.RoleSesion).(string)
	rspPayload, err := b.uc.FindById(id, userId, roleUser)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) getByStatusHandler(ctx *gin.Context) {
	status := ctx.Param("status")
	if status == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Status can't be empty")
		return
	}
	rspPayload, err := b.uc.ViewAllBookingByStatus(status)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) getAllHandler(ctx *gin.Context) {
	rspPayload, err := b.uc.ViewAllBooking()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) sendReportHandler(ctx *gin.Context) {
	var requestPayload map[string]string
	if err := ctx.BindJSON(&requestPayload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	toEmail, ok := requestPayload["to"]
	if !ok || toEmail == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid or missing 'to' field in JSON payload")
		return
	}

	requestJSON, err := json.Marshal(map[string]string{"to": toEmail})
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "Error marshaling JSON")
		return
	}
	rspPayload, err := b.uc.SendReport(string(requestJSON))
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) getReportHandler(ctx *gin.Context) {
	rspPayload, err := b.uc.DownloadReport()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) Route() {
	bc := b.rg.Group(config.BookingGroup)
	bc.POST(config.BookingPost, b.authMiddleware.RequireToken("admin", "employee", "GA"), b.createHandler)
	bc.PUT(config.Approval, b.authMiddleware.RequireToken("GA"), b.UpdateStatusHandler)
	bc.GET(config.BookingGetAll, b.authMiddleware.RequireToken("admin", "GA"), b.getAllHandler)
	bc.GET(config.BookingGet, b.authMiddleware.RequireToken("admin", "employee", "GA"), b.getHandler)
	bc.GET(config.BookingGetAllByStatus, b.authMiddleware.RequireToken("admin", "GA"), b.getByStatusHandler)
	bc.GET(config.DownloadReport, b.authMiddleware.RequireToken("admin", "GA"), b.getReportHandler)
	bc.GET(config.SendReport, b.authMiddleware.RequireToken("admin", "GA"), b.sendReportHandler)

}

func NewBookingController(
	uc usecase.BookingUseCase,
	rg *gin.RouterGroup,
	authMiddleware middleware.AuthMiddleware,
) *BookingController {
	return &BookingController{
		uc:             uc,
		rg:             rg,
		authMiddleware: authMiddleware}
}
