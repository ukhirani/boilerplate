package services

import (
	"fmt"
	"os"

	"github.com/ukhirani/boilerplate/types"
)

// TODO: can't we do a recursive call to nicely print out a tree

func ListDir(path string, showTypes bool) (int, error) {
	entries, err := os.ReadDir(path)

	for _, entry := range entries {
		if showTypes {
			// TODO: one more problem, we are just listing out the folders to see their names.
			// This is limited in it's way to print out more info regarding any templates.

			fileType := "FILE"

			// now we have to read the configs of each templates and then find out the fileType
			var conf types.Config
			if err := ReadConfig(entry.Name(), &conf); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if conf.IsDir {
				fileType = "DIR"
			}

			fmt.Printf("[ %s ] %s\n", fileType, entry.Name())
		} else {
			fmt.Printf("• %s\n", entry.Name())
		}
	}
	return len(entries), err
}
