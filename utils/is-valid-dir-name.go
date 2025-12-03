package utils

import (
	"os"
	"path/filepath"
)

//checks wheter the input string is a valid os directory name

func IsValidDirName(name string) bool {
	// create a temporary directory address and append it with input
	tmp := filepath.Join(os.TempDir(), name)

	//try to create that temporary directory with the input file address
	err := os.Mkdir(tmp, 0o755)

	//if error occurred then it's not a valid file name
	if err != nil {
		return false
	}

	//if yes then delete that temporary directory and return true
	os.Remove(tmp)
	return true
}
