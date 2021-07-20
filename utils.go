package geziyorhelper

import (
	"os"
	"path/filepath"
)

const (
	dirFileMode  = 0755
	fileFileMode = 0644
)

func createDirectory(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dir, dirFileMode); err != nil {
				return err
			}
		}
	}
	return nil
}

func saveFile(filename string, data []byte) error {
	dir := filepath.Dir(filename)
	if dir != "." {
		if err := createDirectory(dir); err != nil {
			return err
		}
	}

	return os.WriteFile(filename, data, fileFileMode)
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}