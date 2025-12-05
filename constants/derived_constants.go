package constants

import (
	"os"
	"path/filepath"
)

var HOME_DIR, _ = os.UserHomeDir()
var USER_CONFIG_DIR, _ = os.UserConfigDir()
var CURR_DIR, _ = os.Getwd()

// Boilerplate Constants
var BOILERPLATE_CONFIG_DIR = filepath.Join(HOME_DIR, BOILERPLATE_DIR, CONFIG_DIR)
var BOILERPLATE_TEMPLATE_DIR = filepath.Join(HOME_DIR, BOILERPLATE_DIR, TEMPLATE_DIR)
