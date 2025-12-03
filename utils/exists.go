package utils

import (
	"errors"
	"io/fs"
	"os"
)

func Exists(path string) (bool, error) {

	// check if the path exists
	_, err := os.Stat(path)

	// if no error, then path exists
	if err == nil {
		return true, nil
	}

	// if file does not exists
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	// none of the above cases, some other kind of error
	return false, err
}
