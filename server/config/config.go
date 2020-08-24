package config

import (
	"fmt"
	"strings"

	"github.com/chakradeb/env"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AppPort int
	LogLevel logrus.Level
	Hosts []string
	Keyspace string
}

type args struct {
	AppPort int `env:"PORT" default:"8000"`
	LogLevel string `env:"LOG_LEVEL" default:"info"`
	Hosts string `env:"DB_HOSTS"`
	Keyspace string `env:"DB_KEYSPACE"`
}

func New() (*Config, []error) {
	args := &args{}

	errs := env.Parse(args)
	if len(errs) > 0 {
		return nil, errs
	}

	return newConfig(args)
}

func newConfig(args *args) (*Config, []error) {
	var errs []error

	logLevel, err := logrus.ParseLevel(args.LogLevel)
	if err != nil {
		errs = append(errs, fmt.Errorf("config: %s", err))
	}

	if len(errs) > 0 {
		return nil, errs
	}

	conf := &Config{
		AppPort: args.AppPort,
		LogLevel: logLevel,
		Hosts: strings.Split(args.Hosts, ","),
		Keyspace: args.Keyspace,
	}
	return conf, nil
}

func (conf Config) ShowConfig() logrus.Fields {
	return logrus.Fields{
		"AppPort": conf.AppPort,
		"LogLevel": conf.LogLevel,
		"Hosts": conf.Hosts,
		"Keyspace": conf.Keyspace,
	}
}
