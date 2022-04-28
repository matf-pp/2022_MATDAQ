package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	order_request "github.com/matf-pp/2022_MATDAQ/internal/request-creator/order-request"
)

func (m *Model) blurAllFields() {
	m.list.FilterInput.Blur()
	m.orderType.FilterInput.Blur()
	m.side.FilterInput.Blur()
	m.amount.Blur()
	m.price.Blur()
	m.order.FilterInput.Blur()
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

func (m *Model) resetState() {
	m.blurAllFields()
	m.state = FocusSelectStock
	m.currentWindow = 0
	m.list.ResetSelected()
	m.stockChoice = ""
	m.orderType.ResetSelected()
	m.orderTypeChoice = ""
	m.side.ResetSelected()
	m.sideChoice = ""
	m.amount.Reset()
	m.price.Reset()
	m.list.FilterInput.Focus()
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
		case "enter":
			if m.state == FocusSendOrder {
				// for now we are sending io.Writer to the SendOrder function
				order_request.SendOrder(m.conn, m.orderTypeChoice, m.sideChoice, m.price.Value(), m.amount.Value())
				m.resetState()
			} else if m.state == FocusSelectStock {
				m.switchWindow()
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
