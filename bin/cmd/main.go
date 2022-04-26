package main

import (
	"net/http"
	"newMail/api"
)

func main() {
	http.HandleFunc("/customer/create", api.CustomerCreate)
	http.HandleFunc("/customer/update", api.CustomerUpdate)
	http.HandleFunc("/customer/read", api.CustomerRead)
	http.HandleFunc("/customer/delete", api.CustomerDelete)
	http.HandleFunc("/customer/readAll", api.CustomerReadAll)

	http.ListenAndServe(":8080", nil)
}
