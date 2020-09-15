package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"warehousing/appLogging"
)

type AppConfig struct {
	logFilePath string
}

func NewConfig() *AppConfig {
	c := &AppConfig{
		logFilePath: "warehousing.log",
	}
	c.logger("debug")
	return c
}

func (c *AppConfig) logger(level string) {
	logFormat := new(logrus.JSONFormatter)
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	log := logrus.Logger{
		Out:       os.Stdout,
		Formatter: logFormat,
		Level:     logLevel,
	}
	var file, errFile = os.OpenFile(c.logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errFile != nil {
		panic(err)
	}
	log.SetOutput(file)
	appLogging.Logger = appLogging.NewAppLogger(log)
}
