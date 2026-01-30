package logger

import (
	"effective_mobile/internal/config"
	"log/slog"
	"os"
	"strings"
)

func New(conf *config.Logger) *slog.Logger {
	var log *slog.Logger

	var l slog.Level
	switch strings.ToLower(conf.Level) {
	case "error":
		l = slog.LevelError
	case "warn":
		l = slog.LevelWarn
	case "debug":
		l = slog.LevelDebug
	default:
		l = slog.LevelInfo
	}

	log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l}),
	)

	return log
}
