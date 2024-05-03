package dtos

import (
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/entity"
)

type CreateUsersResDTO struct {
	ID       string
	UserName string
	Email    string
}

func (urd *CreateUsersResDTO) MapEntityToDto(userEntity *entity.Users) {
	urd.ID = userEntity.ID
	urd.UserName = userEntity.UserName
	urd.Email = userEntity.Email
}

func (urd *CreateUsersResDTO) MapEntityFromDto() *entity.Users {
	return &entity.Users{
		ID:       urd.ID,
		UserName: urd.UserName,
		Email:    urd.Email,
	}
}
