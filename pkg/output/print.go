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
	width := 50
	fmt.Printf("\n╭%s╮\n", strings.Repeat("─", width))
	fmt.Printf("│ %s%s │\n", title, strings.Repeat(" ", width-len(title)-2))
	fmt.Printf("├%s┤\n", strings.Repeat("─", width))
	for _, line := range lines {
		padding := width - len(line) - 2
		if padding < 0 {
			padding = 0
		}
		fmt.Printf("│ %s%s │\n", line, strings.Repeat(" ", padding))
	}
	fmt.Printf("╰%s╯\n\n", strings.Repeat("─", width))
}
