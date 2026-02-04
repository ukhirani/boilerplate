package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/ukhirani/boilerplate/types"
)

// Tests for ExecCmds
func TestExecCmds_EmptyCommands(t *testing.T) {
	// Test with empty command slice
	err := ExecCmds([]string{})
	if err != nil {
		t.Errorf("ExecCmds with empty slice returned error: %v", err)
	}
}

func TestExecCmds_SimpleCommand(t *testing.T) {
	// Test with a simple command that should always work
	cmds := []string{"echo hello"}
	err := ExecCmds(cmds)
	if err != nil {
		t.Errorf("ExecCmds(%v) returned error: %v", cmds, err)
	}
}

func TestExecCmds_InvalidCommand(t *testing.T) {
	// Test with an invalid command that should fail
	cmds := []string{"nonexistentcommand12345"}
	err := ExecCmds(cmds)
	if err == nil {
		t.Errorf("ExecCmds(%v) should have returned an error for invalid command", cmds)
	}
}

// Tests for InitViper
func TestInitViper_SetsConfigType(t *testing.T) {
	// Reset viper before test
	viper.Reset()

	InitViper()

	// Check that config type is set (indirectly by checking defaults)
	// InitViper sets defaults from types.DefaultConfig
	name := viper.GetString("name")
	if name != "default-template-name" {
		t.Errorf("InitViper did not set default name, got %s, want default-template-name", name)
	}
}

func TestInitViper_SetsDefaults(t *testing.T) {
	viper.Reset()

	InitViper()

	// Check that isDir default is false
	isDir := viper.GetBool("isDir")
	if isDir != false {
		t.Errorf("InitViper default isDir = %v, want false", isDir)
	}

	// Check that preCmd default is empty slice
	preCmd := viper.GetStringSlice("preCmd")
	if len(preCmd) != 0 {
		t.Errorf("InitViper default preCmd length = %d, want 0", len(preCmd))
	}

	// Check that postCmd default is empty slice
	postCmd := viper.GetStringSlice("postCmd")
	if len(postCmd) != 0 {
		t.Errorf("InitViper default postCmd length = %d, want 0", len(postCmd))
	}
}

// Tests for ListDir
func TestListDir_EmptyDirectory(t *testing.T) {
	// Create a temporary empty directory
	tmpDir, err := os.MkdirTemp("", "test_listdir_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	count, err := ListDir(tmpDir, false)
	if err != nil {
		t.Errorf("ListDir(%s, false) returned error: %v", tmpDir, err)
	}
	if count != 0 {
		t.Errorf("ListDir(%s, false) returned count = %d, want 0", tmpDir, count)
	}
}

func TestListDir_WithFiles(t *testing.T) {
	// Create a temporary directory with some files
	tmpDir, err := os.MkdirTemp("", "test_listdir_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create some files
	for i := 0; i < 3; i++ {
		filePath := filepath.Join(tmpDir, "file"+string(rune('a'+i))+".txt")
		err := os.WriteFile(filePath, []byte("content"), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	count, err := ListDir(tmpDir, false)
	if err != nil {
		t.Errorf("ListDir(%s, false) returned error: %v", tmpDir, err)
	}
	if count != 3 {
		t.Errorf("ListDir(%s, false) returned count = %d, want 3", tmpDir, count)
	}
}

func TestListDir_NonExistingDirectory(t *testing.T) {
	nonExistingPath := "/tmp/non_existing_dir_12345678"

	_, err := ListDir(nonExistingPath, false)
	if err == nil {
		t.Errorf("ListDir(%s, false) should have returned an error for non-existing directory", nonExistingPath)
	}
}

// Tests for ReadConfig
func TestReadConfig_NonExistingConfig(t *testing.T) {
	viper.Reset()
	InitViper()

	var conf types.Config
	err := ReadConfig("non_existing_template_12345", &conf)

	// Should return an error for non-existing config
	if err == nil {
		t.Errorf("ReadConfig should have returned an error for non-existing config")
	}
}
