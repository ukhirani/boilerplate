package utils

import (
	c "github.com/ukhirani/boilerplate/constants"
	"path/filepath"
)

//TODO: make a full proof mechanism, instead of just seeing whether a directory exists or not

// checks whether a template exists , returns the status, the error (if occurred) and the directory where it needs to be created (if possible)
func IsTemplateExists(templateName string) (bool, string) {

	//destDir = homeDir + location where we store templates + templateName
	destDir := filepath.Join(c.BOILERPLATE_TEMPLATE_DIR, templateName)

	//check if destDir directory exists
	templateExists := Exists(destDir)

	return templateExists, destDir
}
