package data

import (
	"io/fs"
	"os"
	"path/filepath"
	"search-esdb-service/config"
)

func GetRecordCSVFilesEntry(cfg *config.Config) ([]fs.DirEntry, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	directoryPath := cfg.Static.DataPath + cfg.Static.RecordPath

	dataDirPath := filepath.Join(cwd, directoryPath)

	dir, err := os.ReadDir(dataDirPath)
	if err != nil {
		return nil, err
	}

	return dir, nil
}
