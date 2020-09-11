package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/chakradeb/frnd-server/db"
)

func Router(logger *logrus.Logger, db *db.DB, appSecret string) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/signup", SignupHandler(logger, db, appSecret)).Methods(http.MethodPost)
	r.HandleFunc("/api/login", LoginHandler(logger, db, appSecret)).Methods(http.MethodPost)

	return r
}
