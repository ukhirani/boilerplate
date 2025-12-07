package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func Exists(path string) bool {
	// check if the path exists
	_, err := os.Stat(path)

	// if no error, then path exists
	if err == nil {
		return true
	}

	// if file does not exists
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}

	// none of the above cases, some other kind of error
	fmt.Println("[ERROR] Unexpected error occurred : ", err)
	os.Exit(1)
	return false
}
