package main

import (
	"os"

	"github.com/samanthajayasinghe/unit-test/testable/app"
	"github.com/samanthajayasinghe/unit-test/testable/messenger/slack"
	"github.com/sirupsen/logrus"
)

func main() {
	sender := slack.SlackClient{Key: os.Getenv("SLACK_APP_TOKEN")}
	logger := logrus.New()
	app := &app.App{
		MsgClient: &sender,
		Logger:    logger,
	}
	app.SendMessage("#team-thor", "Hi, from testable app")
}
