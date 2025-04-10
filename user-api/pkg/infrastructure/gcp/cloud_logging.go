package gcp

import (
	"context"
	"fmt"
	"time"

	cloudLogging "cloud.google.com/go/logging"
)

// CloudLoggingClient はCloud Loggingのクライアントを表す
type CloudLoggingClient struct {
	Client *cloudLogging.Client
}

// NewCloudLoggingClient はCloud Loggingのクライアントを作成する
func NewCloudLoggingClient(ctx context.Context, projectID string) (*CloudLoggingClient, error) {
	client, err := cloudLogging.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create logging client: %w", err)
	}
	return &CloudLoggingClient{Client: client}, nil
}

// Log はログをCloud Loggingに書き込む
func (c *CloudLoggingClient) Log(ctx context.Context, logName string, severity cloudLogging.Severity, message string) error {
	logger := c.Client.Logger(logName)
	entry := cloudLogging.Entry{
		Timestamp: time.Now(),
		Severity:  severity,
		Payload:   map[string]interface{}{"message": message},
	}
	logger.Log(entry)
	return nil
}
