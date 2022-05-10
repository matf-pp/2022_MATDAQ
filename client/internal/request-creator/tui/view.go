package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
)

func (m *Model) View() string {
	moneyAmount := humanize.Comma(int64(m.money))
	headerLogo := HeaderLogoStyle.Render("MATDAQ")
	// adjust width to accommodate for number of digits in money
	headerMoney := HeaderMoneyStyle.MarginLeft(m.windowWidth*2 - 18 - len(moneyAmount)).Render("Balance: $" + moneyAmount)
	header := lipgloss.JoinHorizontal(lipgloss.Bottom, headerLogo, headerMoney)
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
