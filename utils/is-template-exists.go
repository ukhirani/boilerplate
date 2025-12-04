package utils

import (
	c "bp/constants"
	"os"
	"path/filepath"
)

// checks whether a template exists , returns the status, the error (if occurred) and the directory where it needs to be created (if possible)
func IsTemplateExists(templateName string) (bool, error, string) {

	homeDir, err := os.UserHomeDir()

	//destDir = homeDir + location where we store templates + templateName
	destDir := filepath.Join(homeDir, c.BOILERPLATE_DIR, c.TEMPLATE_DIR, templateName)

	//check if destDir directory exists

	return templateExists, nil, destDir
}
