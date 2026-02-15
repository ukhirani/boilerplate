/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/services"
	"github.com/ukhirani/boilerplate/styles"
	"github.com/ukhirani/boilerplate/types"
	"github.com/ukhirani/boilerplate/utils"

	"github.com/spf13/cobra"
)

var (
	generatedFileName string
	generatedFileDir  string
)

func NameDirValidator(conf *types.Config, cmd *cobra.Command, destDir string, args []string) string {
	// incase template is dir type and file name is used
	if conf.IsDir && len(args) == 2 {
		styles.PrintErrorWithDetails(
			"File name argument is only for file-type templates",
			styles.Highlight(conf.Name)+" is of type Directory",
		)
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
		styles.PrintErrorWithDetails(
			"Template "+styles.Highlight(templateName)+" not found",
			"Expected: "+styles.Path(templateDir),
			"Run "+styles.Code("bp list")+" to see available templates",
		)
		os.Exit(1)
	}

	// read the config for the templateName
	if err := services.ReadConfig(templateName, &config); err != nil {
		styles.PrintError("Couldn't find the config for " + styles.Highlight(templateName))
		os.Exit(1)
	}

	// validate the 2nd arg (as file name) OR --dir flag if passed
	destDir = NameDirValidator(&config, cmd, destDir, args)

	// Execute PreCmd(s) Here
	if len(config.PreCmd) > 0 {
		styles.PrintNewLine()
		styles.PrintHeader("Running Pre-Commands")
		if err := services.ExecCmds(config.PreCmd); err != nil {
			styles.PrintWarning("Pre-command failed: " + err.Error())
			styles.PrintMuted("Template generation will continue...")
		}
	}

	if config.IsDir {

		// copy template in the current directory
		err := utils.CopyDir(templateDir, destDir)
		if err != nil {
			styles.PrintErrorWithDetails(
				"Failed to copy template "+styles.Highlight(templateName),
				err.Error(),
			)
			os.Exit(1)
		}

	} else {
		templateDirFile, err := utils.GetTemplateFileDir(templateName)
		if err != nil {
			styles.PrintError(err.Error())
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
			styles.PrintError(err.Error())
			os.Exit(1)
		}
	}

	styles.PrintNewLine()
	styles.PrintSuccess("Template " + styles.Highlight(templateName) + " generated successfully")

	// Execute PostCmd(s) Here
	if len(config.PostCmd) > 0 {
		styles.PrintNewLine()
		styles.PrintHeader("Running Post-Commands")
		if err := services.ExecCmds(config.PostCmd); err != nil {
			styles.PrintWarning("Post-command failed: " + err.Error())
		}
	}
	styles.PrintNewLine()
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a file or directory from a template",
	Long: `Copy a template to the current directory, preserving its structure and content.

	aliases : generate , create

	`,

	Aliases: []string{"generate", "create"},
	Run:     GenerateCmdRunner,
	Args:    cobra.RangeArgs(1, 2),
	Example: `bp gen cpp-template main.cpp
						bp gen react-template --dir project`,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// defining the flags
	generateCmd.Flags().StringVarP(&generatedFileDir, "dir", "d", "", "Target directory for generation (default: current directory)")
}
