/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ukhirani/boilerplate/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ukhirani/boilerplate/types"
)

var (
	templateName string
	newTemplate  types.Config
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a file or directory as a reusable template",
	Long: `Save a file or directory from the current location as a reusable template.

Template names must contain only letters, numbers, and underscores.

Usage:
  bp add <file-or-directory> --name <template-name>

Examples:
  bp add script.sh --name shell-script
  bp add ./components --name react-components`,
	Args: cobra.ExactArgs(1), // This ensures exactly one fileName is passed
	Run:  AddCmdRunner,
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
		fmt.Println("[ERROR] Failed to get current directory")
		fmt.Println("  Error:", err)
		os.Exit(1)
	}

	currDir = filepath.Join(currDir, fileName)

	//check if current directory exists
	currDirexists, err := utils.Exists(currDir)

	// check for any unknown error
	if err != nil {
		fmt.Println("[ERROR] Unexpected error occurred")
		fmt.Printf("  Path: %s\n", currDir)
		fmt.Printf("  Error: %v\n", err)
		os.Exit(1)
	}

	// if current file does not exist
	if !currDirexists {
		fmt.Println("[ERROR] File or directory not found")
		fmt.Printf("  Path: %s\n", currDir)
		fmt.Println("  Make sure the file exists in the current directory")
		os.Exit(1)
	}

	return currDir

}

func destDirValidator(templateName string) string {

	destDirExists, err, destDir := utils.IsTemplateExists(templateName)

	// check for any unknown error
	if err != nil {
		fmt.Println("[ERROR] Unexpected error occurred")
		fmt.Printf("  Error: %v\n", err)
		os.Exit(1)
	}

	// if template exists
	if destDirExists {
		fmt.Println("[ERROR] Template already exists")
		fmt.Printf("  Template: %s\n", templateName)
		fmt.Printf("  Location: %s\n", destDir)
		fmt.Println("  Use a different template name or remove the existing template")
		os.Exit(1)
	}

	//return the path where the template needs to be stored
	return destDir

}

// TODO: as of now this only generates the configs, can be easily refactored to generate template files also
func GenerateTemplate(newTemplate types.Config) {

	//TODO: isn't it better to make a map and iterate over it ?
	viper.SetConfigName(newTemplate.Name)
	viper.Set("Name", newTemplate.Name)
	viper.Set("IsDir", newTemplate.IsDir)

	if err := viper.SafeWriteConfig(); err != nil {
		fmt.Println(err)
	}

}

func AddCmdRunner(cmd *cobra.Command, args []string) {
	//then get the file name entered as the argument
	fileName := args[0]

	// validate that whether the template name is valid or not
	// since we are going to create a directory of that name, it better be a valid directory name
	if !utils.IsValidDirName(templateName) {
		fmt.Println("[ERROR] Invalid template name")
		fmt.Printf("  Template name: %s\n", templateName)
		fmt.Println("  Allowed characters: letters, numbers, and underscores only")
		os.Exit(1)
	}

	// the logic to check whether the current file exists or not
	currDir := currDirValidator(fileName)

	// the logic to check whether the template can be created or not
	destDir := destDirValidator(templateName)

	//check whether the filetype is direcoty or just file
	isDir, err := utils.IsDirectory(fileName)

	if err != nil {
		fmt.Println("[ERROR] Failed to determine file type")
		fmt.Printf("  Error: %v\n", err)
		os.Exit(1)
	}

	//TODO: refactor everything below this in this function

	//if it's a directory
	if isDir {
		if err := utils.CopyDir(currDir, destDir); err != nil {
			fmt.Println("[ERROR] Failed to create template directory")
			fmt.Printf("  Template: %s\n", templateName)
			fmt.Printf("  Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		//if not, then it's regular file
		//since copyDir function copies all the files rather than one
		if err := utils.CopyFile(currDir, destDir, fileName); err != nil {
			fmt.Println("[ERROR] Failed to create template file")
			fmt.Printf("  Template: %s\n", templateName)
			fmt.Printf("  Error: %v\n", err)
			os.Exit(1)
		}
	}

	newTemplate.Name = templateName
	newTemplate.IsDir = isDir

	GenerateTemplate(newTemplate)
	fmt.Printf("[SUCCESS] Template %v created successfully", templateName)
}

func init() {
	rootCmd.AddCommand(addCmd)

	//defining the flags
	addCmd.Flags().StringVarP(&templateName, "name", "n", "", "Template name (letters, numbers, underscores only)")

	//marking flags as required
	addCmd.MarkFlagRequired("name")
}
