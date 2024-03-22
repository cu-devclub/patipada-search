package logging

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteLogsToFile(dir string, filedes string, log string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	directoryPath := dir + filedes

	dataDirPath := filepath.Join(cwd, directoryPath)

	file, err := os.OpenFile(dataDirPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, log)
	if err != nil {
		return err
	}
	return nil
}
