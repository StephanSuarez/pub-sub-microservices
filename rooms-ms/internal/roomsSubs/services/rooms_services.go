package services

import (
	"github.com/tesis/internal/roomsSubs/repository"
)

type roomSubservice struct {
	rr repository.RoomSubRepository
}

type RoomSubservice interface {
	AddUserToRoom(idRoom, idUser string) error
	RemoveUserInRoom(idRoom, idUser string) error
}

func NewRoomSubService(roomRepository *repository.RoomSubRepository) RoomSubservice {
	return &roomSubservice{
		rr: *roomRepository,
	}
}

// Users in room domain

func (rs *roomSubservice) AddUserToRoom(roomID, userID string) error {
	err := rs.rr.AddUserToRoom(roomID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (rs *roomSubservice) RemoveUserInRoom(roomID, userID string) error {
	err := rs.rr.RemoveUserInRoom(roomID, userID)
	if err != nil {
		return err
	}
	return nil
}
