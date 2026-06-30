package fs

import (
	"os"
	"path/filepath"
)

func WriteFile[D []byte | string](path string, perms os.FileMode, data D, write func(file *os.File, data D) (int, error)) error {
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

	if _, err := write(file, data); err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	return nil
}

func WriteFileBytes(path string, perms os.FileMode, data []byte) error {
	return WriteFile(path, perms, data, writeBytes)
}

func WriteFileString(path string, perms os.FileMode, data string) error {
	return WriteFile(path, perms, data, writeString)
}

func writeBytes(file *os.File, data []byte) (int, error) {
	return file.Write(data)
}

func writeString(file *os.File, data string) (int, error) {
	return file.WriteString(data)
}
