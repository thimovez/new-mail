package api

import (
	"fmt"
	"net/http"
	"newMail/models"
	repo "newMail/repo/department"
)

func DepartmentCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	if !r.Form.Has("address") {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: address is required")
		return
	}

	var department models.PostOffice
	department.Address = r.Form.Get("address")

	newDepartment, err := repo.Create(department)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "New Customer : %+v", newDepartment)
}

func DepartmentRead(w http.ResponseWriter, r *http.Request) {
	id, err := readIDFromForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
	}

	department, err := repo.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}
	fmt.Fprintf(w, "New Customer : %+v", department)
}

// DepartmentUpdate - if department change address - use update
func DepartmentUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := readIDFromForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		return
	}

	var department models.PostOffice
	if !r.Form.Has("address") {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: address is required")
		return
	}

	department.ID = id
	department.Address = r.Form.Get("address")
	newDepartment, err := repo.Update(department)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "All is work %+v", newDepartment)
}

func DepartmentDelete(w http.ResponseWriter, r *http.Request) {
	id, err := readIDFromForm(r)
	if err != nil {
		return
	}

	err = repo.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}
}

func DepartmentReadAll(w http.ResponseWriter, r *http.Request) {
	department, err := repo.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "All is work %+v", department)
}
