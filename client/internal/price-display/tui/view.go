package tui

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/lipgloss"
)

func orderToString(orders []Order) (string, string) {
	prices := ""
	amount := ""
	for i := 0; i < 10 && i < len(orders); i++ {
		prices += "" + fmt.Sprintf("%d", orders[i].price) + "\n"
		amount += "" + fmt.Sprintf("%d", orders[i].orderQty) + "\n"
	}
	return prices, amount
}

func (m *Model) View() string {
	s := m.stocks[m.selectedStockKey]
	sort.Slice(s.buySide, func(i, j int) bool {
		return s.buySide[i].price > s.buySide[j].price
	})
	buyPrices, buyAmount := orderToString(s.buySide)
	sort.Slice(s.sellSide, func(i, j int) bool {
		return s.sellSide[i].price > s.sellSide[j].price
	})
	sellPrices, sellAmount := orderToString(s.sellSide)
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

		style = lipgloss.NewStyle().
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
					Render(buyPrices)

		renderedBuyBoxAmounts = lipgloss.NewStyle().
					Foreground(lipgloss.Color("150")).
					Align(lipgloss.Center).
					Height((m.height - 2) / 3).
					Width((m.width - 10) / 4).
					Render(buyAmount)

		renderedSellBoxPrices = lipgloss.NewStyle().
					Foreground(lipgloss.Color("9")).
					Align(lipgloss.Center).
					Height((m.height - 2) / 3).
					Width((m.width - 10) / 4).
					Render(sellPrices)

		renderedSellBoxAmounts = lipgloss.NewStyle().
					Foreground(lipgloss.Color("9")).
					Align(lipgloss.Center).
					Height((m.height - 2) / 3).
					Width((m.width - 10) / 4).
					Render(sellAmount)
	)

	leftBox := style.Render(m.stockList.View())
	buyBoxPrices := lipgloss.JoinVertical(lipgloss.Left, priceStyle, renderedBuyBoxPrices)
	buyBoxAmount := lipgloss.JoinVertical(lipgloss.Left, amountStyle, renderedBuyBoxAmounts)
	sellBoxPrices := lipgloss.JoinVertical(lipgloss.Left, priceStyle, renderedSellBoxPrices)
	sellBoxAmount := lipgloss.JoinVertical(lipgloss.Left, amountStyle, renderedSellBoxAmounts)
	boyBox := rightBoxBuySideStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left, buyBoxPrices, buyBoxAmount))
	sellBox := rightBoxSellSideStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, sellBoxPrices, sellBoxAmount))
	rightBox := style.Render(lipgloss.JoinVertical(lipgloss.Left, sellBox, boyBox))
	main := lipgloss.JoinHorizontal(lipgloss.Top, leftBox, rightBox)
	return lipgloss.JoinVertical(lipgloss.Left, header, main)
}
