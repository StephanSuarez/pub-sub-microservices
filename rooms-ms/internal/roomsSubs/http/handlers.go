package http

import (
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/tesis/internal/roomsSubs/http/dtos"
	"github.com/tesis/internal/roomsSubs/services"
)

type roomSubHandler struct {
	rs services.RoomSubservice
}

type RoomSubHandler interface {
	AddUserToRoom(ctx *gin.Context)
	RemoveUserInRoom(ctx *gin.Context)

	AddUserToRoomSub(msg *pubsub.Message)
}

func NewRoomSubHandler(roomService *services.RoomSubservice) RoomSubHandler {
	return &roomSubHandler{
		rs: *roomService,
	}
}

func (rh *roomSubHandler) AddUserToRoom(ctx *gin.Context) {
	roomID := ctx.Param("id")
	userIdReq := dtos.UserIdReq{}

	if err := ctx.ShouldBind(&userIdReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := rh.rs.AddUserToRoom(roomID, userIdReq.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user added to room"})
}

func (rh *roomSubHandler) RemoveUserInRoom(ctx *gin.Context) {
	roomID := ctx.Param("id")
	userIdReq := dtos.UserIdReq{}

	if err := ctx.ShouldBind(&userIdReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := rh.rs.RemoveUserInRoom(roomID, userIdReq.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user removed to room"})
}

func (rh *roomSubHandler) AddUserToRoomSub(msg *pubsub.Message) {
	log.Println(msg)
}
