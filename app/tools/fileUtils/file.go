package fileUtils

import (
	"os"

	"github.com/openPanel/core/app/constant"
)

// RequireDataFile opens a file in the data directory.
func RequireDataFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	dirPath := constant.DefaultDataDir
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return nil, err
	}
	return os.OpenFile(dirPath+string(os.PathSeparator)+name, flag, perm)
}
