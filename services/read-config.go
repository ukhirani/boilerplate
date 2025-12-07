package services

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/ukhirani/boilerplate/types"
)

func ReadConfig(templateName string, conf *types.Config) error {
	viper.SetConfigName(templateName)
	if err := viper.ReadInConfig(); err != nil {
		// Config file not found
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, using defaults or environment variables.")
		} else {
			// Config file found but another error occurred
			fmt.Printf("Fatal error reading config file: %s \n", err)
		}
		return err
	}

	// set the config to the things that were found
	conf.Name = viper.GetString("Name")
	conf.IsDir = viper.GetBool("IsDir")
	conf.PreCmd = viper.GetStringSlice("PreCmd")
	conf.PostCmd = viper.GetStringSlice("PostCmd")

	return nil
}
