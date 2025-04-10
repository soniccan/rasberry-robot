package logging

import "context"

var (
	logger          Logger
	streamingLogger StreamingLogger
)

// SetLogger は logger をセットする
func SetLogger(l Logger) { logger = l }

// SetStreamingLogger は streamingLogger をセットする
func SetStreamingLogger(l StreamingLogger) { streamingLogger = l }

// SystemInfo -
func SystemInfo(ctx context.Context, format string, args ...any) {
	logger.Info(ctx, format, args...)
}

// SystemWarn -
func SystemWarn(ctx context.Context, format string, args ...any) {
	logger.Warn(ctx, format, args...)
}

// SystemError -
func SystemError(ctx context.Context, format string, args ...any) {
	logger.Error(ctx, format, args...)
}
