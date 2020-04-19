package logger

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

// Setup logger
func Setup() {
	logger.SetFormatter(&logger.JSONFormatter{
		PrettyPrint: true,
	})

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.DebugLevel)
}
