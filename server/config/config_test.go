package config

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_ = os.Setenv("PORT", "5000")
	_ = os.Setenv("LOG_LEVEL", "info")
	defer os.Clearenv()

	conf, errs := New()

	assert.Equal(t, 0, len(errs), "unexpected error while parsing")
	assert.Equal(t, 5000, conf.AppPort, "invalid value for app port")
	assert.Equal(t, logrus.InfoLevel, conf.LogLevel, "invalid value for log level")
}

func TestNewWithInvalidPortValues(t *testing.T) {
	_ = os.Setenv("PORT", "abcd")
	_ = os.Setenv("LOG_LEVEL", "info")
	defer os.Clearenv()

	_, errs := New()

	assert.Equal(t, 1, len(errs), "unexpected error while parsing")
	assert.Equal(t, errors.New("env: strconv.ParseInt: parsing \"abcd\": invalid syntax"), errs[0], "invalid error while parsing")
}

func TestNewWithInvalidLogLevelValue(t *testing.T) {
	_ = os.Setenv("PORT", "5000")
	_ = os.Setenv("LOG_LEVEL", "all")
	defer os.Clearenv()

	_, errs := New()

	assert.Equal(t, 1, len(errs), "unexpected error while parsing")
	assert.Equal(t, errors.New("config: not a valid logrus Level: \"all\""), errs[0], "invalid error while parsing")
}

func TestNewWithDefaultValue(t *testing.T) {
	_ = os.Setenv("LOG_LEVEL", "debug")
	defer os.Clearenv()

	conf, errs := New()

	assert.Equal(t, 0, len(errs), "unexpected error while parsing")
	assert.Equal(t, 8000, conf.AppPort, "invalid value for app port")
	assert.Equal(t, logrus.DebugLevel, conf.LogLevel, "invalid value for log level")
}

func TestConfig_ShowConfig(t *testing.T) {
	_ = os.Setenv("LOG_LEVEL", "error")
	defer os.Clearenv()

	conf, errs := New()

	assert.Equal(t, 0, len(errs), "unexpected error while parsing")

	fields := conf.ShowConfig()

	assert.Equal(t, 8000, fields["AppPort"], "invalid field value for app port")
	assert.Equal(t, logrus.ErrorLevel, fields["LogLevel"], "invalid field value for log level")
}
