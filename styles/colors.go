package styles

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Color palette - using GitHub/VS Code inspired colors that work well
// on both dark and light terminal backgrounds

// Primary colors - main brand colors (purple/violet like GitHub Copilot)
var (
	// Primary accent colors
	PrimaryDark  = lipgloss.Color("#A371F7") // purple for dark mode
	PrimaryLight = lipgloss.Color("#8250DF") // darker purple for light mode

	// Secondary accent colors (teal/cyan)
	SecondaryDark  = lipgloss.Color("#58A6FF") // blue for dark mode
	SecondaryLight = lipgloss.Color("#0969DA") // darker blue for light mode
)

// Semantic colors - fixed colors for specific meanings
var (
	// Success - green tones
	SuccessDark  = lipgloss.Color("#3FB950") // bright green for dark
	SuccessLight = lipgloss.Color("#1A7F37") // darker green for light

	// Error - red tones
	ErrorDark  = lipgloss.Color("#F85149") // bright red for dark
	ErrorLight = lipgloss.Color("#CF222E") // darker red for light

	// Warning - yellow/orange tones
	WarningDark  = lipgloss.Color("#D29922") // golden for dark
	WarningLight = lipgloss.Color("#9A6700") // darker amber for light

	// Info - blue tones
	InfoDark  = lipgloss.Color("#58A6FF") // bright blue for dark
	InfoLight = lipgloss.Color("#0969DA") // darker blue for light
)

// Neutral colors - for text and backgrounds
var (
	// Muted text colors
	MutedDark  = lipgloss.Color("#8B949E") // gray for dark mode
	MutedLight = lipgloss.Color("#656D76") // darker gray for light mode

	// Subtle background/border colors
	SubtleDark  = lipgloss.Color("#30363D") // dark gray for dark mode
	SubtleLight = lipgloss.Color("#D0D7DE") // light gray for light mode

	// Foreground text colors
	FgDark  = lipgloss.Color("#E6EDF3") // white-ish for dark mode
	FgLight = lipgloss.Color("#1F2328") // near-black for light mode
)

// Helper functions to get the right color for current theme

func Primary() color.Color   { return Adaptive(PrimaryDark, PrimaryLight) }
func Secondary() color.Color { return Adaptive(SecondaryDark, SecondaryLight) }
func Success() color.Color   { return Adaptive(SuccessDark, SuccessLight) }
func Error() color.Color     { return Adaptive(ErrorDark, ErrorLight) }
func Warning() color.Color   { return Adaptive(WarningDark, WarningLight) }
func Info() color.Color      { return Adaptive(InfoDark, InfoLight) }
func Muted() color.Color     { return Adaptive(MutedDark, MutedLight) }
func Subtle() color.Color    { return Adaptive(SubtleDark, SubtleLight) }
func Fg() color.Color        { return Adaptive(FgDark, FgLight) }
