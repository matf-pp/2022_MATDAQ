package tui

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type statusMsg int

type errMsg struct{ error }

const HOST = "http://localhost:8080"

func (m *Model) Init() tea.Cmd {
	return func() tea.Msg { return checkServer(m.username, m.money) }
}

func toStringUsername(username [20]byte) string {
	parsedUsername := ""

	for _, c := range username {
		if c == 0 {
			break
		}
		parsedUsername += string(c)
	}

	return parsedUsername
}

func checkServer(username [20]byte, money int32) tea.Msg {
	url := HOST + "/login?username=" + toStringUsername(username) + "&money=" + strconv.Itoa(int(money))
	fmt.Println(url)
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := c.Get(url)
	if err != nil {
		fmt.Println("Error happened: {}", err)
		return errMsg{err}
	}

	fmt.Println("{}", res)
	return statusMsg(res.StatusCode)
}
