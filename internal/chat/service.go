package chat

import (
	"context"
	stream "github.com/GetStream/stream-chat-go/v5"
)

type Service struct {
	Client *stream.Client
}

func NewService(client *stream.Client) *Service {
	return &Service{Client: client}
}

// CreateChannel creates a Stream channel for a proposal
func (s *Service) CreateChannel(proposalID string, members []string) (*stream.Channel, error) {
	channel := s.Client.channe("messaging", proposalID, stream.ChannelOptions{
		Members: members,
	})
	if err != nil {
		return nil, err
	}

	_, err = channel.Create(context.Background())
	if err != nil {
		return nil, err
	}

	return channel, nil
}

// SendMessage sends a message to a channel
func (s *Service) SendMessage(channel *stream.Channel, userID, text string) error {
	_, err := s.SendMessage(channel, userID, text)
	return err
}
