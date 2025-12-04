package utils

import (
	"os"
)

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	//check whether it's a directory
	if fileInfo.IsDir() {
		// return true and no error if it's a directory
		return true, nil
	} else if fileInfo.Mode().IsRegular() {
		// return false and no error if it's a file
		return false, nil
	}

	//any other case
	return false, err
}
