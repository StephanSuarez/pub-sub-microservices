package main

import (
	"fmt"

	"github.com/StephanSuarez/chat-rooms-users-ms/cmd"
	// "log"
	// "os"
	// "github.com/joho/godotenv"
	// "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	// fmt.Println("preparing all")

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal("Error loading .env file", err)
	// }

	// port := "8080"
	// envPort := os.Getenv("PORT_SERVER")

	// fmt.Println(envPort)
	// if envPort != "" {
	// 	port = envPort
	// }
	// fmt.Println("ready")

	// if err := funcframework.Start(port); err != nil {
	// 	log.Fatalf("funcframework.Start: %v\n", err)
	// } else {
	// 	log.Println("Function framework started successfully")
	// }

	// fmt.Println("fin")

	app := cmd.NewApp()
	fmt.Println(app)
	app.Start()
}
