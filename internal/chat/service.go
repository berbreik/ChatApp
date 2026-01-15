package chat

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v7"
)

type Service struct {
	Client *stream.Client
}

func NewService(client *stream.Client) *Service {
	return &Service{Client: client}
}

// CreateChannel creates a new channel with either a fixed ID or distinct members.
// If channelID is empty, Stream will create a distinct channel based on members.
func (s *Service) CreateChannel(ctx context.Context, channelType, channelID, creatorID string, members []string) (*stream.CreateChannelResponse, error) {
	data := &stream.ChannelRequest{
		Members: members,
	}

	channel, err := s.Client.CreateChannel(ctx, channelType, channelID, creatorID, data)
	if err != nil {
		return nil, err
	}
	return channel, nil
}

// SendMessage sends a message to an existing channel
func (s *Service) SendMessage(ctx context.Context, channelType, channelID, userID, text string) error {
	channel := s.Client.Channel(channelType, channelID)
	msg := &stream.Message{Text: text}
	_, err := channel.SendMessage(ctx, msg, userID)
	return err
}

// AddMembers adds users to a channel
func (s *Service) AddMembers(ctx context.Context, channelType, channelID string, members []string) error {
	channel := s.Client.Channel(channelType, channelID)
	_, err := channel.AddMembers(ctx, members)
	return err
}

// RemoveMembers removes users from a channel
func (s *Service) RemoveMembers(ctx context.Context, channelType, channelID string, members []string) error {
	channel := s.Client.Channel(channelType, channelID)
	_, err := channel.RemoveMembers(ctx, members, nil)
	return err
}
