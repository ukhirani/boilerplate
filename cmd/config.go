package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configName string
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "used to configure your templates",
	Long:  `used to configure your templates`,
	Run:   ConifgCmdRunner,
}

func ConifgCmdRunner(cobra *cobra.Command, args []string) {
	fmt.Println("Config Command Called")
}

func init() {
	rootCmd.AddCommand(configCmd)

	//declaring the flags
	configCmd.Flags().StringVarP(&configName, "name", "n", "", "used to rename your templates")
}
