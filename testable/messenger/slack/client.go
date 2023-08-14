package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SlackClient struct {
	Key    string
	client *slack.Client
}

func (s *SlackClient) init() error {
	if s.Key == "" {
		return fmt.Errorf("key is empty")
	}
	s.client = slack.New(s.Key)
	return nil
}

func (s *SlackClient) SendMessage(channel string, message string) (interface{}, error) {
	err := s.init()

	if err != nil {
		return nil, err
	}

	if channel == "" {
		return nil, fmt.Errorf("channel name is empty")
	}
	if message == "" {
		return nil, fmt.Errorf("message is empty")
	}

	x, y, z, err := s.client.SendMessage(channel, slack.MsgOptionText(message, false))
	return SlackMessageResponse{X: x, Y: y, Z: z}, err
}
