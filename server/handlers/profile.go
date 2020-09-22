package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/chakradeb/frnd-server/db"
	"github.com/chakradeb/frnd-server/lib"
)

func ProfileHandler(logger *logrus.Logger, db db.IDBClient, appSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["id"]

		profile, err := db.GetProfile(username)
		if err != nil {
			msg := fmt.Sprintf("profile: db: user %s doesn't exist", username)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusNotFound, logger)
			return
		}

		lib.WriteResponse(w, profile, http.StatusOK, logger)
		return
	}
}
