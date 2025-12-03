/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bp/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	templateName string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a template to your arsenal",
	Long:  `Add a template to your arsenal`,
	Run:   AddCmdRunner,
}

func AddCmdRunner(cmd *cobra.Command, args []string) {

	if len(templateName) == 0 {
		fmt.Println("Template name not mentioned. Use the flag --name or -n to pass the template name")
		os.Exit(1)
	}

	if len(args) != 1 {
		fmt.Println(`invalid amount of file/folder(s) passed as a template : `, len(args), "passed !")
		os.Exit(1)
	}

	//get the current directory where you are
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory")
		os.Exit(1)
	}

	currDir = filepath.Join(currDir, args[0])

	// check if current directory exists
	currDirexists, err := utils.Exists(currDir)

	// check for any unknown error
	if err != nil {
		fmt.Println("Unexpected error occurred : ", err)
		os.Exit(1)
	}

	// if current file does not exist
	if !currDirexists {
		fmt.Println("File Not Found : ", currDir)
		os.Exit(1)
	}

	fmt.Println("File Exists : ", currDir, currDirexists)

}

func init() {
	rootCmd.AddCommand(addCmd)

	//defining the flags
	addCmd.Flags().StringVarP(&templateName, "name", "n", "", "Name of the template that you wish to add")
}
