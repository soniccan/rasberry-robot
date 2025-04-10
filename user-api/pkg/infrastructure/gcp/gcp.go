package gcp

import (
	"context"
)

type GCP struct {
	CloudStorage *CloudStorage
	BigQuery     *BigQuery
}

// NewGCP はGCPのクライアントを作成する
func NewGCP(ctx context.Context, projectID string) (*GCP, error) {
	cloudStorage, err := NewCloudStorage(ctx, projectID)
	if err != nil {
		return nil, err
	}
	bigQuery, err := NewBigQuery(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &GCP{
		CloudStorage: cloudStorage,
		BigQuery:     bigQuery,
	}, nil
}
