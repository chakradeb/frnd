package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"github.com/chakradeb/frnd-server/constants"
	"github.com/chakradeb/frnd-server/db"
	"github.com/chakradeb/frnd-server/lib"
	"github.com/chakradeb/frnd-server/models"
)

func RefreshHandler(logger *logrus.Logger, db db.IDBClient, appSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Refresh")
		auth = strings.TrimSpace(auth)

		if auth == "" {
			msg := "refresh: missing refresh token"
			logger.Error(msg)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(appSecret), nil
		})
		if err != nil {
			msg := fmt.Errorf("refresh: error while parsing token: %s", err)
			logger.Error(msg)
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			logger.Error("refresh: token validation failed")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user := &models.User{
			Username: r.Header.Get("Username"),
		}

		accessToken, err := TokenCreator(user.Username, time.Now(), constants.ACCESS_TOKEN_LIVE_TIME, appSecret)
		if err != nil {
			msg := fmt.Sprintf("refresh: create token: not able to sign token: %s", err)
			logger.Error(msg)
			lib.WriteResponse(w, msg, http.StatusInternalServerError, logger)
			return
		}

		session := models.Session{AccessToken: accessToken}
		lib.WriteResponse(w, session, http.StatusCreated, logger)
		return
	}
}
