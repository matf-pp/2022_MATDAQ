package tui

import "github.com/charmbracelet/lipgloss"

func (m *Model) View() string {
	var style1 = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(2).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).Height(m.height - 2).Width((m.width - 10) / 2)

	var priceStyle = lipgloss.NewStyle().Padding(0, 0, 0, 0).
		BorderStyle(lipgloss.NormalBorder()).Bold(true).Align(lipgloss.Center).
		BorderForeground(lipgloss.Color("63")).Height((m.height - 2) / 8).Width((m.width - 2) / 4).Render("Price($)")

	var amountStyle = lipgloss.NewStyle().Padding(0, 0, 0, 0).
		BorderStyle(lipgloss.NormalBorder()).Bold(true).Align(lipgloss.Center).
		BorderForeground(lipgloss.Color("63")).Height((m.height - 2) / 8).Width((m.width - 2) / 4).Render("Amount(MSFT)")

	var listStyle1 = lipgloss.NewStyle().Padding(0, 0, 0, 0).Foreground(lipgloss.Color("150")).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63")).Align(lipgloss.Center).Height((m.height - 2) / 4).Width((m.width - 2) / 4).Render("180\n179\n178\n177\n176\n")
	var listStyle2 = lipgloss.NewStyle().Padding(0, 0, 0, 0).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63")).Align(lipgloss.Center).Height((m.height - 2) / 4).Width((m.width - 2) / 4).Render("182\n179\n178\n177\n173\n")
	var amStyle1 = lipgloss.NewStyle().Padding(0, 0, 0, 0).Foreground(lipgloss.Color("150")).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63")).Align(lipgloss.Center).Height((m.height - 2) / 4).Width((m.width - 2) / 4).Render("500\n400\n600\n700\n820\n")
	var amStyle2 = lipgloss.NewStyle().Padding(0, 0, 0, 0).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63")).Align(lipgloss.Center).Height((m.height - 2) / 4).Width((m.width - 2) / 4).Render("300\n450\n1000\n200\n488\n")

	leftBox := style1.Render(m.list.View())
	priceBox1 := lipgloss.JoinVertical(lipgloss.Top, priceStyle, listStyle1)
	priceBox2 := lipgloss.JoinVertical(lipgloss.Top, priceStyle, listStyle2)
	amountBox1 := lipgloss.JoinVertical(lipgloss.Top, amountStyle, amStyle1)
	amountBox2 := lipgloss.JoinVertical(lipgloss.Top, amountStyle, amStyle2)
	rightBox1 := lipgloss.JoinHorizontal(lipgloss.Left, priceBox1, amountBox1)
	rightBox2 := lipgloss.JoinHorizontal(lipgloss.Left, priceBox2, amountBox2)

	return lipgloss.JoinHorizontal(lipgloss.Left, leftBox, lipgloss.JoinVertical(lipgloss.Top, rightBox1, rightBox2))
}
