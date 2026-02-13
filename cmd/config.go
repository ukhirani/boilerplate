package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/services"
	"github.com/ukhirani/boilerplate/utils"
)

var customEditor string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "used to configure your templates",
	Long:  `used to configure your templates`,
	RunE:  ConifgCmdRunner,
	Args:  cobra.ExactArgs(1),
}

func ConifgCmdRunner(cmd *cobra.Command, args []string) error {
	// get the first arg which will be the template name
	templateName := args[0]

	// use default editor, if no custom editor is passed via --editor flag
	editorOfChoice := constants.DEFAULT_EDITOR
	if cmd.Flags().Changed("editor") {
		editorOfChoice = customEditor
	}

	// can't let the -e/--editor flag be empty (because you can't open any editor with that cmd)
	if editorOfChoice == "" {
		return fmt.Errorf("--editor cannot be empty")
	}

	// check if template exists
	if ok, _ := utils.IsTemplateExists(templateName); !ok {
		fmt.Printf("Template %v doesn't exist", templateName)
		os.Exit(1)
	}

	templateConfigLocation := utils.GetConfigFileLocation(templateName)

	cmdToOpenEditor := strings.Join([]string{editorOfChoice, templateConfigLocation}, " ")
	if err := services.ExecCmds([]string{cmdToOpenEditor}); err != nil {
		fmt.Println("Error opening config", err)
		os.Exit(1)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(configCmd)

	// declaring the flags
	configCmd.Flags().StringVarP(&customEditor, "editor", "e", "", "used to open your configs in your editor")
}
