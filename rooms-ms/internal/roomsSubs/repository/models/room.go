package models

import "github.com/tesis/internal/roomsSubs/entity"

type Room struct {
	Name        string `bson:"name"`
	Description string `bson:"desc"`
	CreatedBy   string `bson:"userId"`
	NumMaxUsers string `bson:"maxUsers"`
	Status      string `bson:"status"` // active, desactive, deleted
}

func (model *Room) MapEntityToModel(roomEntity *entity.Room) {
	model.Name = roomEntity.Name
	model.Description = roomEntity.Description
	model.CreatedBy = roomEntity.CreatedBy
	model.NumMaxUsers = roomEntity.NumMaxUsers
	model.Status = roomEntity.Status
}

func (model *Room) MapEntityFromModel() *entity.Room {
	return &entity.Room{
		Name:        model.Name,
		Description: model.Description,
		CreatedBy:   model.CreatedBy,
		NumMaxUsers: model.NumMaxUsers,
		Status:      string(model.Status),
	}
}
