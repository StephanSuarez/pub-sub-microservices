package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

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

func addUserToRoom(w io.Writer, msg string) error {
	topicID := "add-user-to-room"
	// msg := "Hello World"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %w", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}

func removeUserToRoom(w io.Writer, msg string) error {
	topicID := "remove-user-in-room"

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %w", err)
	}

	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}
