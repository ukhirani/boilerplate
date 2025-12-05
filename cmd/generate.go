/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/utils"

	"github.com/spf13/cobra"
)

var (
	generatedFileName string
	generatedFileDir  string
)

func GenerateCmdRunner(cmd *cobra.Command, args []string) {

	// we are assured we only have one arguments
	templateName := args[0]

	//check whether the template exists or not
	templateExists, templateDir := utils.IsTemplateExists(templateName)

	if !templateExists {
		fmt.Println("[ERROR] Template not found")
		fmt.Printf("  Template: %s\n", templateName)
		fmt.Printf("  Expected location: %s\n", templateDir)
		fmt.Println("  Use 'bp list' to see available templates")
		os.Exit(1)
	}

	//copy template in the current directory
	err := utils.CopyDir(templateDir, constants.CURR_DIR)
	if err != nil {
		fmt.Println("[ERROR] Failed to copy template : ", templateName)
		fmt.Printf("  Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[SUCCESS] Template %v generated successfully", templateName)

}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a file or directory from a template",
	Long: `Copy a template to the current directory, preserving its structure and content.

Usage:
  bp generate <template-name> [flags]

Examples:
  bp generate react-component
  bp gen my-template`,
	Aliases: []string{"gen", "create"},
	Run:     GenerateCmdRunner,
	Args:    cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(generateCmd)

	//defining the flags
	//TODO: use the --dir and the --name flags
	generateCmd.Flags().StringVarP(&generatedFileName, "name", "n", "", "Custom name for the generated file (files only, not directories)")
	generateCmd.Flags().StringVarP(&generatedFileDir, "dir", "d", "", "Target directory for generation (default: current directory)")

	//making the flags as required
	//generateCmd.MarkFlagRequired("name")
}
