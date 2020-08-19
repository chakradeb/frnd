package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chakradeb/frnd-server/config"
	"github.com/chakradeb/frnd-server/factory"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf, errs := config.New()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Error(err)
		}
		log.Fatal("Configuration error. Server could not start")
	}

	f := factory.New(conf)
	logger := f.Logger()
	server := http.Server{
		Addr:              fmt.Sprintf(":%d", conf.AppPort),
		Handler:           nil,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}
	logger.Info("server starting on port", conf.AppPort)
	logger.WithFields(conf.ShowConfig()).Debug("Configuration: ")
	logger.Fatal("could not start server: ", server.ListenAndServe())
}
