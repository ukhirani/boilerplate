package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/ukhirani/boilerplate/services"
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
		fmt.Printf("Template [ %s ] doesn't exist. \n ", templateName)
		os.Exit(1)
	}

	// read the config
	var conf types.Config
	if err := services.ReadConfig(templateName, &conf); err != nil {
		fmt.Printf("Error reading config : %s .", err)
		os.Exit(1)
	}

	// quit if template is type dir
	if conf.IsDir {
		fmt.Printf("Template [ %s ] is of type directory and can't be copied  \n ", templateName)
		os.Exit(1)
	}

	// get the exact template file name, then to read the content inside it
	templateFileName, err := utils.GetTemplateFileDir(templateName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// join the template file path and the file path
	templateDir = filepath.Join(templateDir, templateFileName)

	// read the file content
	data, err := os.ReadFile(templateDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// write to clipboard and exit if errors
	if err = clipboard.WriteAll(string(data)); err != nil {
		fmt.Println("error writing to clipboard :", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(clipCmd)
}
