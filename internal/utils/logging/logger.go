package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger extends the logrus.Logger to add additional methods
type Logger struct {
	*logrus.Logger
}

// NewLogger creates a new logger with default settings
func NewLogger() *Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
	})
	logger.SetLevel(logrus.InfoLevel)

	return &Logger{Logger: logger}
}

// Configure configures the logger with the given level and format
func (l *Logger) Configure(level, format string) {
	// Set the log level
	switch level {
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "info":
		l.SetLevel(logrus.InfoLevel)
	case "warn":
		l.SetLevel(logrus.WarnLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	default:
		l.SetLevel(logrus.InfoLevel)
	}

	// Set the log format
	switch format {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
		})
	case "text":
		l.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
		})
	default:
		l.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
		})
	}
}

// SetOutput sets the output destination for the logger
func (l *Logger) SetOutput(output io.Writer) {
	l.Logger.SetOutput(output)
}

// WithComponent adds a component field to the logger
func (l *Logger) WithComponent(component string) *logrus.Entry {
	return l.WithField("component", component)
}

// WithService adds a service field to the logger
func (l *Logger) WithService(service string) *logrus.Entry {
	return l.WithField("service", service)
}

// WithRequest adds request-specific fields to the logger
func (l *Logger) WithRequest(requestID, method, path string) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"request_id": requestID,
		"method":     method,
		"path":       path,
	})
}

// WithTest adds test-specific fields to the logger
func (l *Logger) WithTest(testID, testType string) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"test_id":   testID,
		"test_type": testType,
	})
}

// WithError logs an error with the given message
func (l *Logger) WithError(err error) *logrus.Entry {
	return l.Logger.WithError(err)
}