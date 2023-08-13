package app

// LoggerInterface defines the message logging interface
type LoggerInterface interface {
	Infoln(...interface{})
	Warnln(...interface{})
	Errorln(...interface{})
}
