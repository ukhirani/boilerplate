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
	Use:   "clip",
	Short: "add a template to your system clipboard",
	Long:  "add a template to your system clipboard",
	Run:   ClipCmdRunner,
	Args:  cobra.ExactArgs(1),
}

func ClipCmdRunner(cmd *cobra.Command, args []string) {
	// guaranteed that we have exactly one argument
	templateName := args[0]

	// check wheter the template exists or not
	templateExists, templateDir := utils.IsTemplateExists(templateName)

	if !templateExists {
		fmt.Printf("Template [ %s ] doesn't exist. \n ", templateName)
		os.Exit(1)
	}

	var conf types.Config
	services.ReadConfig(templateName, &conf)

	if conf.IsDir {
		fmt.Printf("Template [ %s ] is of type directory and can't be copied \n ", templateName)
		os.Exit(1)
	}

	templateFileName, err := utils.GetTemplateFileDir(templateName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	templateDir = filepath.Join(templateDir, templateFileName)
	fmt.Println(templateDir)

	data, err := os.ReadFile(templateDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = clipboard.WriteAll(string(data)); err != nil {
		fmt.Println("error writing to clipboard :", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(clipCmd)
}
