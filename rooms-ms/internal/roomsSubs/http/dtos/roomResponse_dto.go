package dtos

import "github.com/tesis/internal/roomsSubs/entity"

type RoomResDTO struct {
	ID          string
	Name        string
	Description string
	CreatedBy   string
	NumMaxUsers string
	Status      string
	Users       []string
}

func (rrd *RoomResDTO) MapEntityToDto(roomEntity *entity.RoomRes) {
	rrd.ID = roomEntity.ID
	rrd.Name = roomEntity.Name
	rrd.Description = roomEntity.Description
	rrd.CreatedBy = roomEntity.CreatedBy
	rrd.NumMaxUsers = roomEntity.NumMaxUsers
	rrd.Status = roomEntity.Status
	rrd.Users = roomEntity.Users
}
