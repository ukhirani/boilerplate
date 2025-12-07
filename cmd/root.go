/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/ukhirani/boilerplate/services"
)

var version bool

var rootCmd = &cobra.Command{
	Use:   "bp",
	Short: "A CLI tool to ironically skip writing boilerplate code",
	Long: `Manage reusable file and directory templates. Save common code structures 
as templates and generate them instantly.`,

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
	services.InitViper()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// --version and -v both wil call the versionCmd
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, versionCmd.Short)
}
