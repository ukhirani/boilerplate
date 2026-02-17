/*
Copyright © 2025 Umang Hirani
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/ukhirani/boilerplate/services"
	"github.com/ukhirani/boilerplate/styles"
	"github.com/ukhirani/boilerplate/types"
	"github.com/ukhirani/boilerplate/utils"
)

// previewCmd represents the preview command
var workCmd = &cobra.Command{
	Use:   "work",
	Short: "run the pre and post commands of any template without generating the template",
	Long:  "run the pre and post commands of any template without generating the template",
	Run:   WorkCmdRunner,
	Args:  cobra.ExactArgs(1),
}

var (
	runPreCmdsOnly  bool
	runPostCmdsOnly bool
)

func WorkCmdRunner(cmd *cobra.Command, args []string) {
	if runPreCmdsOnly && runPostCmdsOnly {
		styles.PrintError("cannot have both " + styles.Highlight("--pre") + " " + styles.Highlight("--post") + " " + "flags together")
		os.Exit(1)
	}

	// get the template name
	templateName := args[0]

	// check whether the template exists
	if exists, _ := utils.IsTemplateExists(templateName); !exists {
		styles.PrintError("Template " + styles.Highlight(templateName) + " not found")
		os.Exit(1)
	}

	// make a config and read from config files
	var conf types.Config

	if err := services.ReadConfig(templateName, &conf); err != nil {
		styles.PrintErrorWithDetails(
			"Error fetching the config for template: "+styles.Highlight(templateName),
			err.Error(),
		)
		os.Exit(1)
	}

	// Execute PreCmd(s) Here
	if !runPostCmdsOnly && len(conf.PreCmd) > 0 {
		styles.PrintNewLine()
		styles.PrintHeader("Running Pre-Commands")
		if err := services.ExecCmds(conf.PreCmd); err != nil {
			styles.PrintWarning("Pre-command failed: " + err.Error())
		}
	}
	// Execute PostCmd(s) Here
	if !runPreCmdsOnly && len(conf.PostCmd) > 0 {
		styles.PrintNewLine()
		styles.PrintHeader("Running Post-Commands")
		if err := services.ExecCmds(conf.PostCmd); err != nil {
			styles.PrintWarning("Post-command failed: " + err.Error())
		}
	}
	styles.PrintNewLine()
}

func init() {
	rootCmd.AddCommand(workCmd)
	workCmd.Flags().BoolVarP(&runPreCmdsOnly, "pre", "p", false, "run the pre-commands of the workflow only")
	workCmd.Flags().BoolVarP(&runPostCmdsOnly, "post", "P", false, "run the post-commands of the workflow only")
}
