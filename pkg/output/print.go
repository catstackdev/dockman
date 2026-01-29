// pkg/output/print.go
package output

import (
	"fmt"

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
