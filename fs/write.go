package fs

import (
	"os"
	"path/filepath"
)

func WriteBytes(path string, data []byte, perms os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), perms); err != nil {
		return err
	}

	file, fileErr := os.Create(path)

	if fileErr != nil {
		return fileErr
	}

	defer file.Close()

	if err := file.Chmod(perms); err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	return nil
}

func WriteString(path string, data string, perms os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), perms); err != nil {
		return err
	}

	file, fileErr := os.Create(path)

	if fileErr != nil {
		return fileErr
	}

	defer file.Close()

	if err := file.Chmod(perms); err != nil {
		return err
	}

	if _, err := file.WriteString(data); err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	return nil
}
