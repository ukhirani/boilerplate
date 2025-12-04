package utils

import (
	"bp/constants"
	"os"
	"path/filepath"
)

// checks whether a template exists , returns the status, the error (if occurred) and the directory where it needs to be created (if possible)
func IsTemplateExists(templateName string) (bool, error, string) {

	homeDir, _ := os.UserHomeDir()

	//destDir = homeDir + location where we store templates + templateName
	destDir := filepath.Join(homeDir, constants.BOILERPLATE_DIR, constants.TEMPLATE_DIR, templateName)

	//check if destDir directory exists
	templateExists, err := Exists(destDir)

	return templateExists, err, destDir
}
