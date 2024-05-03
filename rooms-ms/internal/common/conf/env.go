package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	IPAddress     string
	ServerAddress string
	PortAddress   string
	MongoServer   string
	MongoUsername string
	MongoPassword string
	MongoCluster  string
	DbName        string
	DbEnviroment  string
}

func NewEnv() *Env {
	env := Env{}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	env.IPAddress = os.Getenv("IP_ADDRESS")
	env.ServerAddress = os.Getenv("SERVER_ADDRESS")
	env.PortAddress = os.Getenv("PORT_SERVER")
	env.MongoServer = os.Getenv("MONGO_SERVER")
	env.MongoUsername = os.Getenv("MONGO_USERNAME")
	env.MongoPassword = os.Getenv("MONGO_PASSWORD")
	env.MongoCluster = os.Getenv("MONGO_CLUSTER")
	env.DbName = os.Getenv("DB_NAME")
	env.DbEnviroment = os.Getenv("DB_ENVIROMENT")

	return &env
}
