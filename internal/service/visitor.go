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
	path := filepath.Join(".", "data", "visitor.json")
	jsonFile, err := os.Open(path)
	if err != nil {
		slog.Warn("File not found!", slog.Any("file", path))
		return model.VisitorCounter{}
	}

	bytes, _ := io.ReadAll(jsonFile)

	var data model.VisitorCounter
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		slog.Error("Error unmarshalling JSON", slog.Any("error", err), slog.Any("file", path))
	}

	return data
}
