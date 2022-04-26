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
}

func (s stock) FilterValue() string {
	return s.Title()
}

func (s stock) Title() string       { return s.title }
func (s stock) Description() string { return s.description }

type Model struct {
	list        list.Model
	stocks      []stock
	sell_prices list.Model
	buy_prices  list.Model
	choice      string
	height      int
	width       int
}

func New() Model {

	stocks := []list.Item{
		stock{title: "AAPL", description: "Apple stock"},
		stock{title: "MSFT", description: "Microsoft stock"},
		stock{title: "NFLX", description: "Netflix stock"},
		stock{title: "TSLA", description: "Tesla stock"},
		stock{title: "CKAM", description: "Tesla stock"},
		stock{title: "TTKA", description: "Tesla stock"},
		stock{title: "OPRM", description: "Tesla stock"},
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
