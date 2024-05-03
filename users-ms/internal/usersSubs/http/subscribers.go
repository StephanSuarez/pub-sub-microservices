package http

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

var projectID string

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	projectID = os.Getenv("PROJECT_ID")
}

func addRoomToList(w io.Writer, UserDep *UsersDependencies) error {
	subID := "add-user-to-room-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}

	defer client.Close()

	sub := client.Subscription(subID)

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		data := string(msg.Data)
		if err := UserDep.uh.addRoomToList(data); err != nil {
			log.Println(err)
		} else {
			log.Println("roon added to user list successfully")
		}
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})

	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}

	fmt.Fprintf(w, "Received %d messages\n", received)

	return nil
}

func removeRoomToList(w io.Writer, UserDep *UsersDependencies) error {
	subID := "remove-user-in-room-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}

	defer client.Close()

	sub := client.Subscription(subID)

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		data := string(msg.Data)
		if err := UserDep.uh.removeRoomInList(data); err != nil {
			log.Println(err)
		} else {
			log.Println("roon removed to user list successfully")
		}
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})

	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}

	fmt.Fprintf(w, "Received %d messages\n", received)

	return nil
}
