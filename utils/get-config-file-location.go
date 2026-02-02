package utils

import (
	"path/filepath"

	"github.com/ukhirani/boilerplate/constants"
)

// WARN : this function doesn't return any errors and assumes that the template alreay exists
func GetConfigFileLocation(templateName string) string {
	return filepath.Join(constants.BOILERPLATE_CONFIG_DIR, templateName+".toml")
}
