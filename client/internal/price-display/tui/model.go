package tui

import (
	"github.com/charmbracelet/bubbles/list"
)

type stock struct {
	title       string
	description string
	sell_price  []int32
	sell_amount []uint32
	buy_price   []int32
	buy_amount  []uint32
}

func (s stock) FilterValue() string {
	return s.Title()
}

func (s stock) Title() string       { return s.title }
func (s stock) Description() string { return s.description }

type Model struct {
	list         list.Model
	stocks       []stock
	sell_price   []int32
	sell_amount  []uint32
	buy_price    []int32
	buy_amount   []uint32
	choice       string
	height       int
	width        int
	windowHeight int
	windowWidth  int
}

func New() Model {

	stocks := []list.Item{
		stock{
			title:       "AAPL",
			description: "Apple stock",
			sell_price:  []int32{100, 120, 130, 140, 1000, 10001, 100, 200, 30, 300},
			sell_amount: []uint32{100, 120, 130, 140, 1000, 10001, 100, 200, 30, 300},
			buy_price:   []int32{100, 120, 130, 140, 1000, 10001, 100, 200, 30, 300},
			buy_amount:  []uint32{100, 120, 130, 140, 1000, 10001, 100, 200, 30, 300},
		},
		stock{
			title:       "AMZN",
			description: "Amazon stock",
			sell_price:  []int32{130, 140, 1000, 10001, 100, 200, 30, 300},
			sell_amount: []uint32{130, 140, 1000, 10001, 100, 200, 30, 300},
			buy_price:   []int32{130, 140, 1000, 10001, 100, 200, 30, 300},
			buy_amount:  []uint32{130, 140, 1000, 10001, 100, 200, 30, 300},
		},
		stock{
			title:       "MSFT",
			description: "Microsoft stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "GOOG",
			description: "Alphabet Inc. stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "NFLX",
			description: "Netflix stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "TSLA",
			description: "Tesla stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "FB",
			description: "Meta stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "TSM",
			description: "Taiwan Semiconductor Manufacturing stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "NVDA",
			description: "NVIDIA stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "V",
			description: "Visa stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "WMT",
			description: "Walmart stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "JPM",
			description: "JP Morgan stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "MA",
			description: "Mastercard stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "PEP",
			description: "PepsiCo  stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "DIS",
			description: "Walt Disney stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "CSCO",
			description: "Cisco Systems stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "VZ",
			description: "Verizon stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "ORCL",
			description: "Oracle Corporation stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "NKE",
			description: "Nike stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "ADBE",
			description: "Adobe stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
		stock{
			title:       "INTC",
			description: "Intel Corporation stock",
			sell_price:  []int32{},
			sell_amount: []uint32{},
			buy_price:   []int32{},
			buy_amount:  []uint32{},
		},
	}

	d := list.NewDefaultDelegate()
	l := list.New(stocks, d, 0, 0)
	l.Title = "Stocks"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return Model{
		list: l,
	}
}
