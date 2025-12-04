/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"fmt"
	c "github.com/ukhirani/boilerplate/constants"
	"os"

	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list out all the templates",
	Long:    `list out all the templates`,
	Aliases: []string{"ls"},
	Run:     ListCmdRunner,
}

func ListCmdRunner(cmd *cobra.Command, args []string) {

	homeDir, _ := os.UserHomeDir()

	//destDir = homeDir + location where we store templates
	templateDir := filepath.Join(homeDir, c.BOILERPLATE_DIR, c.TEMPLATE_DIR)

	//read the templateDir
	entries, err := os.ReadDir(templateDir)
	if err != nil {
		fmt.Printf("Error reading directory '%s': %v\n", templateDir, err)
		os.Exit(1)
	}

	for _, entry := range entries {
		fmt.Printf(" -> %s\n", entry.Name())
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
