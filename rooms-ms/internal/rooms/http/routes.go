package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, appRoomDependencies *RoomDependencies) {
	routesRooms := r.Group("/v1/rooms")
	routesRooms.GET("", appRoomDependencies.rh.GetRooms)
	routesRooms.POST("", appRoomDependencies.rh.CreateRoom)
	routesRooms.GET(":id", appRoomDependencies.rh.GetRoomByID)
	routesRooms.PUT(":id", appRoomDependencies.rh.UpdateRoom)
	routesRooms.DELETE(":id", appRoomDependencies.rh.DeleteRoom)

	// add user to room
	routesRooms.PUT("add-user/:id", appRoomDependencies.rh.AddUserToRoom)
	routesRooms.PUT("remove-user/:id", appRoomDependencies.rh.RemoveUserInRoom)
}
