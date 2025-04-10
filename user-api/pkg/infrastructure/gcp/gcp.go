package gcp

import (
	"context"
	"fmt"
)

type GCP struct {
	bigQuery     *BigQuery
	cloudLogging *CloudLoggingClient
	gcs          *GCS
}

// NewGCP はGCPのクライアントを作成する
func NewGCP(ctx context.Context, projectID string) (*GCP, error) {
	gcs, err := NewGCS(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create gcs client: %w", err)
	}
	bigQuery, err := NewBigQuery(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create bigquery client: %w", err)
	}
	cloudLogging, err := NewCloudLogging(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloud logging client: %w", err)
	}

	return &GCP{
		bigQuery:     bigQuery,
		cloudLogging: cloudLogging,
		gcs:          gcs,
	}, nil
}

// GCS はGCSのクライアントを返す
func (g *GCP) GCS() *GCS {
	return g.gcs
}

// BigQuery はBigQueryのクライアントを返す
func (g *GCP) BigQuery() *BigQuery {
	return g.bigQuery
}

// CloudLogging はCloudLoggingのクライアントを返す
func (g *GCP) CloudLogging() *CloudLoggingClient {
	return g.cloudLogging
}
