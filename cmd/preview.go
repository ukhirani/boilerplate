/*
Copyright © 2025 Umang Hirani
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukhirani/boilerplate/services"
	"github.com/ukhirani/boilerplate/types"
	"github.com/ukhirani/boilerplate/utils"
)

var previewConfig bool

// previewCmd represents the preview command
var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "preview the templates with this command",
	Long:  "preview the templates with this command",
	Run:   PreviewCmdRunner,
	Args:  cobra.ExactArgs(1),
}

func PreviewTemplate(templateName string, conf *types.Config) error {
	fmt.Printf("[ %s ] template preview \n", templateName)
	return nil
}

func PreviewConfig(templateName string, conf *types.Config) error {
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
	return nil
}

func PreviewCmdRunner(cmd *cobra.Command, args []string) {
	// get the template name
	templateName := args[0]

	// make a config
	var conf types.Config

	if err := services.ReadConfig(templateName, &conf); err != nil {
		fmt.Println("Error Fetching the config for the template : ", templateName)
		fmt.Println("[Error] :", err)
	}

	// check whether the template exists
	if exists, _ := utils.IsTemplateExists(templateName); !exists {
		fmt.Printf("Template [%s] not found.", templateName)
	}

	// if we want to view the config
	if previewConfig {
		err := PreviewConfig(templateName, &conf)
		if err != nil {
			fmt.Println(err)
		}
		return
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
