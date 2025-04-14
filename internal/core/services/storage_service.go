package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/example/driveby/internal/config"
	"github.com/example/driveby/internal/core/models"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

// MinioStorageService implements the StorageService interface using Minio
type MinioStorageService struct {
	config *config.Config
	logger *logrus.Logger
	client *minio.Client
}

// NewMinioStorageService creates a new Minio storage service
func NewMinioStorageService(cfg *config.Config, logger *logrus.Logger) (StorageService, error) {
	// Create Minio client
	client, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: cfg.Minio.UseSSL,
		Region: cfg.Minio.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Minio client: %w", err)
	}

	service := &MinioStorageService{
		config: cfg,
		logger: logger,
		client: client,
	}

	// Ensure bucket exists
	if err := service.ensureBucket(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ensure bucket exists: %w", err)
	}

	logger.WithField("endpoint", cfg.Minio.Endpoint).Info("Minio storage service initialized")
	return service, nil
}

// ensureBucket ensures that the configured bucket exists
func (s *MinioStorageService) ensureBucket(ctx context.Context) error {
	bucketName := s.config.Minio.BucketName
	exists, err := s.client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to check if bucket exists: %w", err)
	}

	if !exists {
		s.logger.WithField("bucket", bucketName).Info("Creating bucket")
		if err := s.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
			Region: s.config.Minio.Region,
		}); err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
	}

	return nil
}

// buildObjectKey builds the object key for a test
func (s *MinioStorageService) buildObjectKey(testType models.TestType, testID string) string {
	return fmt.Sprintf("tests/%s/%s/test.json", testType, testID)
}

// buildReportKey builds the object key for a test report
func (s *MinioStorageService) buildReportKey(testType models.TestType, testID string) string {
	// Use a timestamp to ensure uniqueness
	timestamp := time.Now().Format("20060102-150405")
	return fmt.Sprintf("reports/%s/%s/%s-report.md", testType, testID, timestamp)
}

// SaveTest saves a test to storage
func (s *MinioStorageService) SaveTest(ctx context.Context, testType models.TestType, testID string, data interface{}) error {
	// Marshal data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal test data: %w", err)
	}

	// Build object key
	objectKey := s.buildObjectKey(testType, testID)

	// Upload to Minio
	_, err = s.client.PutObject(ctx, s.config.Minio.BucketName, objectKey, bytes.NewReader(jsonData), int64(len(jsonData)),
		minio.PutObjectOptions{
			ContentType: "application/json",
		})
	if err != nil {
		return fmt.Errorf("failed to upload test data: %w", err)
	}

	s.logger.WithFields(logrus.Fields{
		"test_type": testType,
		"test_id":   testID,
		"object":    objectKey,
	}).Info("Test saved to storage")

	return nil
}

// GetTest retrieves a test from storage
func (s *MinioStorageService) GetTest(ctx context.Context, testType models.TestType, testID string, result interface{}) error {
	// Build object key
	objectKey := s.buildObjectKey(testType, testID)

	// Get object from Minio
	obj, err := s.client.GetObject(ctx, s.config.Minio.BucketName, objectKey, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to get test data: %w", err)
	}
	defer obj.Close()

	// Read object
	jsonData, err := io.ReadAll(obj)
	if err != nil {
		return fmt.Errorf("failed to read test data: %w", err)
	}

	// Unmarshal JSON
	if err := json.Unmarshal(jsonData, result); err != nil {
		return fmt.Errorf("failed to unmarshal test data: %w", err)
	}

	s.logger.WithFields(logrus.Fields{
		"test_type": testType,
		"test_id":   testID,
		"object":    objectKey,
	}).Info("Test retrieved from storage")

	return nil
}

// ListTests retrieves all tests of a specific type from storage
func (s *MinioStorageService) ListTests(ctx context.Context, testType models.TestType) ([]string, error) {
	// Build prefix
	prefix := fmt.Sprintf("tests/%s/", testType)

	// List objects with prefix
	var testIDs []string
	opts := minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true,
	}

	for object := range s.client.ListObjects(ctx, s.config.Minio.BucketName, opts) {
		if object.Err != nil {
			return nil, fmt.Errorf("failed to list objects: %v", object.Err)
		}

		// Extract test ID from object key
		// Example: tests/validation/123456/test.json -> 123456
		key := object.Key
		parts := strings.Split(key, "/")
		if len(parts) >= 3 && parts[0] == "tests" && parts[len(parts)-1] == "test.json" {
			testIDs = append(testIDs, parts[2])
		}
	}

	s.logger.WithFields(logrus.Fields{
		"test_type": testType,
		"count":     len(testIDs),
	}).Info("Tests listed from storage")

	return testIDs, nil
}

// SaveReport saves a test report to storage
func (s *MinioStorageService) SaveReport(ctx context.Context, testType models.TestType, testID string, reportContent string) (string, error) {
	// Generate unique report key
	reportKey := s.buildReportKey(testType, testID)

	// Convert report content to bytes
	contentBytes := []byte(reportContent)

	// Upload to Minio
	_, err := s.client.PutObject(ctx, s.config.Minio.BucketName, reportKey, bytes.NewReader(contentBytes), int64(len(contentBytes)),
		minio.PutObjectOptions{
			ContentType: "text/markdown",
		})
	if err != nil {
		return "", fmt.Errorf("failed to upload report: %w", err)
	}

	s.logger.WithFields(logrus.Fields{
		"test_type": testType,
		"test_id":   testID,
		"report":    reportKey,
	}).Info("Report saved to storage")

	// Return the report path (can be used to generate a URL)
	return reportKey, nil
}

// GetReport retrieves a test report from storage
func (s *MinioStorageService) GetReport(ctx context.Context, reportPath string) (string, error) {
	// Get object from Minio
	obj, err := s.client.GetObject(ctx, s.config.Minio.BucketName, reportPath, minio.GetObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get report: %w", err)
	}
	defer obj.Close()

	// Read object
	reportBytes, err := io.ReadAll(obj)
	if err != nil {
		return "", fmt.Errorf("failed to read report: %w", err)
	}

	s.logger.WithField("report", reportPath).Info("Report retrieved from storage")

	return string(reportBytes), nil
}

// GeneratePublicURL generates a pre-signed URL for a report
func (s *MinioStorageService) GeneratePublicURL(ctx context.Context, objectPath string, expiry time.Duration) (string, error) {
	// Generate presigned URL
	url, err := s.client.PresignedGetObject(ctx, s.config.Minio.BucketName, objectPath, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return url.String(), nil
}
