package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(stock)
			if ok {
				m.sell_price = i.sell_price
				m.sell_amount = i.sell_amount
				m.buy_price = i.buy_price
				m.buy_amount = i.buy_amount
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.list.SetWidth(msg.Width / 2)
		m.list.SetHeight(msg.Height - 2)
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(message)
	return m, cmd

}
