package slack

import "github.com/slack-go/slack"

type SlackClient struct {
	Key    string
	client *slack.Client
}

func (s *SlackClient) init() {
	s.client = slack.New(s.Key)
}

func (s *SlackClient) SendMessage(channel string, message string) (interface{}, error) {
	s.init()
	x, y, z, err := s.client.SendMessage(channel, slack.MsgOptionText(message, false))
	return SlackMessageResponse{X: x, Y: y, Z: z}, err
}
