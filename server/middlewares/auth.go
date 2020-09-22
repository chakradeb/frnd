package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"github.com/chakradeb/frnd-server/models"
)

func Auth(next http.HandlerFunc, logger *logrus.Logger, appSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		auth = strings.TrimSpace(auth)

		if auth == "" {
			msg := "auth: missing auth token"
			logger.Error(msg)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(appSecret), nil
		})
		if err != nil {
			msg := fmt.Errorf("auth: error while parsing token: %s", err)
			logger.Error(msg)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			logger.Error("auth: token validation failed")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
