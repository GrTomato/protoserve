package main

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {

	logger := slog.New(
		slog.NewJSONHandler(
			os.Stderr,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)

	return logger
}
