package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Color constants matching the specified design
const (
	ColorMatrixGreen = "#00FF00"
	ColorAlertRed    = "#FF0000"
	ColorTermBlack   = "#000000"
	ColorSoftGreen   = "#33FF33"
	ColorCyberBlue   = "#0066FF"
)

// StyleConfig contains all styles for the UI
type StyleConfig struct {
	Title      lipgloss.Style
	Subtitle   lipgloss.Style
	Normal     lipgloss.Style
	Highlight  lipgloss.Style
	Info       lipgloss.Style
	Error      lipgloss.Style
	Success    lipgloss.Style
	Terminal   lipgloss.Style
	StatusBar  lipgloss.Style
	ProgressBar lipgloss.Style
	CommandLine lipgloss.Style
	Panel      lipgloss.Style
	Button     lipgloss.Style
	ButtonActive lipgloss.Style
	ListItem   lipgloss.Style
	ListItemSelected lipgloss.Style
}

// DefaultStyles returns the default style configuration
func DefaultStyles() StyleConfig {
	return StyleConfig{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(ColorMatrixGreen)).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(ColorMatrixGreen)).
			Padding(1, 2),

		Subtitle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorSoftGreen)).
			Bold(true),

		Normal: lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorSoftGreen)),

		Highlight: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(ColorMatrixGreen)),

		Info: lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorCyberBlue)),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorAlertRed)).
			Bold(true),

		Success: lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMatrixGreen)).
			Bold(true),

		Terminal: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(ColorSoftGreen)).
			Padding(1, 2),

		StatusBar: lipgloss.NewStyle().
			Background(lipgloss.Color(ColorSoftGreen)).
			Foreground(lipgloss.Color(ColorTermBlack)).
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1),

		ProgressBar: lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMatrixGreen)),

		CommandLine: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(ColorSoftGreen)).
			Padding(0, 1),

		Panel: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(ColorSoftGreen)).
			Padding(1, 2),

		Button: lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorSoftGreen)).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(ColorSoftGreen)).
			Padding(0, 3),

		ButtonActive: lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMatrixGreen)).
			Background(lipgloss.Color(ColorTermBlack)).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(ColorMatrixGreen)).
			Bold(true).
			Padding(0, 3),

		ListItem: lipgloss.NewStyle().
			Padding(0, 2),

		ListItemSelected: lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMatrixGreen)).
			Background(lipgloss.Color(ColorTermBlack)).
			Bold(true).
			Padding(0, 2),
	}
}
