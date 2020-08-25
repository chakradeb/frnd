package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/chakradeb/frnd-server/config"
	"github.com/chakradeb/frnd-server/factory"
	"github.com/chakradeb/frnd-server/handlers"
)

func main() {
	conf, errs := config.New()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Errorf("frnd: %s", err)
		}
		log.Fatal("Configuration error. Server could not start")
	}

	f := factory.New(conf)
	logger := f.Logger()
	logger.WithFields(conf.ShowConfig()).Debug("frnd: Configuration: ")
	server := http.Server{
		Addr:              fmt.Sprintf(":%d", conf.AppPort),
		Handler:           handlers.Router(logger, f.DB(), conf.AppSecret),
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}
	logger.Info("frnd: server starting on port: ", conf.AppPort)

	logger.Fatal("frnd: could not start server: ", server.ListenAndServe())
}
