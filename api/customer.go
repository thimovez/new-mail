package api

import (
	"fmt"
	"net/http"
	"newMail/models"
	repo "newMail/repo/customer"
	"strconv"
)

func CustomerCreate(w http.ResponseWriter, r *http.Request) {
	customer, err := readCustomer(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	newCustomer, err := repo.Create(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "New Customer : %+v", newCustomer)
}

func CustomerRead(w http.ResponseWriter, r *http.Request) {
	id, err := readIDFromForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
	}

	customer, err := repo.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "All is work %+v", customer)
}

func CustomerUpdate(w http.ResponseWriter, r *http.Request) {
	customer, err := readCustomer(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	newCustomer, err := repo.Update(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "All is work %+v", newCustomer)
}

func CustomerDelete(w http.ResponseWriter, r *http.Request) {
	id, err := readIDFromForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
	}

	err = repo.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}
	fmt.Fprintf(w, "All is work")
}

func CustomerReadAll(w http.ResponseWriter, _ *http.Request) {
	result, err := repo.ReadAll() // add limit and offset
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "All is work %+v", result)
}

func readCustomer(r *http.Request) (customer models.Customer, err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	if !r.Form.Has("id") {
		return customer, fmt.Errorf("id is required")
	}
	if !r.Form.Has("name") {
		return customer, fmt.Errorf("name is required")
	}
	if !r.Form.Has("surname") {
		return customer, fmt.Errorf("surname is required")
	}

	id, err := strconv.ParseUint(r.Form.Get("id"), 10, 64)
	customer.ID = id
	customer.Name = r.Form.Get("name")
	customer.SurName = r.Form.Get("surname")
	return
}
