package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var logger *slog.Logger = slog.New(
	slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: slog.LevelInfo},
	),
)

// Init initializes the logger
func Init(encoding, levelStr string) {
	level := slog.LevelDebug
	if levelStr == "info" {
		level = slog.LevelInfo
	}

	if encoding == "console" {
		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: level},
			),
		)
	} else {
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: level},
			),
		)
	}

	slog.SetDefault(logger)
}

var requestIDKey = "request_id"

// Info logs an info message
func Info(ctx context.Context, msg string, fields ...any) {
	fields = includeRequestID(ctx, fields...)
	logger.Info(msg, fields...)
}

// Debug logs a debug message
func Debug(ctx context.Context, msg string, fields ...any) {
	fields = includeRequestID(ctx, fields...)
	logger.Debug(msg, fields...)
}

// Error logs an error message
func Error(ctx context.Context, msg string, err error, fields ...any) {
	fields = append(fields, slog.Any("err", fmt.Sprintf("%+v", err)))
	fields = includeRequestID(ctx, fields...)
	logger.Error(msg, fields...)
}

// Warn logs a warning message
func Warn(ctx context.Context, msg string, fields ...any) {
	fields = includeRequestID(ctx, fields...)
	logger.Warn(msg, fields...)
}

// includeRequestID includes request ID in the log message if it exists
func includeRequestID(ctx context.Context, fields ...any) []any {
	requestID := extractRequestID(ctx)
	if requestID == "" {
		return fields
	}

	return append(fields, slog.Any(requestIDKey, requestID))
}

func extractRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	id, _ := ctx.Value("request_id").(string)
	return id
}
