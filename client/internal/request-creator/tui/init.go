package tui

import (
	"net/http"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type statusMsg int

type errMsg struct{ error }

var username = ""
var money = ""

func (m *Model) Init() tea.Cmd {
	username = string(m.username[:])
	money = strconv.Itoa(int(m.money))

	return checkServer
}

func checkServer() tea.Msg {
	url := "http://localhost:8080/login?username=" + username + "&money=" + money
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}

	return statusMsg(res.StatusCode)
}
