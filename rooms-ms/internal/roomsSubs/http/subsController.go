package http

import (
	"log"
)

func ListeningSubs(rsdep *RoomSubDependencies) {
	errChan := make(chan error)

	go func() {
		// errChan <- AddUserToRoomSub(os.Stdout, rsdep)
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			log.Fatal(err)
		}
	}
}
