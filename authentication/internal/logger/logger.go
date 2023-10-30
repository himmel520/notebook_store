package logger

import (
	log "github.com/sirupsen/logrus"
)

var Logger = log.New()

func init() {
	Logger.SetFormatter(&log.TextFormatter{})
	// Logger.SetReportCaller(true)
	Logger.SetLevel(log.InfoLevel)
}
