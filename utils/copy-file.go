package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(srcDir, destDir, fileName string) error {
	// Open the source file for reading.
	sourceFile, err := os.Open(srcDir)
	if err != nil {
		return fmt.Errorf("failed to open source file '%s': %w", srcDir, err)
	}
	defer sourceFile.Close()

	// create template directory
	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		fmt.Println("[ERROR] Failed to create directory")
		fmt.Printf("  Path: %s\n", destDir)
		fmt.Printf("  Error: %v\n", err)
		os.Exit(1)
	}

	// make destDir the actual fileName where it needs to be copied
	destDir = filepath.Join(destDir, fileName)

	if isDestExists := Exists(destDir); isDestExists {
		return fmt.Errorf("File already exists, choose a different destination file ")
	}

	// Create the destination file for writing.
	destinationFile, err := os.Create(destDir)
	if err != nil {
		return fmt.Errorf("failed to create destination file '%s': %w", destDir, err)
	}
	defer destinationFile.Close() // Ensure the destination file is closed.

	// Copy the contents from the source file to the destination file.
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file contents from '%s' to '%s': %w", srcDir, destDir, err)
	}

	return nil // Return nil if the copy was successful.
}
