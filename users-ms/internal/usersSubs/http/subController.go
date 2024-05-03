package http

import (
	"fmt"
	"log"
	"os"
)

func ListeningSubs(usdep *UsersDependencies) {
	fmt.Println("LISTENING SUBS")
	errChan := make(chan error)

	go func() {
		errChan <- addRoomToList(os.Stdout, usdep)
	}()
	go func() {
		errChan <- removeRoomToList(os.Stdout, usdep)
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			log.Fatal(err)
		}
	}
}
