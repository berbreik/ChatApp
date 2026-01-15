package chat

import stream "github.com/GetStream/stream-chat-go/v5"

func NewStreamClient(apiKey, apiSecret string) (*stream.Client, error) {
	return stream.NewClient(apiKey, apiSecret)
}
