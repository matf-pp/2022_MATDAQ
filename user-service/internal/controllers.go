package internal

import (
	"net/http"
	"fmt"
	"strconv"
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitRedis(){
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
	})

}


func LoginHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	money := r.FormValue("money")

	AddUser(username, money)


}

func GMHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")

	money := GetMoney(username)



	fmt.Fprintf(w, "%d\n",money)
}

func DMHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	money, err := strconv.Atoi(r.FormValue("money"))
	if err != nil{
		panic(err)
	}

	DecreseMoney(username, money)

}
