package logger

import (
	"os"
	"time"

	"github.com/fsvxavier/golang-worker-skeleton/pkg/logger/interfaces"
	"github.com/sirupsen/logrus"
)

// Override for testing
var osHostname = os.Hostname

var (
	rootLogger *logrus.Logger
)

func initLogger() *logrus.Logger {
	rootLogger = logrus.New()
	rootLogger.SetNoLock()
	rootLogger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}
	rootLogger.SetLevel(getLogLevel("LOG_LEVEL"))
	return rootLogger
}

// GenericLogger represents log struct
type GenericLogger struct {
	rootLogger    *logrus.Logger
	Log           *logrus.Entry
	Hostname      string
	Module        string
	OperationName string
}

// NewGenericLogger create a new genericlogger
func NewGenericLogger() interfaces.GenericLogger {
	g := &GenericLogger{}
	hostname := "unknown"
	hostname, _ = g.GetHostname()
	g.Hostname = hostname
	if g.rootLogger == nil {
		g.rootLogger = initLogger()
	}
	g.Log = rootLogger.WithFields(logrus.Fields{
		"environment": os.Getenv("ENV"),
		"hostname":    hostname,
		"version":     os.Getenv("VERSION"),
		"app":         os.Getenv("APP"),
		"squad":       os.Getenv("SQUAD"),
		"tribe":       os.Getenv("TRIBE"),
	})
	return g
}

func (g *GenericLogger) SetModule(name string) {
	g.Module = name
}

func (g *GenericLogger) SetOperation(name string) {
	g.OperationName = name
}

func (g *GenericLogger) SetHostname(hostname func() (string, error)) {
	osHostname = hostname
}

func (g *GenericLogger) GetHostname() (string, error) {
	host, err := osHostname()
	if err != nil {
		return "unknown", err
	}
	return host, nil
}

// LogIt log a new message to stdout
func (g *GenericLogger) LogIt(severity, message string, fields map[string]interface{}) {
	logger := g.Log
	logger = logger.WithFields(logrus.Fields{
		"severity": severity,
		"operation": logrus.Fields{
			"name": g.OperationName,
		},
	})
	if fields != nil {
		logger = logger.WithFields(fields)
	}
	switch severity {
	case "DEBUG":
		logger.Warn(message)
	case "INFO":
		logger.Info(message)
	case "WARN":
		logger.Warn(message)
	case "ERROR":
		logger.Error(message)
	case "FATAL":
		logger.Fatal(message)
	case "PANIC":
		logger.Panic(message)
	default:
		logger.Info(message)
	}
}

func getLogLevel(envVariable string) logrus.Level {
	switch os.Getenv(envVariable) {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "FATAL":
		return logrus.FatalLevel
	case "PANIC":
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}
