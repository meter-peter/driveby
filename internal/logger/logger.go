package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Config holds the logger configuration
type Config struct {
	// Level is the minimum log level to output
	Level string `mapstructure:"level"`
	// Format is the log format (json or text)
	Format string `mapstructure:"format"`
	// Output is the output destination (stdout, stderr, or file path)
	Output string `mapstructure:"output"`
	// Fields are additional fields to include in every log entry
	Fields map[string]interface{} `mapstructure:"fields"`
}

// DefaultConfig returns the default logger configuration
func DefaultConfig() Config {
	return Config{
		Level:  "debug",
		Format: "json",
		Output: "stdout",
		Fields: map[string]interface{}{
			"app": "driveby",
		},
	}
}

var log = logrus.New()

// Get returns the configured logger instance
func Get() *logrus.Logger {
	return log
}

// Configure sets up the logger with the given configuration
func Configure(cfg Config) error {
	// Set log level (force debug for verbose output)
	log.SetLevel(logrus.DebugLevel)
	log.Infof("Logger set to DEBUG (verbose) mode")

	// Set log format
	switch strings.ToLower(cfg.Format) {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
	case "text":
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
	default:
		return fmt.Errorf("unsupported log format: %s", cfg.Format)
	}

	// Set output
	switch strings.ToLower(cfg.Output) {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	default:
		file, err := os.OpenFile(cfg.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
		log.SetOutput(file)
	}

	// Set default fields
	if len(cfg.Fields) > 0 {
		// Create a new logger with the default fields
		newLog := logrus.New()
		newLog.SetLevel(log.GetLevel())
		newLog.SetFormatter(log.Formatter)
		newLog.SetOutput(log.Out)
		newLog.WithFields(cfg.Fields)
		log = newLog
	}

	return nil
}

// WithFields returns a new logger with the given fields
func WithFields(fields map[string]interface{}) *logrus.Entry {
	return log.WithFields(fields)
}

// WithField returns a new logger with the given field
func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

// WithError returns a new logger with the given error
func WithError(err error) *logrus.Entry {
	return log.WithError(err)
}

// Debug logs a message at level Debug
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debugf logs a formatted message at level Debug
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Info logs a message at level Info
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof logs a formatted message at level Info
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warn logs a message at level Warn
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Warnf logs a formatted message at level Warn
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Error logs a message at level Error
func Error(args ...interface{}) {
	log.Error(args...)
}

// Errorf logs a formatted message at level Error
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatal logs a message at level Fatal then the process will exit with status set to 1
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf logs a formatted message at level Fatal then the process will exit with status set to 1
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
