package cmd

import (
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/ukhirani/boilerplate/services"
	"github.com/ukhirani/boilerplate/styles"
	"github.com/ukhirani/boilerplate/types"
	"github.com/ukhirani/boilerplate/utils"
)

// clipCmd represents the clip command
var clipCmd = &cobra.Command{
	Use:     "clip",
	Short:   "add a template to your system clipboard",
	Long:    "add a template to your system clipboard",
	Run:     ClipCmdRunner,
	Args:    cobra.ExactArgs(1),
	Example: "bp clip cpp-template",
}

func ClipCmdRunner(cmd *cobra.Command, args []string) {
	// guaranteed that we have exactly one argument
	templateName := args[0]

	// check wheter the template exists or not
	templateExists, templateDir := utils.IsTemplateExists(templateName)

	// quit if template doesn't exist
	if !templateExists {
		styles.PrintError("Template " + styles.Highlight(templateName) + " doesn't exist")
		os.Exit(1)
	}

	// read the config
	var conf types.Config
	if err := services.ReadConfig(templateName, &conf); err != nil {
		styles.PrintError("Error reading config: " + err.Error())
		os.Exit(1)
	}

	// quit if template is type dir
	if conf.IsDir {
		styles.PrintErrorWithDetails(
			"Template "+styles.Highlight(templateName)+" is a directory",
			"Directory templates cannot be copied to clipboard",
		)
		os.Exit(1)
	}

	// get the exact template file name, then to read the content inside it
	templateFileName, err := utils.GetTemplateFileDir(templateName)
	if err != nil {
		styles.PrintError(err.Error())
		os.Exit(1)
	}

	// join the template file path and the file path
	templateDir = filepath.Join(templateDir, templateFileName)

	// read the file content
	data, err := os.ReadFile(templateDir)
	if err != nil {
		styles.PrintError(err.Error())
		os.Exit(1)
	}

	// write to clipboard and exit if errors
	if err = clipboard.WriteAll(string(data)); err != nil {
		styles.PrintError("Error writing to clipboard: " + err.Error())
		os.Exit(1)
	}

	styles.PrintSuccess("Template " + styles.Highlight(templateName) + " copied to clipboard")
}

func init() {
	rootCmd.AddCommand(clipCmd)
}
