package internal

import "github.com/go-redis/redis/v8"

func GetMoney(username string) int {
	money, err := Rdb.Get(Ctx, username).Int()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		panic(err)
	}

	return money
}

func AddUser(username string, money string) {
	err := Rdb.Set(Ctx, username, money, 0).Err()
	if err != nil {
		panic(err)
	}
}

func DecreaseMoney(username string, moneyAmount int) {
	err := Rdb.DecrBy(Ctx, username, int64(moneyAmount)).Err()
	if err != nil {
		panic(err)
	}
}
