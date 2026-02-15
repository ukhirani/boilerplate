package styles

import (
	"fmt"
	"strings"
)

// --- Message Printing Helpers ---

// PrintSuccess prints a success message with the success symbol
func PrintSuccess(message string) {
	symbol := SuccessStyle().Render(SuccessSymbol)
	fmt.Printf("%s %s\n", symbol, message)
}

// PrintError prints an error message with the error symbol
func PrintError(message string) {
	symbol := ErrorStyle().Render(ErrorSymbol)
	label := ErrorStyle().Render("Error:")
	fmt.Printf("%s %s %s\n", symbol, label, message)
}

// PrintWarning prints a warning message with the warning symbol
func PrintWarning(message string) {
	symbol := WarningStyle().Render(WarningSymbol)
	label := WarningStyle().Render("Warning:")
	fmt.Printf("%s %s %s\n", symbol, label, message)
}

// PrintInfo prints an info message with the info symbol
func PrintInfo(message string) {
	symbol := InfoStyle().Render(InfoSymbol)
	fmt.Printf("%s %s\n", symbol, message)
}

// PrintMuted prints muted/secondary text
func PrintMuted(message string) {
	fmt.Println(MutedStyle().Render(message))
}

// --- Structured Message Helpers ---

// PrintErrorWithDetails prints an error with additional details
func PrintErrorWithDetails(message string, details ...string) {
	PrintError(message)
	for _, detail := range details {
		fmt.Printf("  %s %s\n", MutedStyle().Render(ArrowSymbol), MutedStyle().Render(detail))
	}
}

// PrintSuccessWithDetails prints a success with additional details
func PrintSuccessWithDetails(message string, details ...string) {
	PrintSuccess(message)
	for _, detail := range details {
		fmt.Printf("  %s %s\n", MutedStyle().Render(ArrowSymbol), MutedStyle().Render(detail))
	}
}

// PrintWarningWithDetails prints a warning with additional details
func PrintWarningWithDetails(message string, details ...string) {
	PrintWarning(message)
	for _, detail := range details {
		fmt.Printf("  %s %s\n", MutedStyle().Render(ArrowSymbol), MutedStyle().Render(detail))
	}
}

// --- Header/Section Helpers ---

// PrintHeader prints a styled header
func PrintHeader(title string) {
	fmt.Println()
	fmt.Println(HeaderStyle().Render(title))
}

// PrintSubHeader prints a styled subheader
func PrintSubHeader(title string) {
	fmt.Println(SubtitleStyle().Render(title))
}

// PrintTitle prints a main title with decoration
func PrintTitle(title string) {
	styled := TitleStyle().Render(title)
	fmt.Println()
	fmt.Println(styled)
}

// --- List Helpers ---

// PrintList prints a bulleted list of items
func PrintList(items []string) {
	bullet := MutedStyle().Render(BulletSymbol)
	for _, item := range items {
		fmt.Printf("  %s %s\n", bullet, item)
	}
}

// PrintNumberedList prints a numbered list of items
func PrintNumberedList(items []string) {
	for i, item := range items {
		num := StepStyle().Render(fmt.Sprintf("%d.", i+1))
		fmt.Printf("  %s %s\n", num, item)
	}
}

// --- Key-Value Helpers ---

// PrintKeyValue prints a key-value pair
func PrintKeyValue(key, value string) {
	k := KeyStyle().Render(key + ":")
	v := ValueStyle().Render(value)
	fmt.Printf("  %s %s\n", k, v)
}

// PrintKeyValueInline prints a key-value pair inline (no indent)
func PrintKeyValueInline(key, value string) {
	k := KeyStyle().Render(key + ":")
	v := ValueStyle().Render(value)
	fmt.Printf("%s %s\n", k, v)
}

// --- Progress/Step Helpers ---

// PrintStep prints a step in a multi-step process
func PrintStep(current, total int, message string) {
	step := StepStyle().Render(fmt.Sprintf("[%d/%d]", current, total))
	fmt.Printf("%s %s\n", step, message)
}

// PrintCommand prints a command that's being executed
func PrintCommand(cmd string) {
	arrow := MutedStyle().Render(ArrowSymbol)
	command := CommandStyle().Render(cmd)
	fmt.Printf("  %s %s\n", arrow, command)
}

// PrintRunning prints a "running" status message for a command
func PrintRunning(message string) {
	symbol := InfoStyle().Render(SpinnerSymbol)
	fmt.Printf("%s %s\n", symbol, message)
}

// --- Template/Item Display Helpers ---

// PrintTemplateItem prints a template entry with type indicator
func PrintTemplateItem(name string, isDir bool) {
	var typeIndicator string
	if isDir {
		typeIndicator = PrimaryStyle().Render("[DIR]")
	} else {
		typeIndicator = SecondaryStyle().Render("[FILE]")
	}
	fmt.Printf("  %s %s\n", typeIndicator, name)
}

// PrintTreeItem prints an item in a tree structure
func PrintTreeItem(name string, isLast bool) {
	var prefix string
	if isLast {
		prefix = MutedStyle().Render("└─")
	} else {
		prefix = MutedStyle().Render("├─")
	}
	fmt.Printf("  %s %s\n", prefix, name)
}

// --- Path Helper ---

// PrintPath prints a styled file path
func PrintPath(path string) {
	fmt.Println(PathStyle().Render(path))
}

// PrintPathWithLabel prints a file path with a label
func PrintPathWithLabel(label, path string) {
	l := KeyStyle().Render(label + ":")
	p := PathStyle().Render(path)
	fmt.Printf("  %s %s\n", l, p)
}

// --- Divider/Spacing Helpers ---

// PrintDivider prints a subtle divider line
func PrintDivider() {
	divider := MutedStyle().Render(strings.Repeat("─", 40))
	fmt.Println(divider)
}

// PrintNewLine prints an empty line
func PrintNewLine() {
	fmt.Println()
}

// --- Box Helpers ---

// PrintInBox prints content in a bordered box
func PrintInBox(content string) {
	fmt.Println(BoxStyle().Render(content))
}

// PrintSuccessBox prints content in a success-colored box
func PrintSuccessBox(content string) {
	fmt.Println(SuccessBoxStyle().Render(content))
}

// PrintErrorBox prints content in an error-colored box
func PrintErrorBox(content string) {
	fmt.Println(ErrorBoxStyle().Render(content))
}

// PrintWarningBox prints content in a warning-colored box
func PrintWarningBox(content string) {
	fmt.Println(WarningBoxStyle().Render(content))
}

// PrintInfoBox prints content in an info-colored box
func PrintInfoBox(content string) {
	fmt.Println(InfoBoxStyle().Render(content))
}

// --- Inline Styling Helpers ---

// Bold returns the text styled as bold
func Bold(text string) string {
	return AdaptiveStyle().Bold(true).Render(text)
}

// Italic returns the text styled as italic
func Italic(text string) string {
	return AdaptiveStyle().Italic(true).Render(text)
}

// Highlight returns the text with primary color
func Highlight(text string) string {
	return PrimaryStyle().Render(text)
}

// Dim returns the text with muted color
func Dim(text string) string {
	return MutedStyle().Render(text)
}

// Code returns the text styled as inline code
func Code(text string) string {
	return CodeStyle().Render(text)
}

// Path returns the text styled as a file path
func Path(text string) string {
	return PathStyle().Render(text)
}
