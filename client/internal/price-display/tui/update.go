package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			// gets selected stock data
			s := m.stocks[m.list.Index()]
			m.sellPrice = s.GetSellPrices()
			m.sellAmount = s.sellAmount
			m.buyPrice = s.buyPrices
			m.buyAmount = s.buyAmount
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height - 3
		m.windowWidth = msg.Width/2 - 2
		m.height = msg.Height
		m.width = msg.Width
		m.list.SetWidth(msg.Width / 2)
		m.list.SetHeight(msg.Height - 3)

		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(message)
	return m, cmd

}
