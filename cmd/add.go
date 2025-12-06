/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ukhirani/boilerplate/constants"
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
func currDirValidator(fileName string) {

	//get the current directory where you are + fileName
	currDir := filepath.Join(constants.CURR_DIR, fileName)

	//check if current directory exists
	currDirexists := utils.Exists(currDir)

	// check for any unknown error

	// if current file does not exist
	if !currDirexists {
		fmt.Println("[ERROR] File or directory not found")
		fmt.Printf("  Path: %s\n", currDir)
		fmt.Println("  Make sure the file exists in the current directory")
		os.Exit(1)
	}

}

func destDirValidator(templateName string) {

	destDirExists, destDir := utils.IsTemplateExists(templateName)

	// if template exists
	if destDirExists {
		fmt.Printf("[ERROR] Template [%v] already exists", templateName)
		fmt.Printf("  Location: %s\n", destDir)
		fmt.Println("  Use a different template name or remove the existing template")
		os.Exit(1)
	}
}

func GenerateTemplate(fileName, templateName string, isDir bool) {

	currDir := filepath.Join(constants.CURR_DIR, fileName)
	destDir := filepath.Join(constants.BOILERPLATE_TEMPLATE_DIR, templateName)

	//if it's a directory
	if isDir {
		if err := utils.CopyDir(currDir, destDir); err != nil {
			fmt.Printf("[ERROR] Failed to create template [ %s ]", templateName)
			fmt.Printf("  Error: %v\n", err)
			if errors.Is(err, fs.ErrExist) {
				fmt.Println("Some files might be copied upto this point from the template")
				fmt.Println("Please delete all the duplicate files and try again")
			}
			os.Exit(1)
		}
	} else {
		//if not, then it's regular file
		//since copyDir function copies all the files rather than one
		if err := utils.CopyFile(currDir, destDir, fileName); err != nil {
			fmt.Printf("[ERROR] Failed to create template [ %s ]", templateName)
			fmt.Printf("  Error: %v\n", err)
			os.Exit(1)
		}
	}

	// Setting the Viper Configs
	viper.SetConfigName(templateName)
	viper.Set("Name", templateName)
	viper.Set("IsDir", isDir)

	if err := viper.SafeWriteConfig(); err != nil {
		fmt.Println(err)
		// TODO: don't we have to fallback when we can't generate a config ?
	}

}

func AddCmdRunner(cmd *cobra.Command, args []string) {
	//then get the file name entered as the argument
	fileName := args[0]

	// validate that whether the template name is valid or not
	// since we are going to create a directory of that name, it better be a valid directory name
	if !utils.IsValidDirName(templateName) {
		fmt.Println("[ERROR] Invalid template name : ", templateName)
		fmt.Println(" - Allowed characters: letters, numbers, and underscores only")
		os.Exit(1)
	}

	// the logic to check whether the current file exists or not (exits the program if not satisfied)
	currDirValidator(fileName)

	// the logic to check whether the template can be created or not (exits the program if not satisfied)
	destDirValidator(templateName)

	//check whether the filetype is directory or just file
	isDir, err := utils.IsDirectory(fileName)

	if err != nil {
		fmt.Println("[ERROR] Failed to determine file type")
		fmt.Printf("  Error: %v\n", err)
		os.Exit(1)
	}

	GenerateTemplate(fileName, templateName, isDir)

	fmt.Printf("[SUCCESS] Template [ %v ] created successfully", templateName)
}

func init() {
	rootCmd.AddCommand(addCmd)

	//defining the flags
	addCmd.Flags().StringVarP(&templateName, "name", "n", "", "Template name (letters, numbers, underscores only)")

	//marking flags as required
	addCmd.MarkFlagRequired("name")
}
