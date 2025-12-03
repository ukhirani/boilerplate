/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	version bool
)

var rootCmd = &cobra.Command{
	Use:   "bp",
	Short: "A cli tool to ironically skip the boilerplate",
	Long:  "A cli tool to ironically skip the boilerplate",

	Run: RootCmdRunner,
}

func RootCmdRunner(cmd *cobra.Command, args []string) {

	// check if the version command is called
	if version {
		versionCmd.Run(cmd, args)
		os.Exit(0)
	}

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// --version and -v both wil call the versionCmd
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, versionCmd.Short)
}
