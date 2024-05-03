package services

import (
	"fmt"
	"os"

	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/entity"
	"github.com/tesis/internal/rooms/repository"
)

type roomService struct {
	rr repository.RoomRepository
}

type RoomService interface {
	CreateRoom(roomEntity *entity.Room) (string, error)
	GetRooms() ([]entity.RoomRes, error)
	GetRoomByID(id string) (*entity.RoomRes, error)
	UpdateRoom(id string, roomEntity *entity.Room) (*entity.Room, error)
	DeleteRoom(id string) error

	AddUserToRoom(idRoom, idUser string) error
	RemoveUserInRoom(idRoom, idUser string) error
}

func NewRoomService(roomRepository *repository.RoomRepository) RoomService {
	return &roomService{
		rr: *roomRepository,
	}
}

func (rs *roomService) CreateRoom(roomEntity *entity.Room) (string, error) {

	if !utils.CheckRoomStatus(roomEntity.Status) {
		return "", fmt.Errorf("room rtatus is not valid")
	}

	stringID, err := rs.rr.InsertOne(roomEntity)
	if err != nil {
		return "", err
	}
	roomEntity.ID = stringID

	return stringID, nil
}

func (rs *roomService) GetRooms() ([]entity.RoomRes, error) {
	rooms, err := rs.rr.FindAll()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (rs *roomService) GetRoomByID(id string) (*entity.RoomRes, error) {
	room, err := rs.rr.FindOne(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (rs *roomService) UpdateRoom(id string, roomEntity *entity.Room) (*entity.Room, error) {
	roomEntity, err := rs.rr.UpdateOne(id, roomEntity)
	if err != nil {
		return nil, err
	}
	return roomEntity, nil
}

func (rs *roomService) DeleteRoom(id string) error {
	roomDeleted, err := rs.rr.DeleteOne(id)
	if err != nil {
		return err
	}
	if !roomDeleted {
		return fmt.Errorf("room ID was not found")
	}

	return nil
}

// Users in room domain

func (rs *roomService) AddUserToRoom(roomID, userID string) error {
	err := rs.rr.AddUserToRoom(roomID, userID)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(`{"userID": "%s", "roomID": "%s"}`, userID, roomID)
	addUserToRoom(os.Stdout, msg)

	return nil
}

func (rs *roomService) RemoveUserInRoom(roomID, userID string) error {
	err := rs.rr.RemoveUserInRoom(roomID, userID)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(`{"userID": "%s", "roomID": "%s"}`, userID, roomID)
	removeUserToRoom(os.Stdout, msg)
	return nil
}
