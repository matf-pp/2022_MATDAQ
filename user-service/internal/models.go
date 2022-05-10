package internal

func GetMoney(username string) (int, error) {
	money, err := Rdb.Get(Ctx, username).Int()
	if err != nil {
		return 0, err
	}
	return money, nil
}

func AddUser(username string, money int32) error {
	return Rdb.Set(Ctx, username, money, 0).Err()
}

func DecreaseMoney(username string, moneyAmount int32) error {
	return Rdb.DecrBy(Ctx, username, int64(moneyAmount)).Err()
}
