package fs

import (
	"os"
	"path/filepath"
)

func WriteFile(data []byte, perms os.FileMode, pathSegments ...string) error {
	path := filepath.Join(pathSegments...)
	return os.WriteFile(path, data, perms)
}

func Write(path string, data []byte, perms os.FileMode) error {
	if dirErr := os.MkdirAll(filepath.Dir(path), perms); dirErr != nil {
		return dirErr
	}

	return os.WriteFile(path, data, perms)
}

func WriteString(path string, data string, perms os.FileMode) error {
	if dirErr := os.MkdirAll(filepath.Dir(path), perms); dirErr != nil {
		return dirErr
	}

	file, fileErr := os.Create(path)

	if fileErr != nil {
		return fileErr
	}

	defer file.Close()

	if permErr := file.Chmod(perms); permErr != nil {
		return permErr
	}

	if _, writeErr := file.WriteString(data); writeErr != nil {
		return writeErr
	}

	if syncErr := file.Sync(); syncErr != nil {
		return syncErr
	}

	return nil
}
