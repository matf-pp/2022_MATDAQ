package internal

import "github.com/go-redis/redis/v8"

func GetMoney(username string) int {

	val, err := Rdb.Get(Ctx, username).Int()

	if err == redis.Nil {
		return 0
	} else if err != nil {
		//panic(err)
	}

	return val
}

func AddUser(username string, money string) {

	err := Rdb.Set(Ctx, username, money, 0).Err()
	if err != nil {
		//panic(err)
	}

}

func DecreseMoney(username string, money int) {

	val, err := Rdb.Get(Ctx, username).Int()
	if err != nil {
		//panic(err)
	}

	val = val - money

	_ = Rdb.SetXX(Ctx, username, val, 0)
}
