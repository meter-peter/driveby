package types

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger handles validation report logging
type Logger struct {
	logger *logrus.Logger
	path   string
}

// NewLogger creates a new validation logger
func NewLogger(logPath string) (*Logger, error) {
	if logPath == "" {
		logger := logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetLevel(logrus.DebugLevel)
		logger.SetOutput(os.Stdout)
		logger.Infof("[types/logger] Logger set to DEBUG (verbose) mode (stdout fallback)")
		return &Logger{
			logger: logger,
			path:   "",
		}, nil
	}
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	logger.Infof("[types/logger] Logger set to DEBUG (verbose) mode")

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logger.SetOutput(file)

	return &Logger{
		logger: logger,
		path:   logPath,
	}, nil
}

// LogReport logs a validation report
func (l *Logger) LogReport(report *ValidationReport) error {
	report.Timestamp = time.Now()

	entry := l.logger.WithFields(logrus.Fields{
		"type":      "validation_report",
		"version":   report.Version,
		"env":       report.Environment,
		"timestamp": report.Timestamp,
	})

	// Log the full report as JSON
	reportJSON, err := json.Marshal(report)
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	entry.Info(string(reportJSON))
	return nil
}

// GetRecentReports retrieves recent validation reports
func (l *Logger) GetRecentReports(limit int) ([]ValidationReport, error) {
	// Implementation to read and parse recent reports from the log file
	// This would be implemented based on your specific needs
	return nil, nil
}
