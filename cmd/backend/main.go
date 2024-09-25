package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"

	"backend/internal/server"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", slog.Any("error", err))
	}

	server.NewRouter()
}
