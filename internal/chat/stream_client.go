package chat

import (
	"log"
	"os"

	stream "github.com/GetStream/stream-chat-go/v7"
)

// NewStreamClient initializes the Stream Chat client
func NewStreamClient() *stream.Client {
	apiKey := os.Getenv("STREAM_API_KEY")
	apiSecret := os.Getenv("STREAM_API_SECRET")

	client, err := stream.NewClient(apiKey, apiSecret)
	if err != nil {
		log.Fatalf("Failed to initialize Stream client: %v", err)
	}

	return client
}
