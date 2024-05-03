package cmd

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tesis/internal/common/conf"
	"github.com/tesis/internal/rooms/http"
	roomsSubs "github.com/tesis/internal/roomsSubs/http"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Env    *conf.Env
	Router *gin.Engine
	DbConn *mongo.Database
	Rd     *http.RoomDependencies
	Rsd    *roomsSubs.RoomSubDependencies
}

func NewApp() *App {

	app := App{}

	app.Env = conf.NewEnv()

	Dbenv := &conf.DbEnv{
		DbEnviroment: app.Env.DbEnviroment,
		Server:       app.Env.MongoServer,
		Username:     app.Env.MongoUsername,
		Password:     app.Env.MongoPassword,
		Cluster:      app.Env.MongoCluster,
		Dbname:       app.Env.DbName,
	}

	app.DbConn = conf.GetDBInstance(Dbenv)

	app.Rd = http.NewRoomDependencies(app.DbConn)
	app.Rsd = roomsSubs.NewRoomSubDependencies(app.DbConn)

	app.Router = gin.Default()

	return &app
}

func (app *App) Start() {

	go func() {
		roomsSubs.ListeningSubs(app.Rsd)
	}()

	http.Routes(app.Router, app.Rd)

	addr := fmt.Sprintf("%s:%s", app.Env.IPAddress, app.Env.ServerAddress)
	log.Printf("Server is running on: %s", addr)

	err := app.Router.Run(app.Env.PortAddress)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
