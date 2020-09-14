package factory

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/chakradeb/frnd-server/config"
	"github.com/chakradeb/frnd-server/db"
)

type Factory struct {
	logger *logrus.Logger
	db *db.DB
}

func New(config *config.Config) *Factory {
	logger := createLogger(config.LogLevel)
	return &Factory{
		logger: logger,
		db: createDBSession(config.DBHost, config.DBPort, config.DBName, logger),
	}
}

func (f Factory) Logger() *logrus.Logger {
	return f.logger
}

func (f Factory) DB() *db.DB {
	return f.db
}

func createDBSession(dbHost string, dbPort int, dbName string, logger *logrus.Logger) *db.DB {
	dbConn, err := db.New(dbHost, dbPort, dbName)
	if err != nil {
		logger.Fatalf("factory: %s", err)
	}
	logger.Infof("factory: connected to database, %s on port %d", dbHost, dbPort)
	return dbConn
}

func createLogger(logLevel logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logLevel)
	return logger
}
