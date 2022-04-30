package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 7
const listWidth = 20

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Bold(true)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4).Italic(true)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type stock struct {
	title       string
	description string
	sell_price  []float64
	sell_amount []float64
	buy_price   []float64
	buy_amount  []float64
}

func (s stock) FilterValue() string {
	return s.Title()
}

func (s stock) Title() string       { return s.title }
func (s stock) Description() string { return s.description }

type Model struct {
	list        list.Model
	stocks      []stock
	sell_price  []float64
	sell_amount []float64
	buy_price   []float64
	buy_amount  []float64
	choice      string
	height      int
	width       int
}

func New() Model {

	stocks := []list.Item{
		stock{
			title:       "AAPL",
			description: "Apple stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "AMZN",
			description: "Amazon stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "MSFT",
			description: "Microsoft stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "GOOG",
			description: "Alphabet Inc. stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "NFLX",
			description: "Netflix stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "TSLA",
			description: "Tesla stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "FB",
			description: "Meta stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "TSM",
			description: "Taiwan Semiconductor Manufacturing stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "NVDA",
			description: "NVIDIA stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "V",
			description: "Visa stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "WMT",
			description: "Walmart stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "JPM",
			description: "JP Morgan stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "MA",
			description: "Mastercard stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "PEP",
			description: "PepsiCo  stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "DIS",
			description: "Walt Disney stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "CSCO",
			description: "Cisco Systems stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "VZ",
			description: "Verizon stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "ORCL",
			description: "Oracle Corporation stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "NKE",
			description: "Nike stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "ADBE",
			description: "Adobe stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
		stock{
			title:       "INTC",
			description: "Intel Corporation stock",
			sell_price:  []float64{},
			sell_amount: []float64{},
			buy_price:   []float64{},
			buy_amount:  []float64{},
		},
	}

	d := list.NewDefaultDelegate()
	l := list.New(stocks, d, listWidth, listHeight)
	l.Paginator.PerPage = 5
	l.Title = "STOCKS"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return Model{
		list:   l,
		height: 40,
		width:  40,
	}
}
