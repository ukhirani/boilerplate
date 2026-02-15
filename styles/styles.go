package styles

import "charm.land/lipgloss/v2"

// Common spacing constants
const (
	DefaultPadding = 1
	DefaultMargin  = 1
)

// Symbols for different message types - using simple ASCII-friendly symbols
var (
	SuccessSymbol = "✓"
	ErrorSymbol   = "✗"
	WarningSymbol = "!"
	InfoSymbol    = "•"
	ArrowSymbol   = "→"
	BulletSymbol  = "•"
	SpinnerSymbol = "◌"
)

// --- Base Style Builders ---

// SuccessStyle returns a style for success messages
func SuccessStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Success()).
		Bold(true)
}

// ErrorStyle returns a style for error messages
func ErrorStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Error()).
		Bold(true)
}

// WarningStyle returns a style for warning messages
func WarningStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Warning()).
		Bold(true)
}

// InfoStyle returns a style for info messages
func InfoStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Info())
}

// MutedStyle returns a style for secondary/muted text
func MutedStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Muted())
}

// PrimaryStyle returns a style with the primary accent color
func PrimaryStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Primary()).
		Bold(true)
}

// SecondaryStyle returns a style with the secondary accent color
func SecondaryStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Secondary())
}

// --- Component Styles ---

// HeaderStyle returns a style for section headers/titles
func HeaderStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Primary()).
		Bold(true).
		MarginBottom(1)
}

// TitleStyle returns a style for main titles
func TitleStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Primary()).
		Bold(true).
		Padding(0, 1)
}

// SubtitleStyle returns a style for subtitles
func SubtitleStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Muted()).
		Italic(true)
}

// CodeStyle returns a style for inline code/commands
func CodeStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Secondary()).
		Background(Subtle()).
		Padding(0, 1)
}

// PathStyle returns a style for file paths
func PathStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Secondary()).
		Underline(true)
}

// KeyStyle returns a style for key names (in key-value displays)
func KeyStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Muted())
}

// ValueStyle returns a style for values (in key-value displays)
func ValueStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Fg())
}

// --- Box/Container Styles ---

// BoxStyle returns a basic bordered box style
func BoxStyle() lipgloss.Style {
	return AdaptiveStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Subtle()).
		Padding(0, 1)
}

// SuccessBoxStyle returns a box style for success messages
func SuccessBoxStyle() lipgloss.Style {
	return AdaptiveStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Success()).
		Padding(0, 1)
}

// ErrorBoxStyle returns a box style for error messages
func ErrorBoxStyle() lipgloss.Style {
	return AdaptiveStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Error()).
		Padding(0, 1)
}

// WarningBoxStyle returns a box style for warning messages
func WarningBoxStyle() lipgloss.Style {
	return AdaptiveStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Warning()).
		Padding(0, 1)
}

// InfoBoxStyle returns a box style for info messages
func InfoBoxStyle() lipgloss.Style {
	return AdaptiveStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Info()).
		Padding(0, 1)
}

// --- List Item Styles ---

// ListItemStyle returns a style for list items
func ListItemStyle() lipgloss.Style {
	return AdaptiveStyle().
		PaddingLeft(2)
}

// SelectedItemStyle returns a style for selected/highlighted items
func SelectedItemStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Primary()).
		Bold(true).
		PaddingLeft(2)
}

// --- Progress/Status Styles ---

// ProgressStyle returns a style for progress indicators
func ProgressStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Info())
}

// StepStyle returns a style for step numbers/counters
func StepStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Muted()).
		Bold(true)
}

// CommandStyle returns a style for displaying commands being executed
func CommandStyle() lipgloss.Style {
	return AdaptiveStyle().
		Foreground(Secondary()).
		Italic(true)
}
