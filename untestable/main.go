package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

const (
	CHANNEL_NAME string = "#team-thor"
)

var logger = logrus.New()
var sender = slack.New(os.Getenv("SLACK_APP_TOKEN"))

func main() {
	a, b, c, err := sender.SendMessage(
		CHANNEL_NAME,
		slack.MsgOptionText("Hi, How are you?", false),
	)
	if err != nil {
		logger.Infoln("Fail to send message", err)
	}
	logger.Infoln("Message Sent, result: ", a, b, c)

}
