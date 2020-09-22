package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/chakradeb/frnd-server/mocks"
	"github.com/chakradeb/frnd-server/models"
)

func TestLoginHandler(t *testing.T) {
	username := "admin"
	password := "password"
	appSecret := "abcdefgh"

	PasswordChecker = func(pwd, dbpwd []byte, logger *logrus.Logger) bool {
		return true
	}
	TokenCreator = func(username string, t time.Time, d time.Duration, secret string) (string, error) {
		return secret, nil
	}

	user := &models.User{
		Username: username,
		Password: password,
	}
	token := &models.Session{
		Username: username,
		AccessToken: appSecret,
		RefreshToken: appSecret,
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)
	expectedToken, _ := json.Marshal(token)

	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	mockDB := &mocks.IDBClient{}

	mockDB.On("GetUser", username).Return(user, nil)

	router := LoginHandler(logger, mockDB, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code, "invalid response code")
	assert.JSONEq(t, string(expectedToken), res.Body.String(), "invalid response body")

	mock.AssertExpectationsForObjects(t, mockDB)
}

func TestLoginHandlerWithEmptyData(t *testing.T) {
	username := ""
	password := ""
	appSecret := "abcdefgh"

	user := &models.User{
		Username: username,
		Password: password,
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	router := LoginHandler(logger, nil, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code, "invalid response code")
	assert.Equal(t, `"login: fields should not be empty"`, res.Body.String(), "invalid response body")
}

func TestLoginHandlerWhenUserDoesNotExists(t *testing.T) {
	username := "admin"
	password := "password"
	appSecret := "abcdefgh"

	user := &models.User{
		Username: username,
		Password: password,
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)
	expectedBody := fmt.Sprintf(`"login: db: user %s doesn't exist"`, username)

	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	mockDB := &mocks.IDBClient{}

	mockDB.On("GetUser", username).Return(&models.User{}, errors.New("no user"))

	router := LoginHandler(logger, mockDB, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code, "invalid response code")
	assert.Equal(t, expectedBody, res.Body.String(), "invalid response body")

	mock.AssertExpectationsForObjects(t, mockDB)
}

func TestLoginHandlerWhenPasswordMismatch(t *testing.T) {
	username := "admin"
	password := "password"
	appSecret := "abcdefgh"

	PasswordChecker = func(pwd, dbpwd []byte, logger *logrus.Logger) bool {
		return false
	}

	user := &models.User{
		Username: username,
		Password: password,
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)
	expectedBody:= fmt.Sprintf(`"login: db: password didn't match for user %s"`, user.Username)

	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	mockDB := &mocks.IDBClient{}

	mockDB.On("GetUser", username).Return(user, nil)

	router := LoginHandler(logger, mockDB, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code, "invalid response code")
	assert.Equal(t, expectedBody, res.Body.String(), "invalid response body")

	mock.AssertExpectationsForObjects(t, mockDB)
}

func TestLoginHandlerWhenTokenCreationError(t *testing.T) {
	username := "admin"
	password := "password"
	appSecret := "abcdefgh"

	PasswordChecker = func(pwd, dbpwd []byte, logger *logrus.Logger) bool {
		return true
	}
	TokenCreator = func(username string, t time.Time, d time.Duration, secret string) (string, error) {
		return "", errors.New("unknown error")
	}

	user := &models.User{
		Username: username,
		Password: password,
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)
	expectedBody := `"login: create token: not able to sign token: unknown error"`

	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	mockDB := &mocks.IDBClient{}

	mockDB.On("GetUser", username).Return(user, nil)

	router := LoginHandler(logger, mockDB, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "invalid response code")
	assert.Equal(t, expectedBody, res.Body.String(), "invalid response body")

	mock.AssertExpectationsForObjects(t, mockDB)
}
