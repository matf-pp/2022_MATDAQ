package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	header := HeaderStyle.Render("MATDAQ")
	amountInput := TextInputStyle.Render(m.amount.View())
	priceInput := TextInputStyle.Render(m.price.View())
	side := m.side.View()
	ordType := m.orderType.View()
	ord := TextInputStyle.Render(m.order.View())

	leftBox := BubbleStyle.Height(m.windowHeight).Width(m.windowWidth).Render(m.list.View())
	rightBox := BubbleStyle.Render(lipgloss.JoinVertical(lipgloss.Left, ordType, side, priceInput, amountInput, ord))
	main := lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox)

	return lipgloss.JoinVertical(lipgloss.Left, header, main)
}
