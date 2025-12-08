package services

import (
	"fmt"
	"os"
)

func ListDir(path string) (int, error) {
	entries, err := os.ReadDir(path)

	for _, entry := range entries {
		fmt.Printf("  • %s\n", entry.Name())
	}

	return len(entries), err
}
