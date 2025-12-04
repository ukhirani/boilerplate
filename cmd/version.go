/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/ukhirani/boilerplate/constants"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version of the application",
	Long:  `Display the current version number. Equivalent to using --version or -v flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constants.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
