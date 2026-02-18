package services

import (
	"os"

	"github.com/ukhirani/boilerplate/constants"
)

// should return an array/slice of string template names
func GetAllTemplates() ([]string, error) {
	entries, err := os.ReadDir(constants.BOILERPLATE_TEMPLATE_DIR)
	answer := make([]string, len(entries))
	for i, entry := range entries {
		answer[i] = entry.Name()
	}
	return answer, err
}
