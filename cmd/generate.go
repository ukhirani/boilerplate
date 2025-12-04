/*
Copyright © 2025 Umang Hirani
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	generatedFileName string
	generatedFileDir  string
)

func GenerateCmdRunner(cmd *cobra.Command, args []string) {

	fmt.Println("Args Called :", args)
	fmt.Println("generatedFileName Called :", generatedFileName)
	fmt.Println("generatedFileDir Called :", generatedFileDir)

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
	generateCmd.Flags().StringVarP(&generatedFileName, "name", "n", "", "The name of the generated file or directory")
	generateCmd.Flags().StringVarP(&generatedFileDir, "dir", "d", "", "(optional) The directory of the generated file or directory ")

	//making the flags as required
	generateCmd.MarkFlagRequired("name")
}
