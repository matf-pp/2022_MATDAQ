package internal

import (
	"fmt"
	"net/http"
	"strconv"
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

var Ctx Contex
var Rdb *redis.Client

func Init_Redis(){
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
	})

	Ctx = context.Background()
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/login" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	username := r.FormValue("username")
	money := r.FormValue("money")

	AddUser(username, money)

	// fmt.Fprintf(w, "LoginHandler %s is a %s\n", username, money)
}

func GMHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/getmoney" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	username := r.FormValue("username")

	money := GetMoney(username)



	// fmt.Fprintf(w, "GMHandler  %s is a %d\n", username, money)
}

func DMHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/decreasemoney" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	username := r.FormValue("username")
	money, err := strconv.Atoi(r.FormValue("money"))
	if err != nil{
		panic(err)
	}

	DecreseMoney(username, money)

	// fmt.Fprintf(w, "DMHandler  %s is a %s\n", username, money)
}
