/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bp/constants"
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
	Args:  cobra.ExactArgs(1), // This ensures exactly one fileName is passed
	Run:   AddCmdRunner,
}

/*
arguments:
- fileName (string) : the file's name that needs to be copied as a template

This function does the following things :
- Fetches the current directory
- Appends the fileName to it
- Checks for the presence of this file

if yes -> return the concatenated fileName
else ->   exit
*/
func currDirValidator(fileName string) string {

	//get the current directory where you are
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory")
		os.Exit(1)
	}

	currDir = filepath.Join(currDir, fileName)

	//check if current directory exists
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

	return currDir

}

func destDirValidator(templateName string) string {

	homeDir, _ := os.UserHomeDir()

	//destDir = homeDir + location where we store templates + templateName
	destDir := filepath.Join(homeDir, constants.BOILERPLATE_DIR, constants.TEMPLATE_DIR, templateName)

	//check if destDir directory exists
	destDirexists, err := utils.Exists(destDir)

	// check for any unknown error
	if err != nil {
		fmt.Println("Unexpected error occurred : ", err)
		os.Exit(1)
	}

	// if template exists
	if destDirexists {
		fmt.Println("Template Already Exists : ", destDir)
		os.Exit(1)
	}

	//return the path where the template needs to be stored
	return destDir

}

func AddCmdRunner(cmd *cobra.Command, args []string) {
	//then get the file name entered as the argument
	fileName := args[0]

	// the logic to check whether the current file exists or not

	currDir := currDirValidator(fileName)
	fmt.Println("File Found Successfully at : ", currDir)

	destDir := destDirValidator(templateName)
	fmt.Println("Template Can be Created at : ", destDir)

}

func init() {
	rootCmd.AddCommand(addCmd)

	//defining the flags
	addCmd.Flags().StringVarP(&templateName, "name", "n", "", "Name of the template that you wish to add")

	//marking flags as required
	addCmd.MarkFlagRequired("name")
}
