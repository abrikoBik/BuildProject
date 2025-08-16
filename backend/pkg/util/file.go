package util

import (
	"os"
	"path/filepath"
)

func MakeStoragePath(name string) (string, error) {
	currentPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	attachmentPath := filepath.Join(filepath.Dir(currentPath), StoragePath, name)
	return attachmentPath, nil
}
