package lib

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/chakradeb/frnd-server/models"
)

func CreateToken(username string, email string, secret string) (string, error) {
	t := time.Now()
	token := &models.Token{
		Username: username,
		Email:    email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: t.Add(time.Minute * 1).Unix(),
			IssuedAt:  t.Unix(),
		},
	}

	signedToken := jwt.NewWithClaims(jwt.GetSigningMethod(jwt.SigningMethodHS256.Name), token)
	return signedToken.SignedString([]byte(secret))
}
