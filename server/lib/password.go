package lib

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("encrypt: error while encryption: %s", err)
	}
	return string(hash), nil
}

func CheckPassword(pwd, dbPwd []byte, logger *logrus.Logger) bool {
	err := bcrypt.CompareHashAndPassword(dbPwd, pwd)
	if err != nil {
		logger.Errorf("compare: error while comparison: %s", err)
		return false
	}
	return true
}
