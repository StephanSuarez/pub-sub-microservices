package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	MongoServer   string
	MongoUsername string
	MongoPassword string
	MongoCluster  string
	DbName        string
	DbEnviroment  string
	IPAddress     string
	ServerAddress string
	PortServer    string
	ProjectID     string
}

func NewEnv() *Env {
	env := Env{}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	env.MongoServer = os.Getenv("MONGO_SERVER")
	env.MongoUsername = os.Getenv("MONGO_USERNAME")
	env.MongoPassword = os.Getenv("MONGO_PASSWORD")
	env.MongoCluster = os.Getenv("MONGO_CLUSTER")
	env.DbName = os.Getenv("DB_NAME")
	env.DbEnviroment = os.Getenv("DB_ENVIROMENT")
	env.IPAddress = os.Getenv("IP_ADDRESS")
	env.ServerAddress = os.Getenv("SERVER_ADDRESS")
	env.PortServer = os.Getenv("PORT_SERVER")
	env.ProjectID = os.Getenv("PROJECT_ID")

	return &env
}
