package main

import (
	"fmt"
	"net"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/matf-pp/2022_MATDAQ/client/internal/request-creator/tui"
)

func main() {
	fmt.Println("Request creator")

	conn, err := net.Dial("tcp", "request-creator-server:8081")
	if err != nil {
		fmt.Println("Dial failed", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	m := tui.New(conn)

	p := tea.NewProgram(&m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
