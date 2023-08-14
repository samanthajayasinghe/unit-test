package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

var logger = logrus.New()
var sender = slack.New(os.Getenv("SLACK_APP_TOKEN"))

func main() {
	a, b, c, err := sender.SendMessage(
		"#team-thor",
		slack.MsgOptionText("Hi, This is test message", false),
	)
	if err != nil {
		logger.Errorln("Fail to send message", err)
		return
	}
	logger.Infoln("Message Sent, result: ", a, b, c)

}
