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

var Encrypter = lib.Encrypt
var TokenCreator = lib.CreateToken

func SignupHandler(logger *logrus.Logger, db db.IDBClient, appSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}

		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			msg := fmt.Sprintf("signup: unable to read request body: %s", err)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		if user.Email == "" || user.Username == "" || user.Gender == "" || user.Password == "" {
			msg := "signup: fields should not be empty"
			logger.Errorf("signup: %s", msg)
			lib.WriteResponse(w, msg, http.StatusUnauthorized, logger)
			return
		}

		ok := db.CheckUserAlreadyExists(user.Username)
		if ok {
			msg := fmt.Sprintf("signup: db: user %s already exist", user.Username)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusForbidden, logger)
			return
		}

		password, err := Encrypter([]byte(user.Password))
		if err != nil {
			msg := fmt.Sprintf("signup: %s", err)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		err = db.CreateUser(user.Username, password)
		if err != nil {
			msg := fmt.Sprintf("signup: db: can not save user details: %s", err)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		token, err := TokenCreator(user.Username, time.Now(), appSecret)
		if err != nil {
			msg := fmt.Sprintf("signup: create token: not able to sign token: %s", err)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		session := models.Session{Username: user.Username, Token: token}
		lib.WriteResponse(w, session, http.StatusCreated, logger)
		return
	}
}
