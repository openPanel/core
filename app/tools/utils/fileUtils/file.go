package fileUtils

import (
	"os"

	"github.com/openPanel/core/app/constant"
)

// RequireDataFile opens a file in the data directory.
func RequireDataFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	dirPath := constant.DefaultDataDir
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		if !fileInfo.IsDir() {
			return nil, os.ErrInvalid
		}
	}
	return os.OpenFile(dirPath+string(os.PathSeparator)+name, flag, perm)
}

// RequireDataDir creates a directory in the data directory.
func RequireDataDir(dir string) (string, error) {
	dirPath := constant.DefaultDataDir + string(os.PathSeparator) + dir
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", err
	}
	return dirPath, nil
}
