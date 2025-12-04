package utils

import (
	"os"
)

// calls the CopyDir function but just with destDir as the current directory
func CopyTemplateHere(templateDir string) error {

	//get the current directory and return error (if any)
	currDir, err := os.Getwd()
	if err != nil {
		return err
	}

	//CopyDir will return nil on successfull copying, thus return whatever the error is
	return CopyDir(templateDir, currDir)

}
