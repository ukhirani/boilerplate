/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	c "github.com/ukhirani/boilerplate/constants"

	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all available templates",
	Long:    `List all available templates stored in $HOME/boilerplate/templates/`,
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
		fmt.Println("[ERROR] Failed to read templates directory")
		fmt.Printf("  Location: %s\n", templateDir)
		fmt.Printf("  Error: %v\n", err)
		fmt.Println("  Make sure the templates directory exists or add a template first")
		os.Exit(1)
	}

	if len(entries) == 0 {
		fmt.Println("No templates found")
		fmt.Println("  Use 'bp add <file> --name <template-name>' to add your first template")
		return
	}

	fmt.Println("Available templates:")
	for _, entry := range entries {
		fmt.Printf("  • %s\n", entry.Name())
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
