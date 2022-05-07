package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"io"
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
	id          int32
}

func (s stock) FilterValue() string {
	return s.Title()
}

func (s stock) Title() string       { return s.title }
func (s stock) Description() string { return s.description }
func (s stock) Id() int32           { return s.id }

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

// retuns empty string because we dont search on OrderType value
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
	securityId      int32
	sideChoice      string
	orderTypeChoice string
	amount          textinput.Model
	price           textinput.Model
	username        [20]byte
	money           int32
	conn            io.Writer
	windowHeight    int
	windowWidth     int
	state           SessionState
	currentWindow   int
}

func New(conn io.Writer) Model {
	// list of stocks, later this will be loaded from some API, or stored in a file
	stocks := []list.Item{
		stock{title: "AAPL", description: "Apple stock", id: 1},
		stock{title: "AMZN", description: "Amazon stock", id: 2},
		stock{title: "MSFT", description: "Microsoft stock", id: 3},
		stock{title: "GOOG", description: "Alphabet Inc. stock", id: 4},
		stock{title: "NFLX", description: "Netflix stock", id: 5},
		stock{title: "TSLA", description: "Tesla stock", id: 6},
		stock{title: "FB", description: "Meta stock", id: 7},
		stock{title: "TSM", description: "Taiwan Semiconductor Manufacturing stock", id: 8},
		stock{title: "NVDA", description: "NVIDIA stock", id: 9},
		stock{title: "V", description: "Visa stock", id: 10},
		stock{title: "WMT", description: "Walmart stock", id: 11},
		stock{title: "JPM", description: "JP Morgan stock", id: 12},
		stock{title: "MA", description: "Mastercard stock", id: 13},
		stock{title: "PEP", description: "PepsiCo  stock", id: 14},
		stock{title: "DIS", description: "Walt Disney stock", id: 15},
		stock{title: "CSCO", description: "Cisco Systems stock", id: 16},
		stock{title: "VZ", description: "Verizon stock", id: 17},
		stock{title: "ORCL", description: "Oracle Corporation stock", id: 18},
		stock{title: "NKE", description: "Nike stock", id: 19},
		stock{title: "ADBE", description: "Adobe stock", id: 20},
		stock{title: "INTC", description: "Intel Corporation stock", id: 21},
	}

	l := list.New(stocks, list.NewDefaultDelegate(), 0, 0)
	l.DisableQuitKeybindings()
	l.Title = "Stocks"
	l.SetShowStatusBar(false)

	// order side
	sides := []list.Item{
		side{title: "Buy", description: "Buy side"},
		side{title: "Sell", description: "Sell side"},
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

	// read data from file where username (and money) is contained and set current user
	username := [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 'i', 'l', 'i', 'j', 'a'}
	var money int32 = 3953211

	return Model{
		list:          l,
		side:          s,
		orderType:     ot,
		amount:        amount,
		price:         price,
		order:         ob,
		username:      username,
		money:         money,
		conn:          conn,
		state:         FocusSelectStock,
		currentWindow: 0,
	}
}
