package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
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
	templateName := args[0]

	// check wheter the template exists or not
	templateExists, templateDir := utils.IsTemplateExists(templateName)

	if !templateExists {
		fmt.Printf("Template [ %s ] doesn't exists ", templateName)
	}

	templateFileName, err := utils.GetTemplateFileDir(templateName)
	if err != nil {
		fmt.Println(err)
	}

	templateDir = filepath.Join(templateDir, templateFileName)
	fmt.Println(templateDir)
}

func init() {
	rootCmd.AddCommand(clipCmd)
}
