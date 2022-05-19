package main

import (
	"fmt"
	"log"
	"net"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/matf-pp/2022_MATDAQ/client/internal/request-creator/tui"
)

func main() {
	conn, err := net.Dial("tcp", "request-creator-server:8081")
	if err != nil {
		fmt.Println("Dial failed", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	log.Println("Connected to request creator server")

	m := tui.New(conn)
	p := tea.NewProgram(&m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		log.Fatalln(err)
	}
}
