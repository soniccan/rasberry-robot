package logging

import "context"

// Logger はloggerのinterface
type Logger interface {
	Info(ctx context.Context, format string, args ...any)
	Warn(ctx context.Context, format string, args ...any)
	Error(ctx context.Context, format string, args ...any)
	InfoWithTag(ctx context.Context, tag Tag, meta map[string]any, format string, args ...any)
	WarnWithTag(ctx context.Context, tag Tag, meta map[string]any, format string, args ...any)
	ErrorWithTag(ctx context.Context, tag Tag, meta map[string]any, format string, args ...any)
	InfoFileStorage(ctx context.Context, tag Tag, meta map[string]any, format string, args ...any)
}

type StreamingLogger interface {
	Info(ctx context.Context, body map[string]any) error
}

// Tag はlogに付与するタグ
type Tag string

// Tags
const (
	SystemInfoTag  Tag = "batch.system.info"
	SystemWarnTag  Tag = "batch.system.warn"
	SystemErrorTag Tag = "batch.system.error"
)
