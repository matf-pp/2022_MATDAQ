package tui

import (
	nos "github.com/matf-pp/2022_MATDAQ/client/pkg/new-order-single"
	"sort"

	"github.com/charmbracelet/bubbles/list"
)

type order struct {
	price    int32
	orderQty uint32
	side     nos.SideValues
}

type stock struct {
	title       string
	description string
	stockId     int32
	buySide     []order
	sellSide    []order
	// will be deleted later
	sellPrices []int32
	sellAmount []uint32
	buyPrices  []int32
	buyAmount  []uint32
}

func (s stock) GetBuySide() ([]int32, []uint32) {
	buyPrices := []int32{}
	buyAmounts := []uint32{}

	sort.Slice(s.buySide, func(i, j int) bool {
		return s.buySide[i].price > s.buySide[j].price
	})

	for _, x := range s.buySide {
		buyPrices = append(buyPrices, x.price)
		buyAmounts = append(buyAmounts, x.orderQty)
	}
	return buyPrices, buyAmounts
}

func (s stock) GetSellSide() ([]int32, []uint32) {
	sellPrices := []int32{}
	sellAmounts := []uint32{}

	sort.Slice(s.sellSide, func(i, j int) bool {
		return s.sellSide[i].price < s.sellSide[j].price
	})

	for _, x := range s.sellSide {
		sellPrices = append(sellPrices, x.price)
		sellAmounts = append(sellAmounts, x.orderQty)
	}
	// reverse order
	for i, j := 0, len(sellPrices)-1; i < j; i, j = i+1, j-1 {
		sellPrices[i], sellPrices[j] = sellPrices[j], sellPrices[i]
		sellAmounts[i], sellAmounts[j] = sellAmounts[j], sellAmounts[i]
	}

	return sellPrices, sellAmounts
}

func (s stock) FilterValue() string {
	return s.Title()
}

func (s stock) Title() string       { return s.title }
func (s stock) Description() string { return s.description }

type Model struct {
	list         list.Model
	stocks       [NUM_OF_STOCKS]stock
	sellPrice    []int32
	sellAmount   []uint32
	buyPrice     []int32
	buyAmount    []uint32
	choice       string
	height       int
	width        int
	windowHeight int
	windowWidth  int
}

func New() Model {
	d := list.NewDefaultDelegate()
	l := list.New(StocksList, d, 0, 0)
	l.Title = "Stocks"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return Model{
		list:   l,
		stocks: StocksArray,
	}
}
