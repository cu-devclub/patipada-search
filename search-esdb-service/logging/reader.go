package logging

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"search-esdb-service/config"
)

var totalSearch int

func CountExistingLogs() (int, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return 0, err
	}

	dirParh := filepath.Join(cwd, config.GetConfig().Static.LogsPath)

	files, err := os.ReadDir(dirParh)
	if err != nil {
		return 0, fmt.Errorf("failed to read directory: %w", err)
	}

	totalLines := 0
	for _, file := range files {
		if !file.IsDir() {
			f, err := os.Open(dirParh + "/" + file.Name())
			if err != nil {
				return 0, fmt.Errorf("failed to open file: %w", err)
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)
			lines := 0
			for scanner.Scan() {
				lines++
			}

			if err := scanner.Err(); err != nil {
				return 0, fmt.Errorf("failed to scan file: %w", err)
			}

			totalLines += lines
		}
	}

	totalSearch = totalLines

	return totalLines, nil
}

func GetTotalSearch() int {
	return totalSearch
}

func IncrementAndPrintTotalSearch() {
	totalSearch++
	slog.Info("Total Search Amount:", slog.Int("count", totalSearch))
}
