package gcp

import (
	"compress/gzip"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rasberry-robot/user-api/pkg/logging"
	"github.com/rasberry-robot/user-api/pkg/myerr"

	gcs "cloud.google.com/go/storage"
	"github.com/avast/retry-go"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

// CloudStorage -
type CloudStorage struct {
	Client *gcs.Client
}

// NewCloudStorage はCloud Storageのクライアントを作成する
func NewCloudStorage(ctx context.Context, projectID string) (*CloudStorage, error) {
	client, err := gcs.NewClient(ctx, option.WithCredentialsJSON([]byte(os.Getenv("GCS_CREDENTIALS_JSON"))))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gcs client. %w", err)
	}
	return &CloudStorage{Client: client}, nil
}

// Put はCloud Storageにデータを保存する
func (cs *CloudStorage) Put(ctx context.Context, bucket string, key string, data [][]string) error {
	if err := retry.Do(
		func() error {
			var (
				pReader, pWriter = io.Pipe()
				object           = cs.Client.Bucket(bucket).Object(key).NewWriter(ctx)
				zw               = gzip.NewWriter(object)
			)
			defer pReader.Close()
			go func() {
				defer pWriter.Close()
				csv.NewWriter(pWriter).WriteAll(append( //nolint:errcheck
					[][]string{},
					data...,
				))
			}()
			if _, err := io.Copy(zw, pReader); err != nil {
				return fmt.Errorf("failed to pipe csv to gzip writer. %w", err)
			}
			if err := zw.Close(); err != nil {
				return fmt.Errorf("failed to close gzip writer. %w", err)
			}
			if err := object.Close(); err != nil {
				return fmt.Errorf("failed to close object. %w", err)
			}
			// TODO: ファイルサイズをlogに突っ込む
			// uploadedSize = object.Attrs().Size
			return nil
		},
		retry.RetryIf(func(err error) bool {
			if gcsError, ok := myerr.As[*googleapi.Error](err); ok {
				return gcsError.Code != 403 && gcsError.Code != 404
			}
			return !errors.Is(err, context.Canceled)
		}),
		retry.OnRetry(func(n uint, err error) {
			logging.SystemWarn(ctx, "finished retrying saving content to gcs. retried count:[%d] key:[%s] error:[%s]", n+1, key, err)
		}),
		retry.Context(ctx),
		retry.Attempts(3),
		retry.Delay(5*time.Second),
		retry.DelayType(retry.BackOffDelay),
		retry.LastErrorOnly(true),
	); err != nil {
		return fmt.Errorf("max retry count is exceeded. %w", cs.handleErrorResponse(err, bucket, key))
	}
	return nil
}

// handleErrorResponse は GCS のエラーをハンドリングする
func (*CloudStorage) handleErrorResponse(err error, bucket string, key string) error {
	gcsError, ok := myerr.As[*googleapi.Error](err)
	if !ok {
		return fmt.Errorf("bucket:[%s] key:[%s]. %w", bucket, key, err)
	}

	switch gcsError.Code {
	case 403:
		return fmt.Errorf("bucket:[%s] key:[%s]. %w", bucket, key, err)
	case 404:
		return fmt.Errorf("bucket:[%s] key:[%s]. %w", bucket, key, err)
	}
	return fmt.Errorf("bucket:[%s] key:[%s]. %w", bucket, key, err)
}
