package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func intArrayToString(l []int32) string {
	s := ""
	for i := 0; i < 10 && i < len(l); i++ {
		s1 := fmt.Sprintf("%d", l[i])
		s += s1 + "\n"
	}

	return s
}

func uintArrayToString(l []uint32) string {
	s := ""
	for i := 0; i < 10 && i < len(l); i++ {
		s1 := fmt.Sprintf("%d", l[i])
		s += s1 + "\n"
	}

	return s
}

func (m *Model) View() string {

	var (
		header = lipgloss.NewStyle().
			Height(1).
			MarginLeft(1).
			PaddingLeft(1).
			PaddingRight(1).
			Bold(true).
			Background(lipgloss.Color("7")).
			Foreground(lipgloss.Color("#FAFAFA")).
			Render("MATDAQ")

		style1 = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")).
			Height(m.windowHeight).
			Width(m.windowWidth)

		rightBoxSellSideStyle = lipgloss.NewStyle().
					PaddingLeft(1).
					PaddingRight(1).
					BorderStyle(lipgloss.NormalBorder()).
					BorderForeground(lipgloss.Color("9")).
					Height(m.height/2 - 3).
					Width(m.width/2 - 6)

		rightBoxBuySideStyle = lipgloss.NewStyle().
					PaddingLeft(1).
					PaddingRight(1).
					BorderStyle(lipgloss.NormalBorder()).
					BorderForeground(lipgloss.Color("150")).
					Height(m.height/2 - 3).
					Width(m.width/2 - 6)

		priceStyle = lipgloss.NewStyle().
				Bold(true).
				Align(lipgloss.Center).
				Height(2).
				Width((m.width - 10) / 4).
				Render("Price($)")

		amountStyle = lipgloss.NewStyle().
				Bold(true).
				Align(lipgloss.Center).
				Height(2).
				Width((m.width - 10) / 4).
				Render("Amount")

		renderedBuyBoxPrices = lipgloss.NewStyle().
					Foreground(lipgloss.Color("150")).
					Align(lipgloss.Center).
					Height((m.height - 2) / 3).
					Width((m.width - 10) / 4).
					Render(intArrayToString(m.buyPrice))

		renderedBuyBoxAmounts = lipgloss.NewStyle().
					Foreground(lipgloss.Color("150")).
					Align(lipgloss.Center).
					Height((m.height - 2) / 3).
					Width((m.width - 10) / 4).
					Render(uintArrayToString(m.buyAmount))

		renderedSellBoxPrices = lipgloss.NewStyle().
					Foreground(lipgloss.Color("9")).
					Align(lipgloss.Center).
					Height((m.height - 2) / 3).
					Width((m.width - 10) / 4).
					Render(intArrayToString(m.sellPrice))

		renderedSellBoxAmounts = lipgloss.NewStyle().
					Foreground(lipgloss.Color("9")).
					Align(lipgloss.Center).
					Height((m.height - 2) / 3).
					Width((m.width - 10) / 4).
					Render(uintArrayToString(m.sellAmount))
	)

	leftBox := style1.Render(m.list.View())
	buyBoxPrices := lipgloss.JoinVertical(lipgloss.Left, priceStyle, renderedBuyBoxPrices)
	buyBoxAmount := lipgloss.JoinVertical(lipgloss.Left, amountStyle, renderedBuyBoxAmounts)
	sellBoxPrices := lipgloss.JoinVertical(lipgloss.Left, priceStyle, renderedSellBoxPrices)
	sellBoxAmount := lipgloss.JoinVertical(lipgloss.Left, amountStyle, renderedSellBoxAmounts)
	boyBox := rightBoxBuySideStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left, buyBoxPrices, buyBoxAmount))
	sellBox := rightBoxSellSideStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, sellBoxPrices, sellBoxAmount))
	rightBox := style1.Render(lipgloss.JoinVertical(lipgloss.Left, sellBox, boyBox))
	main := lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox)
	return lipgloss.JoinVertical(lipgloss.Left, header, main)
}
