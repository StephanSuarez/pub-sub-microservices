package services

import (
	"log"

	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/entity"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/repository"
)

type userSubService struct {
	ur repository.UserSubRepository
}

type UserSubService interface {
	GetUserByUserName(userName string) (*entity.UsersRes, error)
	GetUserByEmail(email string) (*entity.UsersRes, error)

	AddRoomToList(userID, roomID string) error
	RemoveRoomInList(userID, roomID string) error
}

func NewUserSubService(ur *repository.UserSubRepository) UserSubService {
	return &userSubService{
		ur: *ur,
	}
	
}

func (us *userSubService) GetUserByUserName(userName string) (*entity.UsersRes, error) {

	log.Println(userName)
	user, err := us.ur.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userSubService) GetUserByEmail(email string) (*entity.UsersRes, error) {
	user, err := us.ur.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userSubService) AddRoomToList(userID, roomID string) error {
	err := us.ur.AddRoomToList(userID, roomID)
	if err != nil {
		return err
	}
	return nil
}

func (us *userSubService) RemoveRoomInList(userID, roomID string) error {
	err := us.ur.RemoveRoomInList(userID, roomID)
	if err != nil {
		return err
	}
	return nil
}
