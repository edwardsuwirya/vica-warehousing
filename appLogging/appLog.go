package appLogging

import (
	"github.com/sirupsen/logrus"
)

var Logger *AppLogger

type AppLogger struct {
	logger logrus.Logger
}

func (a *AppLogger) LogError(serviceName, event, message string) {
	a.logger.WithField(serviceName, event).Error(message)
}

func (a *AppLogger) LogDebug(serviceName, event string, message interface{}) {
	a.logger.WithField(serviceName, event).Debug(message)
}

func NewAppLogger(l logrus.Logger) *AppLogger {
	return &AppLogger{logger: l}
}
