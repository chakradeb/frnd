package lib

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/chakradeb/frnd-server/models"
)

func CreateToken(username string, createTime time.Time, liveTime time.Duration, secret string) (string, error) {
	claim := &models.Claims{
		Username: username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: createTime.Add(liveTime).Unix(),
			IssuedAt:  createTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(jwt.SigningMethodHS256.Name), claim)
	return token.SignedString([]byte(secret))
}
