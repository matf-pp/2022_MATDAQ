package internal

import (
	"context"
	"github.com/go-redis/redis/v8"
	api "github.com/matf-pp/2022_MATDAQ/api/user-service"
	"log"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

type UserServerImpl struct {
	api.UnimplementedUserServer
}

func NewUserServer() *UserServerImpl {
	return &UserServerImpl{}
}

func (s *UserServerImpl) LoginUser(ctx context.Context, req *api.LoginUserRequest) (*api.LoginUserResponse, error) {
	log.Println("Login user: ", req)
	if err := AddUser(req.Username, req.Money); err != nil {
		return nil, err
	}
	return &api.LoginUserResponse{}, nil
}

func (s *UserServerImpl) DecreaseMoney(ctx context.Context, req *api.DecreaseMoneyRequest) (*api.DecreaseMoneyResponse, error) {
	log.Println("Decrease money: ", req)
	if err := DecreaseMoney(req.Username, req.MoneyAmount); err != nil {
		return nil, err
	}
	return &api.DecreaseMoneyResponse{}, nil
}
