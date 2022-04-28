package api

import (
	"net/http"
)

func Server() {
	http.HandleFunc("/customer/create", CustomerCreate)
	http.HandleFunc("/customer/update", CustomerUpdate)
	http.HandleFunc("/customer/read", CustomerRead)
	http.HandleFunc("/customer/delete", CustomerDelete)
	http.HandleFunc("/customer/readAll", CustomerReadAll)

	http.HandleFunc("/department/create", DepartmentCreate)
	http.HandleFunc("/department/read", DepartmentRead)
	http.HandleFunc("/department/update", DepartmentUpdate)
	http.HandleFunc("/department/delete", DepartmentDelete)
	http.HandleFunc("/department/readAll", DepartmentReadAll)

	http.ListenAndServe(":8080", nil)
}
