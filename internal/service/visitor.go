package service

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"backend/internal/model"
)

const path = "./data/visit.json"

func GetVisitCounter() model.VisitCounter {
	var data model.VisitCounter

	jsonFile, err := os.Open(path)
	if err != nil {
		slog.Warn("File not found or unable to open", slog.Any("error", err), slog.Any("path", path))
		return data
	}
	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		slog.Error("Error reading file", slog.Any("error", err), slog.Any("path", path))
		return data
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		slog.Error("Error unmarshalling JSON", slog.Any("error", err), slog.Any("path", path))
		return data
	}

	return data
}

func IncrementVisitCounter() model.Success {
	data := GetVisitCounter()
	data.Counter++

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		slog.Error("Unable to create directory", slog.Any("error", err), slog.Any("path", filepath.Dir(path)))
		return model.Success{}
	}

	jsonString, err := json.Marshal(data)
	if err != nil {
		slog.Error("Error marshalling JSON", slog.Any("error", err), slog.Any("path", path))
		return model.Success{}
	}

	err = os.WriteFile(path, jsonString, 0644)
	if err != nil {
		slog.Error("Error writing to file", slog.Any("error", err), slog.Any("path", path))
		return model.Success{}
	}

	return model.Success{
		Success: true,
	}
}
