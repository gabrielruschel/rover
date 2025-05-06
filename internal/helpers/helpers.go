package helpers

import (
	"log/slog"
	"os"
	"strings"
)

func NewLogger(level string, attrs ...any) *slog.Logger {
	opts := &slog.HandlerOptions{}
	switch strings.ToUpper(level) {
	case "DEBUG":
		opts.Level = slog.LevelDebug
	case "INFO":
		opts.Level = slog.LevelInfo
	case "ERROR":
		opts.Level = slog.LevelError
	default:
		opts.Level = slog.LevelWarn
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, opts)).With(attrs...)
}
