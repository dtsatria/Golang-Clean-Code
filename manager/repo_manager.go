package manager

import "final-project-booking-room/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	RoomRepo() repository.RoomRepository
	BookingRepo() repository.BookingRepository
}

type repoManager struct {
	infra InfraManager
}

// BookingRepo implements RepoManager.
func (r *repoManager) BookingRepo() repository.BookingRepository {
	return repository.NewBookingRepository(r.infra.Conn())
}

// RoomRepo implements RepoManager.
func (r *repoManager) RoomRepo() repository.RoomRepository {
	return repository.NewRoomRepository(r.infra.Conn())
}
func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
