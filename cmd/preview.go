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
	fmt.Printf("[ %s ] template preview \n", templateName)

	templateDir := filepath.Join(constants.BOILERPLATE_TEMPLATE_DIR, templateName)

	// then print out the directory as it is
	if conf.IsDir {
		if _, err := services.ListDir(templateDir, false); err != nil {
			fmt.Println("Error listing template content of ", templateName)
			fmt.Println("Error : ", err)
		}
	} else {
		templateFileName, err := utils.GetTemplateFileDir(templateName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		data, err := os.ReadFile(filepath.Join(templateDir, templateFileName))
		if err != nil {
			fmt.Println(err)
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
	fmt.Printf("[ %s ] config preview \n", templateName)

	fmt.Println("Template Name :", conf.Name)
	fmt.Println("Template type :", templateType)

	if len(conf.PreCmd) > 0 {
		fmt.Printf("\nHere are the pre-commands \n")

		// printing out all the preCmds
		for i, v := range conf.PreCmd {
			fmt.Printf(" • PRE-CMD %d [ %s ] \n", i+1, v)
		}
	}

	if len(conf.PostCmd) > 0 {
		fmt.Printf("\nHere are the post-commands \n")

		// printing out all the postCmds
		for i, v := range conf.PostCmd {
			fmt.Printf(" • POST-CMD %d [ %s ] \n", i+1, v)
		}
	}
}

func PreviewCmdRunner(cmd *cobra.Command, args []string) {
	// get the template name
	templateName := args[0]

	// make a config
	var conf types.Config

	if err := services.ReadConfig(templateName, &conf); err != nil {
		fmt.Println("Error Fetching the config for the template : ", templateName)
		fmt.Println("[Error] :", err)
		os.Exit(1)
	}

	// check whether the template exists
	if exists, _ := utils.IsTemplateExists(templateName); !exists {
		fmt.Printf("Template [%s] not found.", templateName)
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
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(previewCmd)
	previewCmd.Flags().BoolVarP(&previewConfig, "config", "c", false, "preview the configuration of the template")
}
