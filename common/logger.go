package common

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()
const DefaultLogLevel = logrus.InfoLevel

func init() {
	Logger.SetLevel(getLogLevel())
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func getLogLevel() logrus.Level {
	logLevel, logLevelPresent := os.LookupEnv("LOG_LEVEL")
	if !logLevelPresent {
		return DefaultLogLevel
	}
	customLogLevel, customLogLevelErr := logrus.ParseLevel(logLevel)
	if customLogLevelErr != nil {
		return DefaultLogLevel
	}
	return customLogLevel
}