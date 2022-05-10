package tui

import (
	"context"
	tea "github.com/charmbracelet/bubbletea"
	api "github.com/matf-pp/2022_MATDAQ/api/user-service"
	"google.golang.org/grpc"
	"log"
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
	// TODO: move connection creation out of checkServer
	var conn *grpc.ClientConn
	var opts []grpc.DialOption
	conn, err := grpc.Dial(":9000", opts...)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	userClient := api.NewUserClient(conn)
	loginUserReq := &api.LoginUserRequest{
		Username: toStringUsername(username),
		Money:    money,
	}
	_, err = userClient.LoginUser(context.Background(), loginUserReq)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	return statusMsg(200)
}
