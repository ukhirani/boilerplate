package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ukhirani/boilerplate/constants"
)

// GetTemplateFileDir : a function for file-type templates which simply returns the name of the file of a file-type template
func GetTemplateFileDir(templateName string) (string, error) {
	// get the files inside template_dir/template_name
	entries, _ := os.ReadDir(filepath.Join(constants.BOILERPLATE_TEMPLATE_DIR, templateName))

	// return error if no / more than one file found in the template's directory
	if len(entries) != 1 {
		return "", fmt.Errorf("a template of file type should have exactly one file inside it's template directory, found %d ", len(entries))
	}

	return entries[0].Name(), nil
}
