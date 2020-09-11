package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/chakradeb/frnd-server/db"
	"github.com/chakradeb/frnd-server/lib"
	"github.com/chakradeb/frnd-server/models"
)

var PasswordChecker = lib.CheckPassword

func LoginHandler(logger *logrus.Logger, db db.IDBClient, appSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}

		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			msg := fmt.Sprintf("login: unable to read request body: %s", err)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		if user.Username == "" || user.Password == "" {
			msg := "login: fields should not be empty"
			logger.Errorf("login: %s", msg)
			lib.WriteResponse(w, msg, http.StatusUnauthorized, logger)
			return
		}

		ok := db.CheckUserAlreadyExists(user.Username)
		if !ok {
			msg := fmt.Sprintf("login: db: user %s doesn't exist", user.Username)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusUnauthorized, logger)
			return
		}

		dbUser := db.GetUser(user.Username)
		if dbUser.Password == "" {
			msg := fmt.Sprintf("login: db: couldn't fetch details of user %s", user.Username)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		ok = PasswordChecker([]byte(user.Password), []byte(dbUser.Password), logger)
		if !ok {
			msg := fmt.Sprintf("login: db: password didn't match for user %s", user.Username)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusUnauthorized, logger)
			return
		}

		token, err := TokenCreator(user.Username, time.Now(), appSecret)
		if err != nil {
			msg := fmt.Sprintf("login: create token: not able to sign token: %s", err)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		session := models.Session{Username: user.Username, Token: token}
		lib.WriteResponse(w, session, http.StatusOK, logger)
		return
	}
}
