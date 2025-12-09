package services

import (
	"os"

	"github.com/spf13/viper"
	c "github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/types"
)

var configPaths = []string{
	c.BOILERPLATE_CONFIG_DIR,
	c.BOILERPLATE_TEMPLATE_DIR,
	// TODO: add c.UserConfigDir here
}

func InitViper() {
	viper.SetConfigType("toml")

	for _, path := range configPaths {
		// creating the config path if not existing
		os.MkdirAll(path, os.ModePerm)
		viper.AddConfigPath(path)
	}

	// setting all the defaults for the viper
	for key, value := range types.DefaultConfig {
		viper.SetDefault(key, value)
	}
}
