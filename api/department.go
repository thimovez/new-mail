package api

import (
	"fmt"
	"net/http"
	"newMail/models"
	repo "newMail/repo/department"
	"strconv"
)

func DepartmentCreate(w http.ResponseWriter, r *http.Request) {
	department, err := readDepartmentWithoutId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	newDepartment, err := repo.Create(department)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "New Department : %+v", newDepartment)
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
	fmt.Fprintf(w, "Department : %+v", department)
}

// DepartmentUpdate - if department change address - use update
func DepartmentUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := readIDFromForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	department, err := readDepartmentWithoutId(r)
	department.ID = id
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

func DepartmentReadAll(w http.ResponseWriter, _ *http.Request) {
	department, err := repo.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "All is work %+v", department)
}

func readDepartmentWithoutId(r *http.Request) (department models.PostOffice, err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	if !r.Form.Has("address") {
		return department, fmt.Errorf("address is required")
	}
	if !r.Form.Has("maxVolume") {
		return department, fmt.Errorf("max volume is required")
	}
	if !r.Form.Has("maxWeight") {
		return department, fmt.Errorf("max weight is required")
	}

	department.Address = r.Form.Get("address")
	volume, err := strconv.ParseFloat(r.Form.Get("maxVolume"), 10)
	department.MaxVolume = volume
	weight, err := strconv.ParseFloat(r.Form.Get("maxWeight"), 10)
	department.MaxWeight = weight
	return
}
