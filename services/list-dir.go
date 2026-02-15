package services

import (
	"os"

	"github.com/ukhirani/boilerplate/styles"
	"github.com/ukhirani/boilerplate/types"
)

// TODO: can't we do a recursive call to nicely print out a tree

func ListDir(path string, showTypes bool) (int, error) {
	entries, err := os.ReadDir(path)

	for i, entry := range entries {
		if showTypes {
			// TODO: one more problem, we are just listing out the folders to see their names.
			// This is limited in it's way to print out more info regarding any templates.

			// now we have to read the configs of each templates and then find out the fileType
			var conf types.Config
			if err := ReadConfig(entry.Name(), &conf); err != nil {
				styles.PrintError(err.Error())
				os.Exit(1)
			}

			styles.PrintTemplateItem(entry.Name(), conf.IsDir)
		} else {
			isLast := i == len(entries)-1
			styles.PrintTreeItem(entry.Name(), isLast)
		}
	}
	return len(entries), err
}
