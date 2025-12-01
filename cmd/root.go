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

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, versionCmd.Short)
}
