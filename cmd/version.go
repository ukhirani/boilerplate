/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"github.com/ukhirani/boilerplate/constants"
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints out the current version of the application",
	Long:  `this command tells you the current version of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constants.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
