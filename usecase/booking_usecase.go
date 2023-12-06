package usecase

import (
	"encoding/json"
	"final-project-booking-room/model"
	"final-project-booking-room/model/dto"
	"final-project-booking-room/repository"
	"final-project-booking-room/utils/common"
	"final-project-booking-room/utils/modelutil"
	"os"

	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type BookingUseCase interface {
	RegisterNewBooking(payload dto.BookingRequestDto, userId string) (model.Booking, error)
	FindById(id string, userId string, roleUser string) (model.Booking, error)
	ViewAllBooking() ([]model.Booking, error)
	ViewAllBookingByStatus(status string) ([]model.Booking, error)
	UpdateStatusBookAndRoom(id string, approval string) (model.Booking, error)
	DownloadReport() ([]model.Booking, error)
	SendReport(requestJSON string) ([]model.Booking, error)
}

type EmailRecipient struct {
	To string `json:"to"`
}

type bookingUseCase struct {
	repo         repository.BookingRepository
	userUC       UserUseCase
	roomUC       RoomUseCase
	emailService common.EmailService
}

func (b *bookingUseCase) DownloadReport() ([]model.Booking, error) {
	bookings, err := b.repo.GetAll()
	if err != nil {
		return nil, err
	}

	file, err := os.Create("Report.xlsx")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"

	// Set header row
	header := []string{"ID", "Name", "Divisi", "Jabatan", "Email", "RoomType", "BookingDate", "BookingDateEnd", "Status", "Description"}
	for colIndex, colName := range header {
		cell := fmt.Sprintf("%c%d", 'A'+colIndex, 1)
		xlsx.SetCellValue(sheetName, cell, colName)
	}

	// Write data rows
	for rowIndex, row := range bookings {
		bookingDetails, err := b.repo.GetBookingDetailsByBookingID(row.Id)
		if err != nil {
			return nil, err
		}

		for _, v := range bookingDetails {
			data := []string{
				row.Id,
				row.Users.Name,
				row.Users.Divisi,
				row.Users.Jabatan,
				row.Users.Email,
				v.Rooms.RoomType,
				v.BookingDate.Format("2006-01-02"),
				v.BookingDateEnd.Format("2006-01-02"),
				v.Rooms.Status,
				v.Description,
			}

			for colIndex, cellValue := range data {
				cell := fmt.Sprintf("%c%d", 'A'+colIndex, rowIndex+2)
				xlsx.SetCellValue(sheetName, cell, cellValue)
			}
		}
	}

	// Save the xlsx file
	err = xlsx.SaveAs("Report.xlsx")
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (b *bookingUseCase) SendReport(requestJSON string) ([]model.Booking, error) {
	var emailRecipients EmailRecipient
	err := json.Unmarshal([]byte(requestJSON), &emailRecipients)
	if err != nil {
		return nil, fmt.Errorf("error parsing email JSON: %v", err)
	}

	err = b.emailService.SendEmailFile(modelutil.BodySender{
		To:          []string{emailRecipients.To},
		Subject:     "Report",
		Body:        "Booking Room Report",
		CSVFilePath: "Report.xlsx",
	})

	if err != nil {
		return nil, fmt.Errorf("error sending email: %v", err)
	}

	return nil, nil
}

// UpdateStatusBookAndRoom implements BookingUseCase.
func (b *bookingUseCase) UpdateStatusBookAndRoom(id string, approval string) (model.Booking, error) {
	if approval != "accept" && approval != "decline" {
		return model.Booking{}, fmt.Errorf(`please give approval: "accept" or "decline", not %s`, approval)
	}

	status, err := b.repo.GetBookStatus(id)
	if err != nil {
		return model.Booking{}, err
	}

	if status != "pending" {
		return model.Booking{}, fmt.Errorf("booking status with id %s is already changed (not pending)", id)
	}

	statusRoom, err := b.roomUC.GetRoomStatusByBdId(id)
	if err != nil {
		return model.Booking{}, fmt.Errorf(`sorry, id booking detail %s is not found`, id)
	}

	if statusRoom == "booked" {
		return model.Booking{}, fmt.Errorf("sorry, room is already booked")
	}

	if approval != "accept" && approval != "decline" {
		return model.Booking{}, fmt.Errorf(`please give approval: "accept" or "decline", not %s`, approval)
	}

	booking, err := b.repo.UpdateStatus(id, approval)
	if err != nil {
		return model.Booking{}, fmt.Errorf("booking detail with id %s not found", id)
	}

	return booking, nil
}

// ViewAllBookingByStatus implements BookingUseCase.
func (b *bookingUseCase) ViewAllBookingByStatus(status string) ([]model.Booking, error) {
	bookings, err := b.repo.GetAllByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to get data error: %v", err)
	}
	return bookings, nil
}

// ViewAllBooking implements BookingUseCase.
func (b *bookingUseCase) ViewAllBooking() ([]model.Booking, error) {
	bookings, err := b.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all bookings: %v", err)
	}
	return bookings, nil
}

// FindById implements BookingUseCase.
func (b *bookingUseCase) FindById(id string, userId string, roleUser string) (model.Booking, error) {
	booking, err := b.repo.Get(id, userId, roleUser)
	if err != nil {
		return model.Booking{}, fmt.Errorf("booking with id %s not found", id)
	}
	return booking, nil
}

// RegisterNewBooking implements BookingUseCase.
func (b *bookingUseCase) RegisterNewBooking(payload dto.BookingRequestDto, userId string) (model.Booking, error) {

	user, err := b.userUC.FindById(userId)
	if err != nil {
		return model.Booking{}, fmt.Errorf("user with ID %s not found", userId)
	}

	var bookingDetails []model.BookingDetail
	for _, v := range payload.BoookingDetails {
		room, err := b.roomUC.FindById(v.Rooms.Id)
		if err != nil {
			return model.Booking{}, fmt.Errorf("room with id %s is not found", v.Rooms.Id)
		}

		status, _ := b.roomUC.GetRoomStatus(v.Rooms.Id)
		if status != "available" {
			return model.Booking{}, fmt.Errorf("room status with id %s is not available", v.Rooms.Id)
		}

		bookingDetails = append(bookingDetails, model.BookingDetail{

			Rooms:       room,
			Description: v.Description,
			Status:      v.Status,
		})
	}

	newBookingPayload := model.Booking{
		Users:          user,
		BookingDetails: bookingDetails,
	}

	booking, err := b.repo.Create(newBookingPayload, userId)
	if err != nil {
		return model.Booking{}, err
	}

	return booking, nil
}

func NewBookingUseCase(
	repo repository.BookingRepository,
	userUC UserUseCase,
	roomUC RoomUseCase,
	emailService common.EmailService,
) BookingUseCase {
	return &bookingUseCase{
		repo:         repo,
		userUC:       userUC,
		roomUC:       roomUC,
		emailService: emailService,
	}
}
