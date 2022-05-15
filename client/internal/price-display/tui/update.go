package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	api "github.com/matf-pp/2022_MATDAQ/api/matching-engine"
	matching_engine "github.com/matf-pp/2022_MATDAQ/client/internal/price-display/matching-engine"
)

func (m *Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			// gets selected stock data
			stockKey := m.stockList.SelectedItem().FilterValue()
			m.selectedStockKey = stockKey
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height - 3
		m.windowWidth = msg.Width/2 - 2
		m.height = msg.Height
		m.width = msg.Width
		m.stockList.SetWidth(msg.Width / 2)
		m.stockList.SetHeight(msg.Height - 3)
		return m, nil
		/*
			here we will handle Msg's that happen when matching engine sends request about new trade or Order
			based on that Msg we will update amount for the given Order (if amount == 0 removes it), or add new
			Order to the given side of the book

			based on securityId and side we will add new Order
			NOTE: this struct Order is not the same as the Order struct in model.go
			struct Order {
				price: int32
				amount: uint32
				securityId: int32
				side: Buy/Sell
			}

			based on securityId and Side we will get the top Order from LOB and update its amount
			struct trade {
				amount: uint32
				securityId: int32
				side: Buy/Sell
			}
		*/
	case *matching_engine.OrderResponse:
		return handleOrderResponse(m, msg)
	case *matching_engine.TradeResponse:
		return handleTradeResponse(m, msg)
	}

	var cmd tea.Cmd
	m.stockList, cmd = m.stockList.Update(message)
	return m, cmd

}

func handleOrderResponse(m *Model, orderResponse *matching_engine.OrderResponse) (*Model, tea.Cmd) {
	security := orderResponse.SecurityOrder

	stockKey := m.stockIdIndex[security.SecurityId]
	stock := m.stocks[stockKey]

	side := security.OrderSide
	price := security.Price
	orderQty := security.OrderQuantity

	order := Order{price, orderQty, side}
	if side == api.SecurityOrder_Buy {
		stock.buySide = append(stock.buySide, order)
	} else {
		stock.sellSide = append(stock.sellSide, order)
	}
	//fmt.Println("stock handle", stock)
	fmt.Println("stock key", stockKey)
	m.stocks[stockKey] = stock
	return m, nil
}

func handleTradeResponse(m *Model, tradeResponse *matching_engine.TradeResponse) (*Model, tea.Cmd) {

	return m, nil
}
