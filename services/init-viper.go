package services

import (
	"os"

	"github.com/spf13/viper"
	c "github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/types"
)

var configPaths = []string{
	c.BOILERPLATE_CONFIG_DIR,
	//TODO: add c.USER_CONFIG_DIR here
}

func InitViper() {

	for _, path := range configPaths {
		//creating the config paths if not existing
		os.MkdirAll(path, os.ModePerm)

		//setting all the possible config paths
		viper.AddConfigPath(path)
	}

	//setting all the defaults for the viper
	for key, value := range types.DefaultConfig {
		viper.SetDefault(key, value)
	}
}
