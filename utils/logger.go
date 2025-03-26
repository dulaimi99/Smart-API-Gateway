package utils

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(cfg *Config) *logrus.Logger {
	Log := logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// Open a file for logging
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatalf("Failed to open log file: %v", err)
	}

	// Set the log output to both stdout and the log file
	multiWriter := io.MultiWriter(os.Stdout, file)
	Log.SetOutput(multiWriter)

	level, err := logrus.ParseLevel(cfg.Logging.Level)
	if err != nil {
		Log.Fatalf("Invalid log level: %s", cfg.Logging.Level)
	}
	Log.SetLevel(level)
	return Log
}
