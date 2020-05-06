package logger

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

// Setup logger
func Setup() {
	logger.SetFormatter(&logger.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "02-Jan-2006 15:04:05", // https://golang.org/src/time/format.go
		FullTimestamp:   true,
	})

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.DebugLevel)
}
