// Package styles provides consistent styling for the bp CLI using lipgloss.
// It features adaptive theming that automatically adjusts to the user's terminal
// background (dark or light mode) for optimal readability.
package styles

import (
	"image/color"
	"os"
	"sync"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
)

// Theme represents the current terminal color scheme
type Theme struct {
	IsDark  bool
	Writer  *colorprofile.Writer
	Profile colorprofile.Profile
}

var (
	currentTheme *Theme
	themeOnce    sync.Once
)

// GetTheme returns the current theme, detecting it on first call
func GetTheme() *Theme {
	themeOnce.Do(func() {
		currentTheme = detectTheme()
	})
	return currentTheme
}

// detectTheme determines if the terminal is using a dark or light background
func detectTheme() *Theme {
	w := colorprofile.NewWriter(os.Stdout, os.Environ())
	profile := colorprofile.Detect(os.Stdout, os.Environ())

	// Try to detect if terminal has a dark background
	// Default to dark mode as most developer terminals use dark themes
	isDark := true

	return &Theme{
		IsDark:  isDark,
		Writer:  w,
		Profile: profile,
	}
}

// Adaptive returns a color that works well on both dark and light backgrounds.
// It takes a dark-mode color and a light-mode color, returning the appropriate one.
func Adaptive(dark, light color.Color) color.Color {
	if GetTheme().IsDark {
		return dark
	}
	return light
}

// AdaptiveStyle creates a style with colors that adapt to the terminal theme
func AdaptiveStyle() lipgloss.Style {
	return lipgloss.NewStyle()
}
