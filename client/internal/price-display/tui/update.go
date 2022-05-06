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
			m.sellPrice, m.sellAmount = s.GetSellSide()
			m.buyPrice, m.buyAmount = s.GetBuySide()
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
		/*
			here we will handle Msg's that happen when matching engine sends request about new trade or order
			based on that Msg we will update amount for the given order (if amount == 0 removes it), or add new
			order to the given side of the book

			based on securityId and side we will add new order
			NOTE: this struct order is not the same as the order struct in model.go
			struct order {
				price: int32
				amount: uint32
				securityId: int32
				side: Buy/Sell
			}

			based on securityId and Side we will get the top order from LOB and update its amount
			struct trade {
				amount: uint32
				securityId: int32
				side: Buy/Sell
			}
		*/
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(message)
	return m, cmd

}
