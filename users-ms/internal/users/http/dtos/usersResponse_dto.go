package dtos

import (
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/entity"
)

type UsersResDTO struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func (urd *UsersResDTO) MapEntityToDto(userEntity *entity.UsersRes) {
	urd.ID = userEntity.ID
	urd.UserName = userEntity.UserName
	urd.Email = userEntity.Email
}

func (urd *UsersResDTO) MapEntityFromDto() *entity.UsersRes {
	return &entity.UsersRes{
		ID:       urd.ID,
		UserName: urd.UserName,
		Email:    urd.Email,
	}
}
