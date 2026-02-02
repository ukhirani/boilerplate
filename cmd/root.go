package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
	"github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/services"
)

var rootCmd = &cobra.Command{
	Use:   "bp",
	Short: "A CLI tool to ironically skip writing boilerplate code",
	Long: `Manage reusable file and directory templates.
Save common code structures as templates and generate them instantly.`,
}

func Execute() {
	services.InitViper()

	// THIS is the only version wiring you need
	rootCmd.SetVersionTemplate(constants.VERSION)
	rootCmd.Version = constants.VERSION

	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}
