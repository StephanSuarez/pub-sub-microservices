package http

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var projectID string

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	projectID = os.Getenv("PROJECT_ID")
}

// func AddUserToRoomSub(w io.Writer, rsdep *RoomSubDependencies) error {
// 	subID := "add-user-to-room-sub"
// 	ctx := context.Background()
// 	client, err := pubsub.NewClient(ctx, projectID)
// 	if err != nil {
// 		return fmt.Errorf("pubsub.NewClient: %w", err)
// 	}

// 	defer client.Close()

// 	sub := client.Subscription(subID)

// 	var received int32
// 	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
// 		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
// 		atomic.AddInt32(&received, 1)
// 		rsdep.rh.AddUserToRoomSub(msg)
// 		msg.Ack()
// 	})

// 	if err != nil {
// 		return fmt.Errorf("sub.Receive: %w", err)
// 	}

// 	fmt.Fprintf(w, "Received %d messages\n", received)

// 	return nil
// }
