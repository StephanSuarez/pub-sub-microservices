package http

import (
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/repository"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type UsersDependencies struct {
	ur repository.UserSubRepository
	us services.UserSubService
	uh UserSubHandler
}

func NewUsersSubDependencies(dbInstance *mongo.Database) *UsersDependencies {
	userSubRepository := repository.NewUserSubRepository(dbInstance)
	userSubService := services.NewUserSubService(&userSubRepository)
	userSubHandler := NewUserSubHandler(&userSubService)

	return &UsersDependencies{
		ur: userSubRepository,
		us: userSubService,
		uh: userSubHandler,
	}
}
