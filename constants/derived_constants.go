package constants

import (
	"os"
	"path/filepath"
)

var (
	HOME_DIR, _      = os.UserHomeDir()
	UserConfigDir, _ = os.UserConfigDir()
	CURR_DIR, _      = os.Getwd()
)

// Boilerplate Constants
var (
	BOILERPLATE_CONFIG_DIR   = filepath.Join(HOME_DIR, BOILERPLATE_DIR, CONFIG_DIR)
	BOILERPLATE_TEMPLATE_DIR = filepath.Join(HOME_DIR, BOILERPLATE_DIR, TEMPLATE_DIR)
)
