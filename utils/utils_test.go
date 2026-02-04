package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ukhirani/boilerplate/constants"
)

// Tests for Exists function
func TestExists_ExistingFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_exists_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	if !Exists(tmpFile.Name()) {
		t.Errorf("Exists(%s) = false, want true", tmpFile.Name())
	}
}

func TestExists_ExistingDirectory(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_exists_dir_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	if !Exists(tmpDir) {
		t.Errorf("Exists(%s) = false, want true", tmpDir)
	}
}

func TestExists_NonExistingPath(t *testing.T) {
	nonExistingPath := filepath.Join(os.TempDir(), "non_existing_file_12345678.txt")
	if Exists(nonExistingPath) {
		t.Errorf("Exists(%s) = true, want false", nonExistingPath)
	}
}

// Tests for IsDirectory function
func TestIsDirectory_Directory(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_isdir_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	isDir, err := IsDirectory(tmpDir)
	if err != nil {
		t.Errorf("IsDirectory(%s) returned error: %v", tmpDir, err)
	}
	if !isDir {
		t.Errorf("IsDirectory(%s) = false, want true", tmpDir)
	}
}

func TestIsDirectory_File(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_isdir_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	isDir, err := IsDirectory(tmpFile.Name())
	if err != nil {
		t.Errorf("IsDirectory(%s) returned error: %v", tmpFile.Name(), err)
	}
	if isDir {
		t.Errorf("IsDirectory(%s) = true, want false", tmpFile.Name())
	}
}

// Tests for IsValidDirName function
func TestIsValidDirName_ValidNames(t *testing.T) {
	validNames := []string{
		"myproject",
		"my-project",
		"my_project",
		"MyProject123",
		"test",
	}

	for _, name := range validNames {
		if !IsValidDirName(name) {
			t.Errorf("IsValidDirName(%s) = false, want true", name)
		}
	}
}

func TestIsValidDirName_InvalidNames(t *testing.T) {
	invalidNames := []string{
		"",
		"my/project",
		"my\x00project",
	}

	for _, name := range invalidNames {
		if IsValidDirName(name) {
			t.Errorf("IsValidDirName(%s) = true, want false", name)
		}
	}
}

// Tests for GetConfigFileLocation function
func TestGetConfigFileLocation(t *testing.T) {
	templateName := "my-template"
	expected := filepath.Join(constants.BOILERPLATE_CONFIG_DIR, templateName+".toml")

	result := GetConfigFileLocation(templateName)

	if result != expected {
		t.Errorf("GetConfigFileLocation(%s) = %s, want %s", templateName, result, expected)
	}
}

func TestGetConfigFileLocation_EndsWithToml(t *testing.T) {
	templateName := "test-template"
	result := GetConfigFileLocation(templateName)

	if !strings.HasSuffix(result, ".toml") {
		t.Errorf("GetConfigFileLocation(%s) = %s, should end with .toml", templateName, result)
	}
}

func TestGetConfigFileLocation_ContainsTemplateName(t *testing.T) {
	templateName := "unique-template-name"
	result := GetConfigFileLocation(templateName)

	if !strings.Contains(result, templateName) {
		t.Errorf("GetConfigFileLocation(%s) = %s, should contain the template name", templateName, result)
	}
}

// Tests for CopyDir function
func TestCopyDir_Success(t *testing.T) {
	srcDir, err := os.MkdirTemp("", "test_copydir_src_*")
	if err != nil {
		t.Fatalf("Failed to create temp source dir: %v", err)
	}
	defer os.RemoveAll(srcDir)

	testFile := filepath.Join(srcDir, "test.txt")
	err = os.WriteFile(testFile, []byte("hello world"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	destDir := filepath.Join(os.TempDir(), "test_copydir_dest_"+filepath.Base(srcDir))
	defer os.RemoveAll(destDir)

	err = CopyDir(srcDir, destDir)
	if err != nil {
		t.Errorf("CopyDir(%s, %s) returned error: %v", srcDir, destDir, err)
	}

	if !Exists(destDir) {
		t.Errorf("Destination directory %s was not created", destDir)
	}

	copiedFile := filepath.Join(destDir, "test.txt")
	if !Exists(copiedFile) {
		t.Errorf("File %s was not copied", copiedFile)
	}

	content, err := os.ReadFile(copiedFile)
	if err != nil {
		t.Fatalf("Failed to read copied file: %v", err)
	}
	if string(content) != "hello world" {
		t.Errorf("Copied file content = %q, want %q", string(content), "hello world")
	}
}

// Tests for CopyFile function
func TestCopyFile_Success(t *testing.T) {
	srcFile, err := os.CreateTemp("", "test_copyfile_src_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp source file: %v", err)
	}
	srcPath := srcFile.Name()
	defer os.Remove(srcPath)

	testContent := "test content for copy"
	_, err = srcFile.WriteString(testContent)
	if err != nil {
		t.Fatalf("Failed to write to source file: %v", err)
	}
	srcFile.Close()

	destDir, err := os.MkdirTemp("", "test_copyfile_dest_*")
	if err != nil {
		t.Fatalf("Failed to create temp dest dir: %v", err)
	}
	defer os.RemoveAll(destDir)

	newDestDir := filepath.Join(destDir, "subdir")
	destFileName := "copied_file.txt"

	err = CopyFile(srcPath, newDestDir, destFileName)
	if err != nil {
		t.Errorf("CopyFile(%s, %s, %s) returned error: %v", srcPath, newDestDir, destFileName, err)
	}

	copiedPath := filepath.Join(newDestDir, destFileName)
	if !Exists(copiedPath) {
		t.Errorf("Copied file %s does not exist", copiedPath)
	}

	content, err := os.ReadFile(copiedPath)
	if err != nil {
		t.Fatalf("Failed to read copied file: %v", err)
	}
	if string(content) != testContent {
		t.Errorf("Copied file content = %q, want %q", string(content), testContent)
	}
}

func TestCopyFile_FileAlreadyExists(t *testing.T) {
	srcFile, err := os.CreateTemp("", "test_copyfile_src_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp source file: %v", err)
	}
	srcPath := srcFile.Name()
	defer os.Remove(srcPath)
	srcFile.Close()

	destDir, err := os.MkdirTemp("", "test_copyfile_dest_*")
	if err != nil {
		t.Fatalf("Failed to create temp dest dir: %v", err)
	}
	defer os.RemoveAll(destDir)

	destFileName := "existing_file.txt"
	existingFile := filepath.Join(destDir, destFileName)
	err = os.WriteFile(existingFile, []byte("existing content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create existing file: %v", err)
	}

	err = CopyFile(srcPath, destDir, destFileName)
	if err == nil {
		t.Errorf("CopyFile should have returned an error when destination file already exists")
	}
}
