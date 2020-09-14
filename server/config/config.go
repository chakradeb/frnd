package config

import (
	"errors"
	"fmt"
	"github.com/chakradeb/env"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AppPort int
	LogLevel logrus.Level
	DBHost string
	DBPort int
	DBName string
	AppSecret string
}

type args struct {
	AppPort int `env:"PORT" default:"80"`
	LogLevel string `env:"LOG_LEVEL" default:"info"`
	DBHost string `env:"DB_HOST"`
	DBPort int `env:"DB_PORT" default:"5432"`
	DBName string `env:"DB_NAME"`
	AppSecret string `env:"APP_SECRET"`
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

	if len(args.AppSecret) < 8 {
		errs = append(errs, errors.New("config: app secret should be minimum of 8 characters"))
	}

	if len(errs) > 0 {
		return nil, errs
	}

	conf := &Config{
		AppPort: args.AppPort,
		LogLevel: logLevel,
		DBHost: args.DBHost,
		DBPort: args.DBPort,
		DBName: args.DBName,
		AppSecret: args.AppSecret,
	}
	return conf, nil
}

func (conf Config) ShowConfig() logrus.Fields {
	return logrus.Fields{
		"AppPort": conf.AppPort,
		"LogLevel": conf.LogLevel,
		"DBHost": conf.DBHost,
		"DBPort": conf.DBPort,
		"DBName": conf.DBName,
		"AppSecret": "[SECRET]",
	}
}
