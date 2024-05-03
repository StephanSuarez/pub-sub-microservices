package cmd

import (
	"fmt"
	"log"

	"github.com/StephanSuarez/chat-rooms-users-ms/internal/common/config"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/common/middleware"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/http"
	userSubs "github.com/StephanSuarez/chat-rooms-users-ms/internal/usersSubs/http"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Env        *config.Env
	DbConn     *mongo.Database
	UsersDep   *http.UsersDependencies
	UserSubDep *userSubs.UsersDependencies
	Router     *gin.Engine
}

func NewApp() *App {

	app := &App{}

	app.Env = config.NewEnv()

	dbenv := &config.DbEnv{
		DbEnviroment: app.Env.DbEnviroment,
		Server:       app.Env.MongoServer,
		Username:     app.Env.MongoUsername,
		Password:     app.Env.MongoPassword,
		Cluster:      app.Env.MongoCluster,
		Dbname:       app.Env.DbName,
	}

	app.DbConn = config.GetDBInstance(dbenv)

	app.UsersDep = http.NewUsersDependencies(app.DbConn)
	app.UserSubDep = userSubs.NewUsersSubDependencies(app.DbConn)

	app.Router = gin.Default()
	app.Router.Use(middleware.CorsMiddleware())

	return app
}

func (app *App) Start() {
	go func() {
		userSubs.ListeningSubs(app.UserSubDep)
	}()

	addr := fmt.Sprintf("%s:%s", app.Env.IPAddress, app.Env.ServerAddress)
	log.Printf("Server is running on: %s", addr)

	http.Routes(app.Router, app.UsersDep)

	err := app.Router.Run(app.Env.PortServer)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
