package models

import "github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/entity"

type Users struct {
	UserName string `bson:"userName"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func (model *Users) MapEntityToModel(userEntity *entity.Users) {
	model.UserName = userEntity.UserName
	model.Email = userEntity.Email
	model.Password = userEntity.Password
}

func (model *Users) MapEntityFromModel() *entity.Users {
	return &entity.Users{
		UserName: model.UserName,
		Email:    model.Email,
		Password: model.Password,
	}
}
