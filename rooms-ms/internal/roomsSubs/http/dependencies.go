package http

import (
	"github.com/tesis/internal/roomsSubs/repository"
	"github.com/tesis/internal/roomsSubs/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomSubDependencies struct {
	rr repository.RoomSubRepository
	rs services.RoomSubservice
	rh RoomSubHandler
}

func NewRoomSubDependencies(dbInstanceConn *mongo.Database) *RoomSubDependencies {
	dbInstance := dbInstanceConn
	roomSubRepository := repository.NewRoomSubRepository(dbInstance)
	roomSubService := services.NewRoomSubService(&roomSubRepository)
	roomSubHandler := NewRoomSubHandler(&roomSubService)

	return &RoomSubDependencies{
		rr: roomSubRepository,
		rs: roomSubService,
		rh: roomSubHandler,
	}
}
