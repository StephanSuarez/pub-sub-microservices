package http

import (
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/repository"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type UsersDependencies struct {
	ur repository.UserRepository
	us services.UserService
	uh UserHandler
}

func NewUsersDependencies(dbInstance *mongo.Database) *UsersDependencies {
	userRepository := repository.NewUserRepository(dbInstance)
	userService := services.NewUserService(&userRepository)
	userHandler := NewUserHandler(&userService)

	return &UsersDependencies{
		ur: userRepository,
		us: userService,
		uh: userHandler,
	}
}
