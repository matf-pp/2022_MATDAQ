package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/matf-pp/2022_MATDAQ/internal/request-creator/tui"
)

func main() {
	fmt.Println("Request creator")

	m := tui.New()

	p := tea.NewProgram(&m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
