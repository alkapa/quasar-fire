package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stderr)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(false)
	logger.SetLevel(
		func() logrus.Level {
			if lvl, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL")); err == nil {
				return lvl
			}
			return logrus.InfoLevel
		}(),
	)

	return logger
}
