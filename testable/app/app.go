package app

type App struct {
	MsgClient MessageClientInterface
	Logger    LoggerInterface
}

func (c *App) SendMessage(channel string, message string) error {

	result, err := c.MsgClient.SendMessage(channel, message)
	if err != nil {
		c.Logger.Infoln("Fail to send message", err)
		return err
	}
	c.Logger.Infoln("Message Sent, result: ", result)
	return nil
}
