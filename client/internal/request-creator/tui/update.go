package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) blurAllFields() {
	m.list.FilterInput.Blur()
	m.orderType.FilterInput.Blur()
	m.side.FilterInput.Blur()
	m.amount.Blur()
	m.price.Blur()
}

func (m *Model) switchWindow() {
	m.currentWindow = 1
	m.blurAllFields()
	m.orderType.FilterInput.Focus()
	m.state = FocusSelectOrderType
}

// TODO: implement toggling between different fields in the app
func (m *Model) nextField(forward bool) {
	if m.currentWindow == 0 ||
		forward && m.state == FocusSendOrder ||
		!forward && m.state == FocusSelectOrderType {
		return
	} else if forward {
		m.state += 1
	} else {
		m.state -= 1
	}

	switch m.state {
	case FocusSelectStock:
		m.blurAllFields()
		m.list.FilterInput.Focus()
	case FocusSelectOrderType:
		m.blurAllFields()
		m.orderType.FilterInput.Focus()
	case FocusSelectSide:
		m.blurAllFields()
		m.side.FilterInput.Focus()
	case FocusSelectPrice:
		m.blurAllFields()
		m.price.Focus()
	case FocusSelectAmount:
		m.blurAllFields()
		m.amount.Focus()
	case FocusSendOrder:
		m.blurAllFields()
		m.order.FilterInput.Focus()
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height - 3
		m.windowWidth = msg.Width/2 - 2
		m.list.SetWidth(msg.Width / 2)
		m.list.SetHeight(msg.Height - 3)
		m.side.SetWidth(msg.Width / 2)
		m.orderType.SetWidth(msg.Width / 2)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		// TODO: if we press enter and current field is amount we should send an order
		// implement better logic for moving through fields
		case "enter":
			if m.state == FocusSendOrder {
				// IMPLEMENT: SEND ORDER AND RESET FIELDS
				// here we take all inputs and send them to order-request-service
			} else if m.state == FocusSelectStock {
				m.switchWindow()
				// we dont need to do this here actually
				s, ok := m.list.SelectedItem().(stock)
				// we would need to use stock ID instead of title
				if ok {
					m.stockChoice = string(s.title)
				}
			}
		case "tab":
			m.nextField(true)
		case "shift+tab":
			m.nextField(false)
		}
	}

	// TODO: clean this up
	if m.list.FilterInput.Focused() {
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.orderType.FilterInput.Focused() {
		m.orderType, cmd = m.orderType.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.side.FilterInput.Focused() {
		m.side, cmd = m.side.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.amount, cmd = m.amount.Update(msg)
	cmds = append(cmds, cmd)

	m.price, cmd = m.price.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
