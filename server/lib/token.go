package lib

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/chakradeb/frnd-server/models"
)

func CreateToken(username string, t time.Time, secret string) (string, error) {
	claim := &models.Claims{
		Username: username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: t.Add(time.Minute * 1).Unix(),
			IssuedAt:  t.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(jwt.SigningMethodHS256.Name), claim)
	return token.SignedString([]byte(secret))
}
