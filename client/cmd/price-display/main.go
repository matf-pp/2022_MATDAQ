package main

import (
	"fmt"
	matching_engine "github.com/matf-pp/2022_MATDAQ/client/internal/price-display/matching-engine"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/matf-pp/2022_MATDAQ/client/internal/price-display/tui"
)

func main() {
	orderResponses := make(chan *matching_engine.OrderResponse, 1024)
	tradeResponses := make(chan *matching_engine.TradeResponse, 1024)
	matching_engine.StartMatchingEngine(orderResponses, tradeResponses)

	m := tui.New()

	p := tea.NewProgram(&m, tea.WithAltScreen())

	go matching_engine.HandleBubbleTea(p, orderResponses, tradeResponses)

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
