package factory

import (
	"os"

	"github.com/chakradeb/frnd-server/config"
	"github.com/sirupsen/logrus"
)

type Factory struct {
	logger *logrus.Logger
}

func New(config *config.Config) *Factory {
	return &Factory{
		logger: createLogger(config.LogLevel),
	}
}

func (f Factory) Logger() *logrus.Logger {
	return f.logger
}

func createLogger(logLevel logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logLevel)
	return logger
}
