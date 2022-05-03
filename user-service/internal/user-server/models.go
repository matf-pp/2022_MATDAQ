package main

import "github.com/go-redis/redis/v8"

func FindUsername(string username) int {

	val, err := rdb.Get(ctx, username).Int()

	if err == redis.Nil {
		return -1
	} else if err != nil {
		panic(err)
	} else {
		return val
	}
}

func LoginUser(string) {

	_, err := rdb.Set(ctx, username, val1)
	if err != nill {
		panic(err)
	}

}

func DecreseMoney(string username, int money) {

	val1, err := rdb.Get(ctx, username).Int()
	if err != nil {
		panic(err)
	}

	val1 = val1 - money

	val2, err := rdb.SetXX(ctx, username, val1)
}
