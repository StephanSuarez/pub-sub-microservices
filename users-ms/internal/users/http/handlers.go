package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/entity"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/http/dtos"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/services"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	us services.UserService
}

type UserHandler interface {
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetUserByUserNameOrEmail(ctx *gin.Context)
}

func NewUserHandler(us *services.UserService) UserHandler {
	return &userHandler{
		us: *us,
	}
}

func (uh *userHandler) CreateUser(ctx *gin.Context) {
	userDto := dtos.UsersRequestDTO{}

	if err := ctx.ShouldBind(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Body",
		})
		return
	}

	if err := userDto.ValidateString(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if uh.userNameRegistered(userDto.UserName) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "userName already exist",
		})
		return
	}

	if uh.emailRegistered(userDto.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email already exist",
		})
		return
	}

	userID, err := uh.us.CreateUser(userDto.MapEntityFromDto())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userDtoRes := dtos.UsersResDTO{
		ID:       userID,
		UserName: userDto.UserName,
		Email:    userDto.Email,
	}

	ctx.JSON(http.StatusOK, userDtoRes)
}

func (uh *userHandler) GetUsers(ctx *gin.Context) {
	users, err := uh.us.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var usersResDto []dtos.UsersResDTO
	for i := 0; i < len(users); i++ {
		usersResDto = append(usersResDto, dtos.UsersResDTO(users[i]))
	}

	ctx.JSON(http.StatusOK, usersResDto)
}

func (uh *userHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uh.us.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userResDto := dtos.UsersResDTO{}
	userResDto.MapEntityToDto(user)

	ctx.JSON(http.StatusOK, userResDto)
}

func (uh *userHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userDto := dtos.UsersRequestDTO{}
	if err := ctx.ShouldBind(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := userDto.ValidateString(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if uh.userNameRegistered(userDto.UserName) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "userName already exist",
		})
		return
	}

	if uh.emailRegistered(userDto.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email already exist",
		})
		return
	}

	userEntity, err := uh.us.UpdateUser(id, userDto.MapEntityFromDto())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userEntity)
}

func (uh *userHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uh.us.DeleteUser(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user with id: %s deleted", id)})
}

func (uh *userHandler) GetUserByUserNameOrEmail(ctx *gin.Context) {
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

func (uh *userHandler) userNameRegistered(userName string) bool {
	user, _ := uh.us.GetUserByUserName(userName)
	return user != nil
}

func (uh *userHandler) emailRegistered(email string) bool {
	user, _ := uh.us.GetUserByEmail(email)
	return user != nil
}
