package lib

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("encrypt: error while encryption: %s", err)
	}
	return string(hash), nil
}
