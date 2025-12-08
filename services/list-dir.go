package services

import (
	"fmt"
	"os"
)

// TODO: can't we do a recursive call to nicely print out a tree

func ListDir(path string) (int, error) {
	entries, err := os.ReadDir(path)

	for _, entry := range entries {
		fmt.Printf("  • %s\n", entry.Name())
	}

	return len(entries), err
}
