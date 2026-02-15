/*
Copyright © 2025 Umang Hirani umanghirani.exe@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	degit "github.com/qiushiyan/degit/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ukhirani/boilerplate/constants"
	"github.com/ukhirani/boilerplate/styles"
	"github.com/ukhirani/boilerplate/types"
	"github.com/ukhirani/boilerplate/utils"
)

const (
	hubAPIEndpoint = "https://bp-hub-render-service.onrender.com/getTemplates"
)

var aliasName string

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone <username/template-name>",
	Short: "Clone a template from the bp-hub",
	Long: `Clone a template from the bp-hub repository using username/template-name format.

The template will be saved locally with the alias name you provide.

Usage:
  bp clone <username/template-name> --alias <local-template-name>

Examples:
  bp clone ukhirani/cpp-template --alias cpp-starter
  bp clone ukhirani/react-tailwind -a tailwind-starter`,
	Args:    cobra.ExactArgs(1),
	Run:     CloneCmdRunner,
	Aliases: []string{"install", "get"},
	Example: `bp clone ukhirani/cpp-template --alias cpp-starter
  bp clone ukhirani/react-tailwind -a tailwind-starter`,
}

// fetchTemplates fetches all templates from the bp-hub API
func fetchTemplates() ([]types.HubTemplate, error) {
	resp, err := http.Get(hubAPIEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch templates from hub: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("hub API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var templates []types.HubTemplate
	if err := json.Unmarshal(body, &templates); err != nil {
		return nil, fmt.Errorf("failed to parse templates JSON: %w", err)
	}

	return templates, nil
}

func createViperConfig(template *types.HubTemplate, alias, templateDir, configPath string, isDir bool) error {
	v := viper.New()
	v.SetConfigType("toml")
	v.Set("Name", alias)
	v.Set("IsDir", isDir)
	v.Set("PreCmd", template.PreCmds)
	v.Set("PostCmd", template.PostCmds)

	// Write the config file
	if err := v.WriteConfigAs(configPath); err != nil {
		// Cleanup on failure
		os.RemoveAll(templateDir)
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func isDirType(template *types.HubTemplate) bool {
	return strings.EqualFold(template.Type, "dir")
}

// findTemplate finds a template by username and template name
func findTemplate(templates []types.HubTemplate, username, templateName string) *types.HubTemplate {
	for _, t := range templates {
		if strings.EqualFold(t.Username, username) && strings.EqualFold(t.TemplateName, templateName) {
			return &t
		}
	}
	return nil
}

// parseTemplateArg parses the username/template-name argument
func parseTemplateArg(arg string) (username, templateName string, err error) {
	// Find the first "/" to split username and template name
	idx := strings.Index(arg, "/")
	if idx == -1 {
		return "", "", fmt.Errorf("invalid format: expected 'username/template-name', got '%s'", arg)
	}

	username = strings.TrimSpace(arg[:idx])
	templateName = strings.TrimSpace(arg[idx+1:])

	if username == "" {
		return "", "", fmt.Errorf("username cannot be empty")
	}
	if templateName == "" {
		return "", "", fmt.Errorf("template name cannot be empty")
	}

	return username, templateName, nil
}

// createLocalTemplate creates the local template with the cloned code
func createLocalTemplate(template *types.HubTemplate, alias string) error {
	// Create the template directory
	templateDir := filepath.Join(constants.BOILERPLATE_TEMPLATE_DIR, alias)
	isDir := isDirType(template)

	// Check if template already exists
	if utils.Exists(templateDir) {
		return fmt.Errorf("template with alias '%s' already exists", alias)
	}

	if isDir {
		if err := degit.Clone(template.GithubRepoLink, templateDir, false, false); err != nil {
			return fmt.Errorf("failed to clone repository: %w", err)
		}
	} else {
		// Create the template directory
		if err := os.MkdirAll(templateDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create template directory: %w", err)
		}

		// Determine the file extension based on the code content or template type
		fileName := determineFileName(template)

		// Create the template file with the code
		filePath := filepath.Join(templateDir, fileName)
		if err := os.WriteFile(filePath, []byte(template.Code), 0o644); err != nil {
			// Cleanup on failure
			os.RemoveAll(templateDir)
			return fmt.Errorf("failed to write template file: %w", err)
		}
	}

	// Create the config file using viper
	configPath := filepath.Join(constants.BOILERPLATE_CONFIG_DIR, alias+".toml")

	// WARN: here isDir is always false as of now for prototyping purposes

	// Create a new viper instance for this config
	createViperConfig(template, alias, templateDir, configPath, isDir)

	return nil
}

// TODO: add a filename field in the database

// determineFileName attempts to determine an appropriate file name based on the template
func determineFileName(template *types.HubTemplate) string {
	// Check for common patterns in the code to determine file extension
	code := template.Code

	// Check for various language patterns
	switch {
	case strings.Contains(code, "import React") || strings.Contains(code, "from 'react'"):
		return "index.jsx"
	case strings.Contains(code, "from 'vue'") || strings.Contains(code, "defineComponent"):
		return "index.vue"
	case strings.Contains(code, "package main") || strings.Contains(code, "func main()"):
		return "main.go"
	case strings.Contains(code, "def ") && strings.Contains(code, ":"):
		return "main.py"
	case strings.Contains(code, "#!/bin/bash") || strings.Contains(code, "#!/bin/sh"):
		return "script.sh"
	case strings.Contains(code, "<!DOCTYPE html") || strings.Contains(code, "<html"):
		return "index.html"
	case strings.Contains(code, "function") || strings.Contains(code, "const ") || strings.Contains(code, "export "):
		if strings.Contains(code, ": ") && (strings.Contains(code, "interface ") || strings.Contains(code, "type ")) {
			return "index.ts"
		}
		return "index.js"
	case strings.Contains(code, "#include"):
		if strings.Contains(code, "iostream") || strings.Contains(code, "std::") {
			return "main.cpp"
		}
		return "main.c"
	case strings.Contains(code, "class ") && strings.Contains(code, "public static void main"):
		return "Main.java"
	case strings.Contains(code, "fn main()") || strings.Contains(code, "use std::"):
		return "main.rs"
	default:
		// Default to a text file if we can't determine the type
		return "template.txt"
	}
}

func CloneCmdRunner(cmd *cobra.Command, args []string) {
	// Parse the username/template-name argument
	username, templateName, err := parseTemplateArg(args[0])
	if err != nil {
		styles.PrintErrorWithDetails(
			err.Error(),
			"Usage: "+styles.Code("bp clone <username/template-name> --alias <local-name>"),
		)
		os.Exit(1)
	}

	// Validate the alias name
	if !utils.IsValidDirName(aliasName) {
		styles.PrintErrorWithDetails(
			"Invalid alias name: "+styles.Highlight(aliasName),
			"Allowed characters: letters, numbers, and underscores only",
		)
		os.Exit(1)
	}

	// Check if template with alias already exists
	if exists, _ := utils.IsTemplateExists(aliasName); exists {
		styles.PrintErrorWithDetails(
			"Template with alias "+styles.Highlight(aliasName)+" already exists",
			"Use a different alias name or remove the existing template",
		)
		os.Exit(1)
	}

	styles.PrintRunning("Fetching templates from bp-hub...")

	// Fetch templates from the hub
	templates, err := fetchTemplates()
	if err != nil {
		styles.PrintErrorWithDetails(
			err.Error(),
			"Please check your internet connection and try again",
		)
		os.Exit(1)
	}

	// Find the requested template
	template := findTemplate(templates, username, templateName)
	if template == nil {
		styles.PrintErrorWithDetails(
			"Template "+styles.Highlight(username+"/"+templateName)+" not found in bp-hub",
			"Make sure the username and template name are correct",
		)
		os.Exit(1)
	}

	styles.PrintNewLine()
	styles.PrintInfo("Found: " + styles.Highlight(template.TemplateName) + " by " + styles.Dim(template.Username))
	styles.PrintKeyValue("Description", template.Description)
	styles.PrintKeyValue("Stats", fmt.Sprintf("%d stars • %d clones", template.Stars, template.Clones))

	// Create the local template
	if err := createLocalTemplate(template, aliasName); err != nil {
		styles.PrintError("Failed to create local template: " + err.Error())
		os.Exit(1)
	}

	styles.PrintNewLine()
	styles.PrintSuccess("Template cloned successfully!")
	styles.PrintKeyValue("Local name", aliasName)
	styles.PrintInfo("Run " + styles.Code("bp gen "+aliasName) + " to use this template")

	// Show pre/post commands if any
	if len(template.PreCmds) > 0 {
		styles.PrintNewLine()
		styles.PrintSubHeader("Pre-commands configured")
		styles.PrintList(template.PreCmds)
	}
	if len(template.PostCmds) > 0 {
		styles.PrintNewLine()
		styles.PrintSubHeader("Post-commands configured")
		styles.PrintList(template.PostCmds)
	}
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	// Define the --alias / -a flag
	cloneCmd.Flags().StringVarP(&aliasName, "alias", "a", "", "Local template name (required)")

	// Mark the alias flag as required
	if err := cloneCmd.MarkFlagRequired("alias"); err != nil {
		styles.PrintError(err.Error())
	}
}
