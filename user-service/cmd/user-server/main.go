package main

import (
	"log"
	"net/http"

	"github.com/matf-pp/2022_MATDAQ/user-service/internal"
)

func main() {

	internal.InitRedis()

	http.HandleFunc("/login", internal.LoginHandler)
	http.HandleFunc("/getMoney", internal.GMHandler)
	http.HandleFunc("/decreaseMoney", internal.DMHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
