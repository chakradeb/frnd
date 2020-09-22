package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/chakradeb/frnd-server/constants"
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

		dbUser, err := db.GetUser(user.Username)
		if err != nil {
			msg := fmt.Sprintf("login: db: user %s doesn't exist", user.Username)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusUnauthorized, logger)
			return
		}

		ok := PasswordChecker([]byte(user.Password), []byte(dbUser.Password), logger)
		if !ok {
			msg := fmt.Sprintf("login: db: password didn't match for user %s", user.Username)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusUnauthorized, logger)
			return
		}

		accessToken, err := TokenCreator(user.Username, time.Now(), constants.ACCESS_TOKEN_LIVE_TIME, appSecret)
		if err != nil {
			msg := fmt.Sprintf("login: create token: not able to sign token: %s", err)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		refreshToken, err := TokenCreator("", time.Now(), constants.REFRESH_TOKEN_LIVE_TIME, appSecret)
		if err != nil {
			msg := fmt.Sprintf("signup: create refresh token: not able to sign token: %s", err)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		session := models.Session{Username: user.Username, AccessToken: accessToken, RefreshToken: refreshToken}
		lib.WriteResponse(w, session, http.StatusCreated, logger)
		return
	}
}
