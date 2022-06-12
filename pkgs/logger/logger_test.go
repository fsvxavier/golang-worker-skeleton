package logger_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/fsvxavier/golang-worker-skeleton/pkg/logger"
	"github.com/stretchr/testify/assert"
)

// Override for testing
var osHostname = os.Hostname

func TestLogger(t *testing.T) {
	t.Run("success-severity", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "DEBUG")

		logger := logger.NewGenericLogger()
		logger.SetModule("worker")
		logger.SetOperation("Initialize")

		logger.LogIt("DEBUG", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)
		logger.LogIt("INFO", fmt.Sprintf("[test] - Info: %s", "teste"), nil)
		logger.LogIt("WARN", fmt.Sprintf("[test] - Warn: %s", "teste"), nil)
		logger.LogIt("ERROR", fmt.Sprintf("[test] - Error: %s", "teste"), nil)

		assert.NotNil(t, logger)
	})
	t.Run("success-loglevel-debug", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "INFO")

		logger := logger.NewGenericLogger()
		logger.SetModule("worker")
		logger.SetOperation("Initialize")

		logger.LogIt("DEBUG", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logger)
	})
	t.Run("success-loglevel-info", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "WARN")

		logger := logger.NewGenericLogger()
		//logger.GetLogger()
		logger.SetModule("worker")
		logger.SetOperation("Initialize")

		logger.LogIt("INFO", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logger)
	})
	t.Run("success-loglevel-warn", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "ERROR")

		logger := logger.NewGenericLogger()
		logger.SetModule("worker")
		logger.SetOperation("Initialize")

		logger.LogIt("WARN", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logger)
	})

	t.Run("success-loglevel-default", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "")

		logger := logger.NewGenericLogger()
		logger.SetModule("worker")
		logger.SetOperation("Initialize")

		logger.LogIt("", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logger)
	})

	t.Run("success-logrusloglevel-fatal", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "FATAL")

		logger := logger.NewGenericLogger()
		logger.SetModule("worker")
		logger.SetOperation("Initialize")

		logger.LogIt("", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logger)
	})

	t.Run("success-logrusloglevel-panic", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "PANIC")

		logger := logger.NewGenericLogger()
		logger.SetModule("worker")
		logger.SetOperation("Initialize")

		logger.LogIt("", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)

		assert.NotNil(t, logger)
	})

	t.Run("success-field", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "")

		field := map[string]interface{}{"teste": ""}

		logger := logger.NewGenericLogger()
		logger.SetModule("worker")
		logger.SetOperation("Initialize")

		logger.LogIt("WARN", fmt.Sprintf("[test] - Debug: %s", "teste"), field)

		assert.NotNil(t, logger)
	})

	t.Run("error-root", func(t *testing.T) {
		os.Setenv("ENV", "teste")
		os.Setenv("SQUAD", "teste")
		os.Setenv("TRIBE", "teste")
		os.Setenv("APP", "teste")
		os.Setenv("LOG_LEVEL", "DEBUG")

		logger := logger.NewGenericLogger()
		logger.SetModule("worker")
		logger.SetOperation("Initialize")

		logger.LogIt("WARN", fmt.Sprintf("[test] - Debug: %s", "teste"), nil)
		logger = nil
		assert.Nil(t, logger)
	})
}
func TestGetHostnameFails(t *testing.T) {
	defer func() { osHostname = os.Hostname }()
	os.Setenv("ENV", "teste")
	os.Setenv("SQUAD", "teste")
	os.Setenv("TRIBE", "teste")
	os.Setenv("APP", "teste")
	os.Setenv("LOG_LEVEL", "DEBUG")

	logger := logger.NewGenericLogger()
	logger.SetModule("worker")
	logger.SetOperation("Initialize")
	osHostname = func() (string, error) { return "", errors.New("fail") }
	logger.SetHostname(osHostname)
	hostname, err := logger.GetHostname()
	if err == nil {
		t.Errorf("getHostname() = (%v, nil), want error", hostname)
	}
	hostname = "unknown"
}
