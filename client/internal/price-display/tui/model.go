package tui

import (
	"fmt"
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

func (s stock) GetSellPrices() []int32 {
	sellPrices := []int32{}
	sort.Slice(s.sellSide, func(i, j int) bool {
		return s.sellSide[i].price < s.sellSide[i].price
	})
	for i := 0; i < 10; i++ {
		sellPrices[i] = s.sellSide[i].price
		fmt.Println(sellPrices[i])
	}
	return sellPrices
}

func (s stock) GetBuyPrices() [10]int32 {
	buyPrices := [10]int32{}
	sort.Slice(s.buySide, func(i, j int) bool {
		return s.buySide[i].price < s.buySide[i].price
	})
	for i := 0; i < 10; i++ {
		buyPrices[i] = s.buySide[i].price
		fmt.Println(buyPrices[i])
	}
	return buyPrices
}

func (s stock) GetSellAmount() [10]uint32 {
	sellAmount := [10]uint32{}

	return sellAmount
}

func (s stock) GetBuyAmount() [10]uint32 {
	buyAmount := [10]uint32{}

	return buyAmount
}

func (s stock) FilterValue() string {
	return s.Title()
}

func (s stock) Title() string       { return s.title }
func (s stock) Description() string { return s.description }

type Model struct {
	list         list.Model
	stocks       []stock
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
