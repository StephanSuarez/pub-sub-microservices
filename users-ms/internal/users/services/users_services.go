package services

import (
	"errors"
	"log"

	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/entity"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/repository"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	ur repository.UserRepository
}

type UserService interface {
	CreateUser(userEntity *entity.Users) (string, error)
	GetUsers() ([]entity.UsersRes, error)
	GetUserByID(id string) (*entity.UsersRes, error)
	UpdateUser(id string, userEntity *entity.Users) (*entity.UsersRes, error)
	DeleteUser(id string) error
	GetUserByUserName(userName string) (*entity.UsersRes, error)
	GetUserByEmail(email string) (*entity.UsersRes, error)
}

func NewUserService(ur *repository.UserRepository) UserService {
	return &userService{
		ur: *ur,
	}
}

func (us *userService) CreateUser(userEntity *entity.Users) (string, error) {
	hashPass, err := hashPassword(userEntity.Password)
	if err != nil {
		return "", err
	}
	userEntity.Password = hashPass

	userID, err := us.ur.InsertOne(userEntity)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (us *userService) GetUsers() ([]entity.UsersRes, error) {
	users, err := us.ur.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userService) GetUserByID(id string) (*entity.UsersRes, error) {
	user, err := us.ur.FindOne(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) UpdateUser(id string, userEntity *entity.Users) (*entity.UsersRes, error) {
	passwordHashed, err := hashPassword(userEntity.Password)
	if err != nil {
		return nil, err
	}
	userEntity.Password = passwordHashed

	user, err := us.ur.UpdateOne(id, userEntity)
	if err != nil {
		return nil, err
	}
	return user, nil
}

var ErrUserNotFound = errors.New("user ID was not found")

func (us *userService) DeleteUser(id string) error {
	userDeleted, err := us.ur.DeleteOne(id)
	if !userDeleted {
		return ErrUserNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (us *userService) GetUserByUserName(userName string) (*entity.UsersRes, error) {

	log.Println(userName)
	user, err := us.ur.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) GetUserByEmail(email string) (*entity.UsersRes, error) {
	user, err := us.ur.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}
	return user, nil
}
