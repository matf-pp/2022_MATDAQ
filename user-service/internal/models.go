package internal

import "github.com/go-redis/redis/v8"

func GetMoney(username string) int {

	val, err := Rdb.Get(Ctx, username).Int()

	if err == redis.Nil {
		return -1
	} else if err != nil {
		panic(err)
	}
		
	return val
}

func AddUser(username string, money string)  {

	err := Rdb.Set(Ctx, username, money,0).Err()
	if err != nil {
		panic(err)
	}

}

func DecreseMoney(username string, money int) {

	val1, err1 := Rdb.Get(Ctx, username).Int()
	if err1 != nil {
		panic(err1)
	}

	val1 = val1 - money

	err2 := Rdb.SetXX(Ctx, username, val1,0)
}
