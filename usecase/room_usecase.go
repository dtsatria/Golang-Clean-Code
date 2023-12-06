package usecase

import (
	"final-project-booking-room/model"
	"final-project-booking-room/repository"
	"fmt"
	"strings"
)

type RoomUseCase interface {
	RegisterNewRoom(payload model.Room) (model.Room, error)
	FindById(id string) (model.Room, error)
	FindByRoomType(roomType string) (model.Room, error)
	ViewAllRooms() ([]model.Room, error)
	DeleteById(id string) (model.Room, error)
	UpdateById(id string, payload model.Room) (model.Room, error)
	GetRoomStatus(id string) (string, error)
	GetRoomStatusByBdId(id string) (string, error)
	ChangeRoomStatus(id string) error
	GetAllRoomByStatus(status string) ([]model.Room, error)
}

type roomUseCase struct {
	repo repository.RoomRepository
}

func (r *roomUseCase) GetAllRoomByStatus(status string) ([]model.Room, error) {
	room, err := r.repo.GetAllRoomByStatus(status)

	if err != nil {
		return []model.Room{}, fmt.Errorf("room with status %s not found", status)
	}
	if status != "available" {
		return []model.Room{}, fmt.Errorf("room with status %s not found", status)
	}

	return room, err
}

// ViewAllRooms implements RoomUseCase.
func (r *roomUseCase) ViewAllRooms() ([]model.Room, error) {
	room, err := r.repo.GetAllRoom()
	if err != nil {
		return nil, err
	}

	return room, err
}

// ChangeRoomStatus implements RoomUseCase.
func (r *roomUseCase) ChangeRoomStatus(id string) error {
	err := r.repo.ChangeStatus(id)
	if err != nil {
		return fmt.Errorf("room with id %s not found", id)
	}
	return err
}

// GetRoomStatusByBdId implements RoomUseCase.
func (r *roomUseCase) GetRoomStatusByBdId(id string) (string, error) {
	getStatus, err := r.repo.GetStatusByBd(id)
	if err != nil {
		return "Can't get room status from booking details ID", fmt.Errorf("room with booking details id %s not found", id)
	}
	return getStatus, nil
}

// GetRoomStatus implements RoomUseCase.
func (r *roomUseCase) GetRoomStatus(id string) (string, error) {
	getStatus, err := r.repo.GetStatus(id)
	if err != nil {
		return "Can't get room status", fmt.Errorf("room with id %s not found", id)
	}
	return getStatus, nil
}

// FindByRoomType implements RoomUseCase.
func (r *roomUseCase) FindByRoomType(roomType string) (model.Room, error) {
	findRoom, err := r.repo.GetByRoomType(roomType)
	if err != nil {
		return model.Room{}, fmt.Errorf("room with roomType %s not found", roomType)
	}
	return findRoom, err
}

// UpdateById implements RoomUseCase.
func (r *roomUseCase) UpdateById(id string, payload model.Room) (model.Room, error) {

	updatedRoom, err := r.repo.Get(id)
	if err != nil {
		return model.Room{}, fmt.Errorf("room with id %s not found", id)
	}

	if updatedRoom.Id == id {
		if strings.TrimSpace(payload.RoomType) != "" {
			updatedRoom.RoomType = payload.RoomType
		}
		if payload.MaxCapacity != 0 {
			updatedRoom.MaxCapacity = payload.MaxCapacity
		}
		if strings.TrimSpace(payload.Status) != "" {
			updatedRoom.Status = payload.Status
		}
		if strings.TrimSpace(payload.Facility.RoomDescription) != "" {
			updatedRoom.Facility.RoomDescription = payload.Facility.RoomDescription
		}
		if strings.TrimSpace(payload.Facility.Fwifi) != "" {
			updatedRoom.Facility.Fwifi = payload.Facility.Fwifi
		}
		if strings.TrimSpace(payload.Facility.FsoundSystem) != "" {
			updatedRoom.Facility.FsoundSystem = payload.Facility.FsoundSystem
		}
		if strings.TrimSpace(payload.Facility.Fprojector) != "" {
			updatedRoom.Facility.Fprojector = payload.Facility.Fprojector
		}
		if strings.TrimSpace(payload.Facility.FscreenProjector) != "" {
			updatedRoom.Facility.FscreenProjector = payload.Facility.FscreenProjector
		}
		if strings.TrimSpace(payload.Facility.Fchairs) != "" {
			updatedRoom.Facility.Fchairs = payload.Facility.Fchairs
		}
		if strings.TrimSpace(payload.Facility.Ftables) != "" {
			updatedRoom.Facility.Ftables = payload.Facility.Ftables
		}
		if strings.TrimSpace(payload.Facility.FsoundProof) != "" {
			updatedRoom.Facility.FsoundProof = payload.Facility.FsoundProof
		}
		if strings.TrimSpace(payload.Facility.FsmonkingArea) != "" {
			updatedRoom.Facility.FsmonkingArea = payload.Facility.FsmonkingArea
		}
		if strings.TrimSpace(payload.Facility.Ftelevison) != "" {
			updatedRoom.Facility.Ftelevison = payload.Facility.Ftelevison
		}
		if strings.TrimSpace(payload.Facility.FAc) != "" {
			updatedRoom.Facility.FAc = payload.Facility.FAc
		}
		if strings.TrimSpace(payload.Facility.Fbathroom) != "" {
			updatedRoom.Facility.Fbathroom = payload.Facility.Fbathroom
		}
		if strings.TrimSpace(payload.Facility.FcoffeMaker) != "" {
			updatedRoom.Facility.FcoffeMaker = payload.Facility.FcoffeMaker
		}
	}

	update, err := r.repo.Update(id, updatedRoom)
	if err != nil {
		return model.Room{}, err
	}

	return update, err
}

// DeleteById implements RoomUseCase.
func (r *roomUseCase) DeleteById(id string) (model.Room, error) {
	_, err := r.repo.Delete(id)
	if err != nil {
		return model.Room{}, fmt.Errorf("room with id %s not found", id)
	}

	return model.Room{}, err
}

// FindById implements RoomUseCase.
func (r *roomUseCase) FindById(id string) (model.Room, error) {
	findRoom, err := r.repo.Get(id)
	if err != nil {
		return model.Room{}, fmt.Errorf("room with ID %s not found", id)
	}

	return findRoom, err
}

// RegisterNewRoom implements RoomUseCase.
func (r *roomUseCase) RegisterNewRoom(payload model.Room) (model.Room, error) {
	newRoom, err := r.repo.Create(payload)
	if err != nil {
		panic(err)
	}
	return newRoom, err
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
	return &roomUseCase{repo: repo}
}
