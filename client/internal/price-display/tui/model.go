package tui

import (
	"github.com/charmbracelet/bubbles/list"
	api "github.com/matf-pp/2022_MATDAQ/api/matching-engine"
)

type Order struct {
	price    int32
	orderQty uint32
	side     api.SecurityOrder_OrderSide
}

type Stock struct {
	title       string
	description string
	stockId     uint32
	buySide     []Order
	sellSide    []Order
}

func (s Stock) FilterValue() string { return s.Title() }
func (s Stock) Title() string       { return s.title }
func (s Stock) Description() string { return s.description }

type Model struct {
	stockList        list.Model
	stocks           map[string]Stock
	stockIdIndex     map[uint32]string
	selectedStockKey string

	height       int
	width        int
	windowHeight int
	windowWidth  int
}

func New() Model {
	stockIdIndex := make(map[uint32]string)
	stocks := make(map[string]Stock)
	stockListData := make([]list.Item, 0, len(STOCK_DATA))
	for i, stock := range STOCK_DATA {
		// TODO: this should be set in data
		stock.stockId = uint32(i + 1)

		stockIdIndex[stock.stockId] = stock.title
		stocks[stock.title] = stock
		stockListData = append(stockListData, list.Item(stock))
	}

	stockList := list.New(stockListData, list.NewDefaultDelegate(), 0, 0)
	stockList.Title = "Stocks"
	stockList.SetShowStatusBar(false)
	stockList.SetFilteringEnabled(false)

	return Model{
		stocks:       stocks,
		stockList:    stockList,
		stockIdIndex: stockIdIndex,
	}
}
