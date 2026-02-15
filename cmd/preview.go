/*
Copyright © 2025 Umang Hirani
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/services"
	"github.com/ukhirani/boilerplate/styles"
	"github.com/ukhirani/boilerplate/types"
	"github.com/ukhirani/boilerplate/utils"
)

var previewConfig bool

// previewCmd represents the preview command
var previewCmd = &cobra.Command{
	Use:     "preview",
	Short:   "preview the templates with this command",
	Long:    "preview the templates with this command",
	Run:     PreviewCmdRunner,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"view"},
}

func PreviewTemplate(templateName string, conf *types.Config) error {
	styles.PrintHeader("Template: " + styles.Highlight(templateName))

	templateDir := filepath.Join(constants.BOILERPLATE_TEMPLATE_DIR, templateName)

	// then print out the directory as it is
	if conf.IsDir {
		if _, err := services.ListDir(templateDir, false); err != nil {
			styles.PrintErrorWithDetails(
				"Error listing template content",
				err.Error(),
			)
		}
	} else {
		templateFileName, err := utils.GetTemplateFileDir(templateName)
		if err != nil {
			styles.PrintError(err.Error())
			os.Exit(1)
		}
		data, err := os.ReadFile(filepath.Join(templateDir, templateFileName))
		if err != nil {
			styles.PrintError(err.Error())
			os.Exit(1)
		}

		fmt.Println(string(data))
		return nil
	}
	return nil
}

func PreviewConfig(templateName string, conf *types.Config) {
	// determining the file type
	templateType := "File"
	if conf.IsDir {
		templateType = "Directory"
	}
	styles.PrintHeader("Config: " + styles.Highlight(templateName))

	styles.PrintKeyValue("Template Name", conf.Name)
	styles.PrintKeyValue("Template Type", templateType)

	if len(conf.PreCmd) > 0 {
		styles.PrintNewLine()
		styles.PrintSubHeader("Pre-commands")

		// printing out all the preCmds
		for i, v := range conf.PreCmd {
			styles.PrintStep(i+1, len(conf.PreCmd), v)
		}
	}

	if len(conf.PostCmd) > 0 {
		styles.PrintNewLine()
		styles.PrintSubHeader("Post-commands")

		// printing out all the postCmds
		for i, v := range conf.PostCmd {
			styles.PrintStep(i+1, len(conf.PostCmd), v)
		}
	}
}

func PreviewCmdRunner(cmd *cobra.Command, args []string) {
	// get the template name
	templateName := args[0]

	// make a config
	var conf types.Config

	if err := services.ReadConfig(templateName, &conf); err != nil {
		styles.PrintErrorWithDetails(
			"Error fetching the config for template: "+styles.Highlight(templateName),
			err.Error(),
		)
		os.Exit(1)
	}

	// check whether the template exists
	if exists, _ := utils.IsTemplateExists(templateName); !exists {
		styles.PrintError("Template " + styles.Highlight(templateName) + " not found")
		os.Exit(1)
	}

	// if we want to view the config
	if previewConfig {
		PreviewConfig(templateName, &conf)
		os.Exit(0)
	}

	// else show the exact template if
	err := PreviewTemplate(templateName, &conf)
	if err != nil {
		styles.PrintError(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(previewCmd)
	previewCmd.Flags().BoolVarP(&previewConfig, "config", "c", false, "preview the configuration of the template")
}
