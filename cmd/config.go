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

// var openEditor bool

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "used to configure your templates",
	Long:  `used to configure your templates`,
	Run:   ConifgCmdRunner,
	Args:  cobra.ExactArgs(1),
}

func ConifgCmdRunner(cobra *cobra.Command, args []string) {
	// get the first arg which will be the template name
	templateName := args[0]

	// check if template exists
	if ok, _ := utils.IsTemplateExists(templateName); !ok {
		fmt.Printf("Template %v doesn't exist", templateName)
		os.Exit(1)
	}

	templateConfigLocation := utils.GetConfigFileLocation(templateName)

	if err := services.ExecCmds([]string{strings.Join([]string{constants.DEFAULT_EDITOR, templateConfigLocation}, " ")}); err != nil {
		fmt.Println("Error opening config", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)

	// declaring the flags
	// configCmd.Flags().BoolVarP(&openEditor, "edit", "e", false, "used to open your configs in your editor")
}
