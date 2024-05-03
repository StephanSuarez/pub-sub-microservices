package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/entity"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/http/dtos"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/services"

	"github.com/gin-gonic/gin"
)

type userSubHandler struct {
	us services.UserSubService
}

type UserSubHandler interface {
	GetUserByUserNameOrEmail(ctx *gin.Context)

	addRoomToList(msg string) error
	removeRoomInList(data string) error
}

func NewUserSubHandler(us *services.UserSubService) UserSubHandler {
	return &userSubHandler{
		us: *us,
	}
}

func (uh *userSubHandler) GetUserByUserNameOrEmail(ctx *gin.Context) {
	userName := ctx.Query("username")
	email := ctx.Query("email")

	log.Println(userName, email)
	if userName == "" && email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UserName or Email is require"})
		return
	}

	log.Println("Username:", userName)

	var userEntity *entity.UsersRes
	var err error

	if userName != "" {
		userEntity, err = uh.us.GetUserByUserName(userName)
	} else if email != "" {
		userEntity, err = uh.us.GetUserByEmail(email)
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"res": nil, "error": err.Error()})
		return
	}

	var userRDto dtos.UsersResDTO
	userRDto.MapEntityToDto(userEntity)

	ctx.JSON(http.StatusOK, userRDto)
}

func (uh *userSubHandler) userNameRegistered(userName string) bool {
	user, _ := uh.us.GetUserByUserName(userName)
	return user != nil
}

func (uh *userSubHandler) emailRegistered(email string) bool {
	user, _ := uh.us.GetUserByEmail(email)
	return user != nil
}

func (uh *userSubHandler) addRoomToList(data string) error {
	addRoomDTO := dtos.AddRoom{}
	err := json.Unmarshal([]byte(data), &addRoomDTO)
	if err != nil {
		return err
	}

	if err := uh.us.AddRoomToList(addRoomDTO.UserID, addRoomDTO.RoomID); err != nil {
		return err
	}

	return nil
}

func (uh *userSubHandler) removeRoomInList(data string) error {
	addRoomDTO := dtos.AddRoom{}
	err := json.Unmarshal([]byte(data), &addRoomDTO)
	if err != nil {
		return err
	}

	if err := uh.us.RemoveRoomInList(addRoomDTO.UserID, addRoomDTO.RoomID); err != nil {
		return err
	}

	return nil
}
