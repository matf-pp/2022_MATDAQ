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
	HeaderLogoStyle = lipgloss.NewStyle().
			Height(1).
			MarginLeft(1).
			PaddingLeft(1).
			PaddingRight(1).
			Bold(true).
			Background(lipgloss.Color("7")).
			Foreground(lipgloss.Color("#FAFAFA"))
	HeaderMoneyStyle = lipgloss.NewStyle().
				Align(lipgloss.Center).
				Height(1).
				PaddingLeft(1).
				PaddingRight(1).
				Bold(true).
				Background(lipgloss.Color("7")).
				Foreground(lipgloss.Color("#FAFAFA"))
)
