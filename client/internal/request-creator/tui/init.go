package tui

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	api "github.com/matf-pp/2022_MATDAQ/api/user-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type statusMsg int

const PORT int = 9000

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
	var conn *grpc.ClientConn
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("user-service:%d", PORT), opts...)
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
