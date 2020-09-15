package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"warehousing/appLogging"
)

type AppConfig struct {
	logFilePath string
}

func NewConfig() *AppConfig {
	c := &AppConfig{}
	logFile := c.viperGetEnv("LOG_FILE", "/tmp/app.log")
	logLevel := c.viperGetEnv("LOG_LEVEL", "debug")
	c.logFilePath = logFile
	c.logger(logLevel)
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

func (c *AppConfig) viperGetEnv(key, defaultValue string) string {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}
