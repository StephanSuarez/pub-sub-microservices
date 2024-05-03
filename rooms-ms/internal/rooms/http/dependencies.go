package http

import (
	"github.com/tesis/internal/rooms/repository"
	"github.com/tesis/internal/rooms/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomDependencies struct {
	rr repository.RoomRepository
	rs services.RoomService
	rh RoomHandler
}

func NewRoomDependencies(dbInstanceConn *mongo.Database) *RoomDependencies {
	dbInstance := dbInstanceConn
	roomRepository := repository.NewRoomRepository(dbInstance)
	roomService := services.NewRoomService(&roomRepository)
	roomHandler := NewRoomHandler(&roomService)

	return &RoomDependencies{
		rr: roomRepository,
		rs: roomService,
		rh: roomHandler,
	}
}
