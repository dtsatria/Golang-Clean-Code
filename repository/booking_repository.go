package repository

import (
	"database/sql"
	"final-project-booking-room/model"
	"final-project-booking-room/utils/common"

	"fmt"
	"time"
)

type BookingRepository interface {
	Create(payload model.Booking, userId string) (model.Booking, error)
	Get(id string, userId string, roleUser string) (model.Booking, error)
	GetAll() ([]model.Booking, error)
	GetAllByStatus(status string) ([]model.Booking, error)
	UpdateStatus(id string, approval string) (model.Booking, error)
	GetBookStatus(id string) (string, error)
	GetBookingDetailsByBookingID(bookingID string) ([]model.BookingDetail, error)
	GetReport(requestJSON string) ([]model.Booking, error)
}

type bookingRepository struct {
	db *sql.DB
}

// GetBookStatus implements BookingRepository.
func (b *bookingRepository) GetBookStatus(id string) (string, error) {
	var status string
	err := b.db.QueryRow("SELECT status FROM booking_details WHERE id = $1", id).Scan(&status)
	if err != nil {
		return "", fmt.Errorf("ID booking details %s is not found", id)
	}
	return status, nil
}

func (b *bookingRepository) GetReport(requestJSON string) ([]model.Booking, error) {
	var result []model.Booking
	_, err := b.db.Exec(common.DownloadReport, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UpdateStatus implements BookingRepository.
func (b *bookingRepository) UpdateStatus(id string, approval string) (model.Booking, error) {
	var booking model.Booking

	// Memulai transaksi
	tx, err := b.db.Begin()
	if err != nil {
		return model.Booking{}, err
	}

	// Update status booking_details
	var bookingId, roomId string
	err = tx.QueryRow(`UPDATE booking_details SET status = $1
	WHERE id = $2 RETURNING bookingid, roomid`, approval, id).Scan(&bookingId, &roomId)
	if err != nil {
		tx.Rollback()
		return model.Booking{}, err
	}

	// Update status rooms based on approval
	status := "available"
	if approval == "accept" {
		status = "booked"
	}
	_, err = tx.Query(`UPDATE rooms SET status = $1 WHERE id = $2`, status, roomId)
	if err != nil {
		tx.Rollback()
		return model.Booking{}, err
	}

	// Commit transaksi
	if err := tx.Commit(); err != nil {
		return model.Booking{}, err
	}

	err = b.db.QueryRow(`
		SELECT b.id, u.id, u.name, u.divisi, u.jabatan, u.email, u.role, u.createdat, u.updatedat, b.createdat, b.updatedat 
		FROM booking b 
		JOIN users u ON u.id = b.userid
		WHERE b.id = $1`, bookingId).Scan(
		&booking.Id,
		&booking.Users.Id,
		&booking.Users.Name,
		&booking.Users.Divisi,
		&booking.Users.Jabatan,
		&booking.Users.Email,
		&booking.Users.Role,
		&booking.Users.CreatedAt,
		&booking.Users.UpdatedAt,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return model.Booking{}, err
	}

	// Menggunakan getBookingDetailsByBookingID untuk mendapatkan data booking details
	bookingDetails, err := b.GetBookingDetailsByBookingID(bookingId)
	if err != nil {
		return model.Booking{}, err
	}

	booking.BookingDetails = bookingDetails

	return booking, nil
}

// GetAllByStatus implements BookingRepository.
func (b *bookingRepository) GetAllByStatus(status string) ([]model.Booking, error) {
	var bookings []model.Booking

	rows, err := b.db.Query(`SELECT b.id, u.id, u.name, u.divisi, u.jabatan, u.email, u.role, u.createdat, u.updatedat, b.createdat, b.updatedat 
	FROM 
	booking b JOIN users u ON u.id = b.userid JOIN booking_details bd ON bd.bookingid = b.id WHERE status = $1`, status)

	if err != nil {
		return nil, fmt.Errorf("can't find data with status : %s", status)
	}

	defer rows.Close()

	for rows.Next() {
		var booking model.Booking
		err := rows.Scan(
			&booking.Id,
			&booking.Users.Id,
			&booking.Users.Name,
			&booking.Users.Divisi,
			&booking.Users.Jabatan,
			&booking.Users.Email,
			&booking.Users.Role,
			&booking.Users.CreatedAt,
			&booking.Users.UpdatedAt,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Ambil data booking_details untuk setiap booking
		bookingDetails, err := b.GetBookingDetailsByBookingID(booking.Id)
		if err != nil {
			return nil, err
		}

		booking.BookingDetails = bookingDetails
		bookings = append(bookings, booking)
	}

	if len(bookings) == 0 {
		return nil, fmt.Errorf("can't find data with status: %s", status)
	}
	return bookings, nil
}

// GetAll implements BookingRepository.
func (b *bookingRepository) GetAll() ([]model.Booking, error) {
	var bookings []model.Booking

	rows, err := b.db.Query(`SELECT b.id, u.id, u.name, u.divisi, u.jabatan, u.email, u.role, u.createdat, u.updatedat, b.createdat, b.updatedat 
	FROM 
	booking b JOIN users u ON u.id = b.userid`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var booking model.Booking
		err := rows.Scan(
			&booking.Id,
			&booking.Users.Id,
			&booking.Users.Name,
			&booking.Users.Divisi,
			&booking.Users.Jabatan,
			&booking.Users.Email,
			&booking.Users.Role,
			&booking.Users.CreatedAt,
			&booking.Users.UpdatedAt,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Ambil data booking_details untuk setiap booking
		bookingDetails, err := b.GetBookingDetailsByBookingID(booking.Id)
		if err != nil {
			return nil, err
		}

		booking.BookingDetails = bookingDetails
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (b *bookingRepository) GetBookingDetailsByBookingID(bookingID string) ([]model.BookingDetail, error) {
	var bookingDetails []model.BookingDetail

	rows, err := b.db.Query(`SELECT bd.id, bd.bookingdate, bd.bookingdateend, bd.status, bd.description, bd.createdat, bd.updatedat, r.id, r.roomtype, r.capacity, r.status, r.createdat, r.updatedat, f.id, f.roomdescription, f.fwifi, f.fsoundsystem, f.fprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat
	FROM 
	booking_details bd JOIN rooms r ON r.id = bd.roomid
	JOIN facilities f ON f.id = r.facilities 
	WHERE bd.bookingid = $1`, bookingID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var bookingDetail model.BookingDetail
		err := rows.Scan(
			&bookingDetail.Id,
			&bookingDetail.BookingDate,
			&bookingDetail.BookingDateEnd,
			&bookingDetail.Status,
			&bookingDetail.Description,
			&bookingDetail.CreatedAt,
			&bookingDetail.UpdatedAt,
			&bookingDetail.Rooms.Id,
			&bookingDetail.Rooms.RoomType,
			&bookingDetail.Rooms.MaxCapacity,
			&bookingDetail.Rooms.Status,
			&bookingDetail.Rooms.CreatedAt,
			&bookingDetail.Rooms.UpdatedAt,
			&bookingDetail.Rooms.Facility.Id,
			&bookingDetail.Rooms.Facility.RoomDescription,
			&bookingDetail.Rooms.Facility.Fwifi,
			&bookingDetail.Rooms.Facility.FsoundSystem,
			&bookingDetail.Rooms.Facility.Fprojector,
			&bookingDetail.Rooms.Facility.Fchairs,
			&bookingDetail.Rooms.Facility.Ftables,
			&bookingDetail.Rooms.Facility.FsoundProof,
			&bookingDetail.Rooms.Facility.FsmonkingArea,
			&bookingDetail.Rooms.Facility.Ftelevison,
			&bookingDetail.Rooms.Facility.FAc,
			&bookingDetail.Rooms.Facility.Fbathroom,
			&bookingDetail.Rooms.Facility.FcoffeMaker,
			&bookingDetail.Rooms.Facility.UpdatedAt,
			&bookingDetail.Rooms.Facility.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		bookingDetails = append(bookingDetails, bookingDetail)
	}

	return bookingDetails, nil
}

// Create implements BookingRepository.
func (b *bookingRepository) Create(payload model.Booking, userId string) (model.Booking, error) {
	tx, err := b.db.Begin()
	fmt.Println(userId)
	if err != nil {
		return model.Booking{}, err
	}

	var booking model.Booking
	err = tx.QueryRow(`INSERT INTO booking (userId, updatedAt) VALUES ($1,$2) RETURNING id,userId,createdAt, updatedAt`, userId, time.Now()).Scan(
		&booking.Id,
		&booking.Users.Id,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return model.Booking{}, tx.Rollback()
	}

	var bookingDetails []model.BookingDetail
	for _, v := range payload.BookingDetails {
		var bookingDetail model.BookingDetail

		// convert booking date end, 3 hari setelah start
		now := time.Now()
		threeHours := 3 * time.Hour
		threeHoursLater := now.Add(threeHours)

		// status awal booking : pending
		bdStatus := "pending"

		err = tx.QueryRow(`INSERT INTO booking_details (bookingid, roomid, bookingdate, bookingdateend, status, description, updatedat) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, bookingid, roomid, bookingdate, bookingdateend, status, description, createdat, updatedat`, booking.Id, v.Rooms.Id, time.Now(), threeHoursLater, bdStatus, v.Description, time.Now()).Scan(
			&bookingDetail.Id,
			&bookingDetail.BookingId,
			&bookingDetail.Rooms.Id,
			&bookingDetail.BookingDate,
			&bookingDetail.BookingDateEnd,
			&bookingDetail.Status,
			&bookingDetail.Description,
			&bookingDetail.CreatedAt,
			&bookingDetail.UpdatedAt,
		)

		if err != nil {
			return model.Booking{}, tx.Rollback()
		}

		bookingDetail.Rooms = v.Rooms
		bookingDetails = append(bookingDetails, bookingDetail)

	}

	booking.Users = payload.Users
	booking.BookingDetails = bookingDetails

	if err := tx.Commit(); err != nil {
		return model.Booking{}, err
	}
	return booking, err
}

// Get implements BookingRepository.
func (b *bookingRepository) Get(id string, userId string, roleUser string) (model.Booking, error) {
	var booking model.Booking
	var isAdminRole string
	if roleUser == "admin" || roleUser == "GA" {
		isAdminRole = "b.id = $1 OR b.id = $2" // Harus begini karena kalo OR b.userid nanti ambil datanya salah tidak sesuai
	} else {
		isAdminRole = "b.id = $1 AND b.userid =$2"
	}
	err := b.db.QueryRow(`
		SELECT b.id, u.id, u.name, u.divisi, u.jabatan, u.email, u.role, u.createdat, u.updatedat, b.createdat, b.updatedat 
		FROM booking b 
		JOIN users u ON u.id = b.userid
		WHERE `+isAdminRole, id, userId).Scan(
		&booking.Id,
		&booking.Users.Id,
		&booking.Users.Name,
		&booking.Users.Divisi,
		&booking.Users.Jabatan,
		&booking.Users.Email,
		&booking.Users.Role,
		&booking.Users.CreatedAt,
		&booking.Users.UpdatedAt,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return model.Booking{}, err
	}

	// Menggunakan getBookingDetailsByBookingID untuk mendapatkan data booking details
	bookingDetails, err := b.GetBookingDetailsByBookingID(id)
	if err != nil {
		return model.Booking{}, err
	}

	booking.BookingDetails = bookingDetails

	return booking, nil
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}
