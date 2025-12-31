/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	c "github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/services"

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

// TODO:: make a seperate service to list out the contents of any provided absolute directory path, then reuse it in this file and the preview template (for directory types)

func ListCmdRunner(cmd *cobra.Command, args []string) {
	templateDir := c.BOILERPLATE_TEMPLATE_DIR // get the template dir from the constants package

	fmt.Println("Available template/s:")

	numEntries, err := services.ListDir(templateDir, true) // read the templateDir
	if err != nil {                                        // catch error while listing directories (if any)
		fmt.Println("[ERROR] Failed to read templates directory")
		fmt.Printf("  Location: %s\n", templateDir)
		fmt.Printf("  Error: %v\n", err)
		fmt.Println("  Make sure the templates directory exists or add a template first")
		os.Exit(1)
	}

	if numEntries == 0 {
		fmt.Println("No templates found")
		fmt.Println("  Use 'bp add <file> --name <template-name>' to add your first template")
		os.Exit(0)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
