package function

import (
	"fmt"
	"net/http"

	"github.com/StephanSuarez/chat-rooms-users-ms/cmd"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

var app *cmd.App

func init() {
	fmt.Println("creating app")
	app = cmd.NewApp()
	// app.Start()
	fmt.Println("app created")

	functions.HTTP("UsersMS", usersMS)
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func usersMS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CALL INTO A FUNC USERSMS")
	// app.Router.ServeHTTP(w, r)
}
