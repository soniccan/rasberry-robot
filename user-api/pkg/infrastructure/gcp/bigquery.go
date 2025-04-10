package gcp

import (
	"context"

	bigquery "cloud.google.com/go/bigquery"
)

type BigQuery struct {
	Client *bigquery.Client
}

func NewBigQuery(ctx context.Context, projectID string) (*BigQuery, error) {
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &BigQuery{Client: client}, nil
}
