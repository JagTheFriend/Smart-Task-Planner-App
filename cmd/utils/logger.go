package utils

import (
	"log/slog"
	"os"

	"github.com/MatusOllah/slogcolor"
)

func SetupLogger() {
	logger := slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions))
	slog.SetDefault(logger)
}
