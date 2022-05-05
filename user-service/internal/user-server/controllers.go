package main

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/login" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	username := r.FormValue("username")
	money := r.FormValue("money")

	LoginUser(username, money)

	fmt.Fprintf(w, "LoginHandler %s is a %s\n", username, money)
}

func GMHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/getmoney" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	username := r.FormValue("username")

	money := FindUsername(username)

	w.Write([]byte(money))

	fmt.Fprintf(w, "GMHandler  %s is a %s\n", username, money)
}

func DMHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/decreasemoney" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	username := r.FormValue("username")
	money := r.FormValue("money")

	DecreseMoney(username, money)

	fmt.Fprintf(w, "DMHandler  %s is a %s\n", username, money)
}
