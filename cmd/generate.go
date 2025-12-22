/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/services"
	"github.com/ukhirani/boilerplate/types"
	"github.com/ukhirani/boilerplate/utils"

	"github.com/spf13/cobra"
)

var (
	generatedFileName string
	generatedFileDir  string
)

func NameDirValidator(conf *types.Config, cmd *cobra.Command, destDir string, args []string) string {
	// incase template is dir type and --name flage is used
	if conf.IsDir && len(args) == 2 {
		fmt.Println("--name flag or file name is only for file type templates")
		fmt.Printf("[ %s ] is of type Directory", conf.Name)
		os.Exit(1)
	}

	// join the dir flag's value and then current directory
	if cmd.Flags().Changed("dir") {
		destDir = filepath.Join(destDir, generatedFileDir)
	}
	return destDir
}

func GenerateCmdRunner(cmd *cobra.Command, args []string) {
	// we are assured we only have one arguments
	templateName := args[0]
	destDir := constants.CurrDir
	var config types.Config

	// check whether the template exists or not
	templateExists, templateDir := utils.IsTemplateExists(templateName)
	if !templateExists {
		fmt.Printf(" [ERROR] Template [ %s ] not found \n", templateName)
		fmt.Println("  Expected location: \n", templateDir)
		fmt.Println("  Use 'bp list' to see available templates")
		os.Exit(1)
	}

	// read the config for the templateName
	if err := services.ReadConfig(templateName, &config); err != nil {
		fmt.Println("[ERROR] Couldn't find the config for :", templateName)
		os.Exit(1)
	}

	// validate the 2nd arg (as file name) OR --dir flag if passed
	destDir = NameDirValidator(&config, cmd, destDir, args)

	// Execute PreCmd(s) Here
	if len(config.PreCmd) > 0 {
		fmt.Printf("\nRunning Pre-Commands . . .  \n")
		if err := services.ExecCmds(config.PreCmd); err != nil {
			fmt.Println("Error executing Pre-Commands :", err)
			fmt.Println("The template generation will still carry on . . .")
		}
	}

	if config.IsDir {

		// copy template in the current directory
		err := utils.CopyDir(templateDir, destDir)
		if err != nil {
			fmt.Println("[ERROR] Failed to copy template : ", templateName)
			fmt.Printf("  Error: %v\n", err)
			os.Exit(1)
		}

	} else {
		templateDirFile, err := utils.GetTemplateFileDir(templateName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// if 2 args are not provided (i.e. no renaming of the existing file in the template directory)
		// use the file name that is already there in the template file
		if len(args) < 2 {
			generatedFileName = templateDirFile
		} else {
			generatedFileName = args[1]
		}

		err = utils.CopyFile(filepath.Join(templateDir, templateDirFile), destDir, generatedFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Printf("\n[SUCCESS] Template %v generated successfully \n\n", templateName)
	// Execute PostCmd(s) Here
	if len(config.PostCmd) > 0 {
		fmt.Printf("Running Post-Commands . . .  \n")
		if err := services.ExecCmds(config.PostCmd); err != nil {
			fmt.Println("Error executing Post-Commands :", err)
		}
	}
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a file or directory from a template",
	Long: `Copy a template to the current directory, preserving its structure and content.

Usage:
  bp generate <template-name> <generated-file-name> [flags]

Examples:
  bp generate react-component index.jsx
  bp gen my-template`,

	Aliases: []string{"gen", "create"},
	Run:     GenerateCmdRunner,
	Args:    cobra.RangeArgs(1, 2),
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// defining the flags
	generateCmd.Flags().StringVarP(&generatedFileDir, "dir", "d", "", "Target directory for generation (default: current directory)")
}
