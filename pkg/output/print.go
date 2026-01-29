// pkg/output/print.go
package output

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

// Info prints info message
func Info(msg string) {
	fmt.Printf("%s %s\n", cyan("ℹ"), msg)
}

// Success prints success message
func Success(msg string) {
	fmt.Printf("%s %s\n", green("✓"), msg)
}

// Error prints error message
func Error(msg string) {
	fmt.Printf("%s %s\n", red("✗"), msg)
}

// Warning prints warning message
func Warning(msg string) {
	fmt.Printf("%s %s\n", yellow("⚠"), msg)
}

// FormatPresetName formats a preset name with color
func FormatPresetName(name string) string {
	return cyan("●") + " " + name
}

// Gray prints gray text
func Gray(text string) string {
	gray := color.New(color.FgHiBlack).SprintFunc()
	return gray(text)
}

// Cyan returns cyan colored text (useful for prompts)
func Cyan(text string) string {
	return cyan(text)
}

// ErrorWithHelp prints error with helpful suggestion
func ErrorWithHelp(msg string, help string) {
	Error(msg)
	if help != "" {
		fmt.Printf("  %s %s\n", yellow("💡"), Gray(help))
	}
}

// Box prints a boxed message
func Box(title string, lines []string) {
	width := 60

	// Top border
	fmt.Printf("\n╭%s╮\n", strings.Repeat("─", width))

	// Title
	titlePadding := width - len(title) - 2
	if titlePadding < 0 {
		titlePadding = 0
	}
	fmt.Printf("│ %s%s%s │\n",
		green("✓"),
		" "+title,
		strings.Repeat(" ", titlePadding-2))

	// Separator
	fmt.Printf("├%s┤\n", strings.Repeat("─", width))

	// Content lines
	for _, line := range lines {
		// Handle color codes in length calculation
		visibleLen := len(stripANSI(line))
		padding := width - visibleLen - 2
		if padding < 0 {
			padding = 0
		}
		fmt.Printf("│ %s%s │\n", line, strings.Repeat(" ", padding))
	}

	// Bottom border
	fmt.Printf("╰%s╯\n\n", strings.Repeat("─", width))
}

// stripANSI removes ANSI color codes for length calculation
func stripANSI(s string) string {
	// Simple ANSI stripper - remove escape sequences
	var result strings.Builder
	inEscape := false

	for _, r := range s {
		if r == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if r == 'm' {
				inEscape = false
			}
			continue
		}
		result.WriteRune(r)
	}

	return result.String()
}
