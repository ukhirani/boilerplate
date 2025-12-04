/*
Copyright © 2025 Umang Hirani
*/
package cmd

import (
	"github.com/ukhirani/boilerplate/utils"
	"fmt"
	"os"

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
	templateExists, _, templateDir := utils.IsTemplateExists(templateName)

	if !templateExists {
		fmt.Printf("Template %v Doesn't Exist at %v", templateName, templateDir)
		os.Exit(1)
	}
	fmt.Printf("Template %v Exist at %v", templateName, templateDir)

	//copy template in the current directory
	err := utils.CopyTemplateHere(templateDir)
	if err != nil {
		fmt.Println("Error Copying Template: ", err)
	}
	fmt.Printf("Template %v recreated successfully !", templateName)

}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate a file/directory from a template",
	Long:    "Generate a file/directory from a template",
	Aliases: []string{"gen", "create"},
	Run:     GenerateCmdRunner,
	Args:    cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(generateCmd)

	//defining the flags
	generateCmd.Flags().StringVarP(&generatedFileName, "name", "n", "", "(not applicable for directories) The name of the generated file")
	generateCmd.Flags().StringVarP(&generatedFileDir, "dir", "d", "", "(optional) The directory of the generated file or directory ")

	//making the flags as required
	//generateCmd.MarkFlagRequired("name")
}
