package ui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

// Styles contains the global styles configuration
var Styles = DefaultStyles()

// Terminal width constants
const (
	TerminalWidth  = 80
	TerminalHeight = 24
)

// RenderTitle creates a formatted title block
func RenderTitle(title string) string {
	return Styles.Title.Render(title)
}

// RenderSubtitle creates a formatted subtitle
func RenderSubtitle(subtitle string) string {
	return Styles.Subtitle.Render(subtitle)
}

// RenderStatusBar creates a status bar with the given content
func RenderStatusBar(content string, width int) string {
	return Styles.StatusBar.Width(width).Render(content)
}

// RenderPanelWithTitle creates a panel with a title and content
func RenderPanelWithTitle(title, content string, width int) string {
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorMatrixGreen)).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color(ColorSoftGreen)).
		Padding(0, 1).
		Width(width - 4)

	contentStyle := lipgloss.NewStyle().
		Width(width - 4).
		Padding(1, 2, 0, 2)

	panel := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(title),
		contentStyle.Render(content),
	)

	return Styles.Panel.Width(width).Render(panel)
}

// RenderCommandPrompt creates a command prompt with the given prompt and input
func RenderCommandPrompt(prompt, input string, width int) string {
	promptStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorMatrixGreen)).
		Bold(true)

	inputStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorSoftGreen))

	commandLine := promptStyle.Render(prompt) + inputStyle.Render(input)
	
	return Styles.CommandLine.Width(width).Render(commandLine)
}

// RenderProgressBar creates a progress bar with the given percent
func RenderProgressBar(percent float64, width int) string {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(width-4),
	)
	
	return Styles.ProgressBar.Render(p.ViewAs(percent))
}

// RenderList creates a list with the given items, and optional selection
func RenderList(items []string, selectedIdx int) string {
	var listItems []string
	
	for i, item := range items {
		if i == selectedIdx {
			listItems = append(listItems, Styles.ListItemSelected.Render("▶ "+item))
		} else {
			listItems = append(listItems, Styles.ListItem.Render("  "+item))
		}
	}
	
	return strings.Join(listItems, "\n")
}

// NewSpinner creates a new spinner with the default style
func NewSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorMatrixGreen))
	return s
}

// AnimateText returns a function that reveals text character by character
func AnimateText(text string, delay time.Duration) func() (string, bool) {
	var currentPos int
	
	return func() (string, bool) {
		if currentPos >= len(text) {
			return text, true
		}
		
		currentPos++
		return text[:currentPos], false
	}
}

// CenterText centers the given text within the specified width
func CenterText(text string, width int) string {
	textWidth := lipgloss.Width(text)
	if textWidth >= width {
		return text
	}
	
	padding := (width - textWidth) / 2
	return strings.Repeat(" ", padding) + text
}
