package app

// MessageClientInterface defines the message sending interface
type MessageClientInterface interface {
	SendMessage(channel, message string) (interface{}, error)
}
