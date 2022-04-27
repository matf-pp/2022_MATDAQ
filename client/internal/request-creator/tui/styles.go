package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	BubbleStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63"))
	TextInputStyle = lipgloss.NewStyle().
			Width(23).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63"))
	HeaderStyle = lipgloss.NewStyle().
			Height(1).
			MarginLeft(1).
			PaddingLeft(1).
			PaddingRight(1).
			Bold(true).
			Background(lipgloss.Color("7")).
			Foreground(lipgloss.Color("#FAFAFA"))
)
