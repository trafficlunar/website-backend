package service

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"backend/internal/model"
)

func GetVisitorCounter() model.VisitorCounter {
	var data model.VisitorCounter

	path := filepath.Join(".", "data", "visitor.json")
	jsonFile, err := os.Open(path)
	if err != nil {
		slog.Warn("File not found or unable to open", slog.Any("error", err), slog.Any("file", path))
		return data
	}
	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		slog.Error("Error reading file", slog.Any("error", err), slog.Any("file", path))
		return data
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		slog.Error("Error unmarshalling JSON", slog.Any("error", err), slog.Any("file", path))
		return data
	}

	return data
}
