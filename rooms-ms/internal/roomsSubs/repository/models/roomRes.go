package models

import "github.com/tesis/internal/roomsSubs/entity"

type RoomRes struct {
	ID          string   `bson:"_id"`
	Name        string   `bson:"name"`
	Description string   `bson:"desc"`
	CreatedBy   string   `bson:"userId"`
	NumMaxUsers string   `bson:"maxUsers"`
	Status      string   `bson:"status"`
	Users       []string `bson:"users"`
}

func (model *RoomRes) MapEntityToModel(roomEntity *entity.Room) {
	model.Name = roomEntity.Name
	model.Description = roomEntity.Description
	model.CreatedBy = roomEntity.CreatedBy
	model.NumMaxUsers = roomEntity.NumMaxUsers
	model.Status = roomEntity.Status
}

func (model *RoomRes) MapEntityFromModel() *entity.RoomRes {
	return &entity.RoomRes{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		CreatedBy:   model.CreatedBy,
		NumMaxUsers: model.NumMaxUsers,
		Status:      string(model.Status),
		Users:       model.Users,
	}
}
