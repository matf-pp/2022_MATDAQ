package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func toString(l []float64) string {
	s := ""
	for i := 0; i < 10 && i < len(l); i++ {
		s1 := fmt.Sprintf("%.2f", l[i])
		s += s1
		s += "\n"
	}
	return s
}

func (m *Model) View() string {

	var style1 = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Height(m.height - 5).
		Width((m.width - 10) / 2)

	var style2 = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Height((m.height - 10) / 2).
		Width((m.width - 20) / 2)

	var priceStyle = lipgloss.NewStyle().
		Bold(true).
		Align(lipgloss.Center).
		Height((m.height - 2) / 10).
		Width((m.width - 10) / 4).
		Render("Price($)")

	var amountStyle = lipgloss.NewStyle().
		Bold(true).
		Align(lipgloss.Center).
		Height((m.height - 2) / 10).
		Width((m.width - 10) / 4).
		Render("Amount(MSFT)")

	var listStyle1 = lipgloss.NewStyle().
		Foreground(lipgloss.Color("150")).
		Align(lipgloss.Center).
		Height((m.height - 2) / 3).
		Width((m.width - 10) / 4).
		Render(toString(m.buy_price))

	var listStyle2 = lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Align(lipgloss.Center).
		Height((m.height - 2) / 3).
		Width((m.width - 10) / 4).
		Render(toString(m.sell_price))

	var amStyle1 = lipgloss.NewStyle().
		Foreground(lipgloss.Color("150")).
		Align(lipgloss.Center).
		Height((m.height - 2) / 3).
		Width((m.width - 10) / 4).
		Render(toString(m.buy_amount))

	var amStyle2 = lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Align(lipgloss.Center).
		Height((m.height - 2) / 3).
		Width((m.width - 10) / 4).
		Render(toString(m.sell_amount))

	leftBox := style1.Render(m.list.View())
	priceBox1 := lipgloss.JoinVertical(lipgloss.Left, priceStyle, listStyle1)
	priceBox2 := lipgloss.JoinVertical(lipgloss.Left, priceStyle, listStyle2)
	amountBox1 := lipgloss.JoinVertical(lipgloss.Left, amountStyle, amStyle1)
	amountBox2 := lipgloss.JoinVertical(lipgloss.Left, amountStyle, amStyle2)
	rightBox1 := style2.Render(lipgloss.JoinHorizontal(lipgloss.Left, priceBox1, amountBox1))
	rightBox2 := style2.Render(lipgloss.JoinHorizontal(lipgloss.Left, priceBox2, amountBox2))
	rightBox := style1.Render(lipgloss.JoinVertical(lipgloss.Left, rightBox1, rightBox2))
	return lipgloss.JoinHorizontal(lipgloss.Left, leftBox, rightBox)
}
