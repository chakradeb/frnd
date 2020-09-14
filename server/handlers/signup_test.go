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

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/chakradeb/frnd-server/mocks"
	"github.com/chakradeb/frnd-server/models"
)

func TestSignupHandler(t *testing.T) {
	username := "admin"
	email := "admin@dc.org"
	password := "password"
	appSecret := "abcdefgh"

	Encrypter = func(pwd []byte) (string, error) {
		return string(pwd), nil
	}
	TokenCreator = func(username string, t time.Time, secret string) (string, error) {
		return secret, nil
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Gender:   "male",
	}
	token := &models.Session{
		Username: username,
		Token: appSecret,
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)
	expectedToken, _ := json.Marshal(token)

	req := httptest.NewRequest("POST", "/api/signup", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	mockDB := &mocks.IDBClient{}

	mockDB.On("GetUser", username).Return(&models.User{}, errors.New("no user"))
	mockDB.On("CreateUser", username, password).Return(nil)

	router := SignupHandler(logger, mockDB, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code, "invalid response code")
	assert.JSONEq(t, string(expectedToken), res.Body.String(), "invalid response body")

	mock.AssertExpectationsForObjects(t, mockDB)
}

func TestSignupHandlerWithEmptyData(t *testing.T) {
	username := ""
	email := ""
	password := ""
	appSecret := "abcdefgh"

	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Gender:   "male",
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/api/signup", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	router := SignupHandler(logger, nil, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code, "invalid response code")
	assert.Equal(t, `"signup: fields should not be empty"`, res.Body.String(), "invalid response body")
}

func TestSignupHandlerWhenUserAlreadyExists(t *testing.T) {
	username := "admin"
	email := "admin@dc.org"
	password := "password"
	appSecret := "abcdefgh"

	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Gender:   "male",
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)
	expectedBody := fmt.Sprintf(`"signup: db: user %s already exist"`, username)

	req := httptest.NewRequest("POST", "/api/signup", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	mockDB := &mocks.IDBClient{}

	mockDB.On("GetUser", username).Return(&models.User{}, nil)

	router := SignupHandler(logger, mockDB, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusForbidden, res.Code, "invalid response code")
	assert.Equal(t, expectedBody, res.Body.String(), "invalid response body")

	mock.AssertExpectationsForObjects(t, mockDB)
}

func TestSignupHandlerWhenEncryptionError(t *testing.T) {
	username := "admin"
	email := "admin@dc.org"
	password := "password"
	appSecret := "abcdefgh"

	Encrypter = func(pwd []byte) (string, error) {
		return "", errors.New("encrypt: error while encryption: unknown error")
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Gender:   "male",
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)
	expectedBody:= `"signup: encrypt: error while encryption: unknown error"`

	req := httptest.NewRequest("POST", "/api/signup", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	mockDB := &mocks.IDBClient{}

	mockDB.On("GetUser", username).Return(&models.User{}, errors.New("no user"))

	router := SignupHandler(logger, mockDB, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "invalid response code")
	assert.Equal(t, expectedBody, res.Body.String(), "invalid response body")

	mock.AssertExpectationsForObjects(t, mockDB)
}

func TestSignupHandlerWhenDBError(t *testing.T) {
	username := "admin"
	email := "admin@dc.org"
	password := "password"
	appSecret := "abcdefgh"

	Encrypter = func(pwd []byte) (string, error) {
		return string(pwd), nil
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Gender:   "male",
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)
	expectedBody := `"signup: db: can not save user details: db error"`

	req := httptest.NewRequest("POST", "/api/signup", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	mockDB := &mocks.IDBClient{}

	mockDB.On("GetUser", username).Return(&models.User{}, errors.New("no user"))
	mockDB.On("CreateUser", username, password).Return(errors.New("db error"))

	router := SignupHandler(logger, mockDB, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "invalid response code")
	assert.Equal(t, expectedBody, res.Body.String(), "invalid response body")

	mock.AssertExpectationsForObjects(t, mockDB)
}

func TestSignupHandlerWhenTokenCreationError(t *testing.T) {
	username := "admin"
	email := "admin@dc.org"
	password := "password"
	appSecret := "abcdefgh"

	Encrypter = func(pwd []byte) (string, error) {
		return string(pwd), nil
	}
	TokenCreator = func(username string, t time.Time, secret string) (string, error) {
		return "", errors.New("unknown error")
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Gender:   "male",
	}

	logger, hook := test.NewNullLogger()
	defer hook.Reset()

	body, _ := json.Marshal(user)
	expectedBody := `"signup: create token: not able to sign token: unknown error"`

	req := httptest.NewRequest("POST", "/api/signup", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	mockDB := &mocks.IDBClient{}

	mockDB.On("GetUser", username).Return(&models.User{}, errors.New("no user"))
	mockDB.On("CreateUser", username, password).Return(nil)

	router := SignupHandler(logger, mockDB, appSecret)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "invalid response code")
	assert.Equal(t, expectedBody, res.Body.String(), "invalid response body")

	mock.AssertExpectationsForObjects(t, mockDB)
}
