package service

import (
	"backend/internal/model"
	"bufio"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetStatuses() model.StatusData {
	data := model.StatusData{
		Success: false,
		Website: 0,
		Api:     0,
		Files:   0,
	}

	req, err := http.NewRequest("GET", os.Getenv("UPTIME_KUMA_URL"), nil)
	if err != nil {
		slog.Error("Error creating request for Uptime Kuma", slog.Any("error", err))
		return data
	}
	req.SetBasicAuth("", os.Getenv("UPTIME_KUMA_API_KEY"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		slog.Error("Error sending request for Uptime Kuma", slog.Any("error", err))
		return data
	}
	defer res.Body.Close()

	regex := regexp.MustCompile(`monitor_name="([^"]+)"[^}]*} (\d+)`)
	statusMap := map[string]*uint8{
		"website": &data.Website,
		"api":     &data.Api,
		"files":   &data.Files,
	}

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "monitor_status{") {
			matches := regex.FindStringSubmatch(line)
			monitorName := matches[1]
			status := matches[2]

			statusCode, err := strconv.Atoi(status)
			if err != nil {
				slog.Error("Error parsing bool for Uptime Kuma", slog.Any("error", err))
				statusCode = 0
			}

			if field, exists := statusMap[monitorName]; exists {
				*field = uint8(statusCode)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		slog.Error("Error reading metrics for Uptime Kuma", slog.Any("error", err))
		return data
	}

	data.Success = true
	return data
}
