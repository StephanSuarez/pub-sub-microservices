package http

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/tesis/internal/common/utils"
	"github.com/tesis/internal/rooms/http/dtos"
	"github.com/tesis/internal/rooms/services"
)

type roomHandler struct {
	rs services.RoomService
}

type RoomHandler interface {
	CreateRoom(ctx *gin.Context)
	GetRooms(ctx *gin.Context)
	GetRoomByID(ctx *gin.Context)
	UpdateRoom(ctx *gin.Context)
	DeleteRoom(ctx *gin.Context)

	AddUserToRoom(ctx *gin.Context)
	RemoveUserInRoom(ctx *gin.Context)

	AddUserToRoomSub(msg *pubsub.Message)
}

func NewRoomHandler(roomService *services.RoomService) RoomHandler {
	return &roomHandler{
		rs: *roomService,
	}
}

func (rh *roomHandler) CreateRoom(ctx *gin.Context) {
	roomDto := dtos.RoomReqDTO{}

	if err := ctx.ShouldBind(&roomDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Body",
		})
		return
	}

	validateData := utils.ValidateData(roomDto)
	if validateData != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": validateData,
		})
		return
	}

	if !utils.CheckRoomStatus(roomDto.Status) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Room Status is not valid",
		})
		return
	}

	stringID, err := rh.rs.CreateRoom(roomDto.MapEntityFromDto())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	roomDtoResponse := dtos.RoomResDTO{
		ID:          stringID,
		Name:        roomDto.Name,
		Description: roomDto.Description,
		CreatedBy:   roomDto.CreatedBy,
		NumMaxUsers: roomDto.NumMaxUsers,
		Status:      roomDto.Status,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": roomDtoResponse,
	})
}

func (rh *roomHandler) GetRooms(ctx *gin.Context) {
	rooms, err := rh.rs.GetRooms()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var roomsResDto []dtos.RoomResDTO
	for i := 0; i < len(rooms); i++ {
		roomsResDto = append(roomsResDto, dtos.RoomResDTO(rooms[i]))
	}

	ctx.JSON(http.StatusOK, roomsResDto)
}

func (rh *roomHandler) GetRoomByID(ctx *gin.Context) {
	id := ctx.Param("id")
	room, err := rh.rs.GetRoomByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	roomResDto := dtos.RoomResDTO{}
	roomResDto.MapEntityToDto(room)

	ctx.JSON(http.StatusOK, roomResDto)
}

func (rh *roomHandler) UpdateRoom(ctx *gin.Context) {
	id := ctx.Param("id")

	roomDto := dtos.RoomReqDTO{}
	if err := ctx.ShouldBind(&roomDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validateData := utils.ValidateData(roomDto)
	if validateData != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": validateData,
		})
		return
	}

	if roomDto.Status != "" {
		if !utils.CheckRoomStatus(roomDto.Status) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Room Status is not valid",
			})
			return
		}
	}

	roomEntity, err := rh.rs.UpdateRoom(id, roomDto.MapEntityFromDto())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roomEntity)
}

func (rh *roomHandler) DeleteRoom(ctx *gin.Context) {
	id := ctx.Param("id")
	err := rh.rs.DeleteRoom(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("room with id: %s deleted", id)})
}

// Users in room domain

func (rh *roomHandler) AddUserToRoom(ctx *gin.Context) {
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

func (rh *roomHandler) RemoveUserInRoom(ctx *gin.Context) {
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

func (rh *roomHandler) AddUserToRoomSub(msg *pubsub.Message) {
	log.Println(msg)
}
