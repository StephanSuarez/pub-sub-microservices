package dtos

import (
	"errors"

	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/entity"
)

type UsersRequestDTO struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (urd *UsersRequestDTO) MapEntityToDto(userEntity *entity.Users) {
	urd.UserName = userEntity.UserName
	urd.Email = userEntity.Email
	urd.Password = userEntity.Password
}

func (urd *UsersRequestDTO) MapEntityFromDto() *entity.Users {
	return &entity.Users{
		UserName: urd.UserName,
		Email:    urd.Email,
		Password: urd.Password,
	}
}

func (urd *UsersRequestDTO) ValidateString() error {
	if urd.UserName == "" {
		return errors.New("the field 'userName' can not be empty")
	}

	if urd.Email == "" {
		return errors.New("the field 'email' can not be empty")
	}

	if urd.Password == "" {
		return errors.New("the field 'password' can not be empty")
	}

	return nil
}
