package factory

import (
	"os"
	"strings"

	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"

	"github.com/chakradeb/frnd-server/config"
)

type Factory struct {
	logger *logrus.Logger
	dbSession *gocql.Session
}

func New(config *config.Config) *Factory {
	logger := createLogger(config.LogLevel)
	return &Factory{
		logger: logger,
		dbSession: createDBSession(config.Hosts, config.Keyspace, logger),
	}
}

func (f Factory) Logger() *logrus.Logger {
	return f.logger
}

func (f Factory) DbSession() *gocql.Session {
	return f.dbSession
}

func createDBSession(clusterIPs []string, keyspace string, logger *logrus.Logger) *gocql.Session {
	cluster := gocql.NewCluster(clusterIPs[:]...)
	cluster.Keyspace = keyspace
	session, err := cluster.CreateSession()
	if err != nil {
		logger.WithError(err).Fatalf("factory: connecting to \"%s\"", strings.Join(clusterIPs, ","))
	}
	logger.Infof("factory: connected to cluster, %s", strings.Join(clusterIPs, ","))
	return session
}

func createLogger(logLevel logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logLevel)
	return logger
}
