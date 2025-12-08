/*
Copyright © 2025 Umang Hirani
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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

func PreviewTemplate(templateName string) error {
	return nil
}

func PreviewConfig(templateName string) error {
	return nil
}

func PreviewCmdRunner(cmd *cobra.Command, args []string) {
	templateName := args[0]

	// if we want to view the config
	if previewConfig {
		err := PreviewConfig(templateName)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// else show the exact template if
	err := PreviewTemplate(templateName)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(previewCmd)
	previewCmd.Flags().BoolVarP(&previewConfig, "config", "c", false, "preview the configuration of the template")
}
