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
func AddCmdRunner(cmd *cobra.Command, args []string) {

	if len(args) != 1 {
		fmt.Println(`invalid amount of file/folder(s) passed as a template : `, len(args), "passed !")
		os.Exit(1)
	}

	//then get the file name entered as the argument
	fileName := args[0]

	// the logic to check whether the current file exists or not

	currDir := currDirValidator(fileName)
	fmt.Println("File Exists : ", currDir)
	// destDir := destDirValidator(templateName)

}

func init() {
	rootCmd.AddCommand(addCmd)

	//defining the flags
	addCmd.Flags().StringVarP(&templateName, "name", "n", "", "Name of the template that you wish to add")

	//marking flags as required
	addCmd.MarkFlagRequired("name")
}
