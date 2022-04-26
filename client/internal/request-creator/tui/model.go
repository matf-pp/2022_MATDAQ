package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type SessionState int

const (
	FocusSelectStock SessionState = iota
	FocusSelectOrderType
	FocusSelectSide
	FocusSelectPrice
	FocusSelectAmount
	FocusSendOrder
)

type stock struct {
	title       string
	description string
}

func (s stock) FilterValue() string {
	return s.Title()
}

func (s stock) Title() string       { return s.title }
func (s stock) Description() string { return s.description }

type side struct {
	title       string
	description string
}

func (s side) FilterValue() string {
	return ""
}

func (s side) Title() string       { return s.title }
func (s side) Description() string { return s.description }

type orderType struct {
	title       string
	description string
}

func (ot orderType) FilterValue() string {
	return ""
}

func (ot orderType) Title() string       { return ot.title }
func (ot orderType) Description() string { return ot.description }

type Model struct {
	list            list.Model
	side            list.Model
	orderType       list.Model
	order           list.Model
	stocks          []stock
	stockChoice     string
	sideChoice      string
	orderTypeChoice string
	amount          textinput.Model
	price           textinput.Model
	windowHeight    int
	windowWidth     int
	state           SessionState
	currentWindow   int
}

func New() Model {
	// TODO: refactor this into multiple files so its not messy
	// list of stocks, later this will be loaded from some API, or stored in a file
	stocks := []list.Item{
		stock{title: "AAPL", description: "Apple stock"},
		stock{title: "AMZN", description: "Amazon stock"},
		stock{title: "MSFT", description: "Microsoft stock"},
		stock{title: "GOOG", description: "Alphabet Inc. stock"},
		stock{title: "NFLX", description: "Netflix stock"},
		stock{title: "TSLA", description: "Tesla stock"},
		stock{title: "FB", description: "Meta stock"},
		stock{title: "TSM", description: "Taiwan Semiconductor Manufacturing stock"},
		stock{title: "NVDA", description: "NVIDIA stock"},
		stock{title: "V", description: "Visa stock"},
		stock{title: "WMT", description: "Walmart stock"},
		stock{title: "JPM", description: "JP Morgan stock"},
		stock{title: "MA", description: "Mastercard stock"},
		stock{title: "PEP", description: "PepsiCo  stock"},
		stock{title: "DIS", description: "Walt Disney stock"},
		stock{title: "CSCO", description: "Cisco Systems stock"},
		stock{title: "VZ", description: "Verizon stock"},
		stock{title: "ORCL", description: "Oracle Corporation stock"},
		stock{title: "NKE", description: "Nike stock"},
		stock{title: "ADBE", description: "Adobe stock"},
		stock{title: "INTC", description: "Intel Corporation stock"},
	}

	l := list.New(stocks, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Stocks"
	l.SetShowStatusBar(false)

	// order side
	// TODO: change styling of But to green and Sell to red
	sides := []list.Item{
		side{title: "Buy", description: "Buy side"},
		stock{title: "Sell", description: "Sell side"},
	}

	s := list.New(sides, list.NewDefaultDelegate(), 0, 0)
	s.Title = "Side"
	s.SetHeight(8)
	s.SetShowHelp(false)
	s.SetFilteringEnabled(false)
	s.SetShowPagination(false)
	s.SetShowStatusBar(false)

	// order type
	orderTypes := []list.Item{
		orderType{title: "Limit Order", description: "Limit order is..."},
		orderType{title: "Market Order", description: "Market order is..."},
	}

	ot := list.New(orderTypes, list.NewDefaultDelegate(), 0, 0)
	ot.Title = "Order Type"
	ot.SetHeight(8)
	ot.SetShowHelp(false)
	ot.SetFilteringEnabled(false)
	ot.SetShowPagination(false)
	ot.SetShowStatusBar(false)

	// order
	orderBtn := []list.Item{
		orderType{title: "Order", description: "Send order request"},
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.Styles.SelectedTitle.Border(lipgloss.NormalBorder(), false, false, false, false).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
		Padding(0, 0, 0, 8)
	ob := list.New(orderBtn, delegate, 0, 0)
	ob.SetShowTitle(false)
	ob.SetShowPagination(false)
	ob.SetShowHelp(false)
	ob.SetShowFilter(false)
	ob.SetShowStatusBar(false)
	ob.ResetSelected()

	// price input field
	price := textinput.New()
	price.Placeholder = "Enter the price"
	price.CharLimit = 20
	price.Prompt = "❯ "

	// amount input field
	amount := textinput.New()
	amount.Placeholder = "Enter the amount"
	amount.CharLimit = 20
	amount.Prompt = "❯ "

	return Model{
		list:          l,
		side:          s,
		orderType:     ot,
		amount:        amount,
		price:         price,
		order:         ob,
		state:         FocusSelectStock,
		currentWindow: 0,
	}
}
