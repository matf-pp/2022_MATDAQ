package main

import (
	"log"
	"net/http"

	"github.com/matf-pp/2022_MATDAQ/user-service/internal"
)

func main() {

	internal.Init_Redis()

	http.HandleFunc("/login", internal.LoginHandler)
	http.HandleFunc("/getmoney", internal.GMHandler)
	http.HandleFunc("/decreasemoney", internal.DMHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
