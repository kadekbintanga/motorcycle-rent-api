package config

import (
	"io"
	"motorcycle-rent-api/app/global"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func InitLogger(config *global.EnvConfig) {
	var writers []io.Writer

	logToStdout := config.LogSTDOUT
	logFileLocation := config.LogFileLocation
	logLevelStr := config.LogLevel

	loggerConf := logrus.New()

	loggerConf.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	if logFileLocation != "" {
		logLocation := logFileLocation

		logDir := filepath.Dir(logLocation)

		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			loggerConf.Fatalf("Failed to create log directory: %v", err)
		}

		global.LogFile, err = os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			loggerConf.Fatalf("Failed to open log file: %v", err)
		}

		writers = append(writers, global.LogFile)
	}

	if logToStdout {
		writers = append(writers, os.Stdout)
	}

	if len(writers) > 0 {
		loggerConf.SetOutput(io.MultiWriter(writers...))
	}

	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = logrus.InfoLevel // Default to INFO if parsing fails
	}
	loggerConf.SetLevel(logLevel)

	global.Logger = loggerConf
}

func CloseLoggerFile() {
	err := global.LogFile.Close()
	if err != nil {
		global.Logger.Fatalf("Unable close log file error : %v", err)
	}
}
