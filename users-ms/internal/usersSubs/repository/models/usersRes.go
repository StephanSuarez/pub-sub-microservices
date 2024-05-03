package models

import "github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/entity"

type UsersR struct {
	ID       string `bson:"_id"`
	UserName string `bson:"userName"`
	Email    string `bson:"email"`
}

func (model *UsersR) MapEntityFromModel() *entity.UsersRes {
	return &entity.UsersRes{
		ID:       model.ID,
		UserName: model.UserName,
		Email:    model.Email,
	}
}
