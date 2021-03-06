package api

import (
	"fmt"
	"net/http"
	"newMail/models"
	repo "newMail/repo/parcel"
	"strconv"
)

func ParcelCreate(w http.ResponseWriter, r *http.Request) {
	parcel, err := readParcelWithoutID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	newParcel, err := repo.Create(parcel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "Parcel : %+v", newParcel)
}

func ParcelRead(w http.ResponseWriter, r *http.Request) {
	id, err := readIDFromForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
	}

	parcel, err := repo.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "All is work %+v", parcel)

}

func ParcelUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := readIDFromForm(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	parcel, err := readParcelWithoutID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	parcel.ID = id
	newParcel, err := repo.Update(parcel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "All is work %+v", newParcel)
}

func ParcelDelete(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "All is work %+v")
}

func ParcelReadAll(w http.ResponseWriter, _ *http.Request) {
	result, err := repo.ReadAll() // add limit and offset
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "API error: %s", err)
		return
	}

	fmt.Fprintf(w, "All is work %+v", result)
}

func readParcelWithoutID(r *http.Request) (parcel models.Parcel, err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	if !r.Form.Has("name") {
		return parcel, fmt.Errorf("name is required")
	}
	if !r.Form.Has("weight") {
		return parcel, fmt.Errorf("weight is required")
	}
	if !r.Form.Has("volume") {
		return parcel, fmt.Errorf("volume is required")
	}
	if !r.Form.Has("price") {
		return parcel, fmt.Errorf("price is required")
	}
	if !r.Form.Has("sender") {
		return parcel, fmt.Errorf("sender is required")
	}
	if !r.Form.Has("receiver") {
		return parcel, fmt.Errorf("receiver is required")
	}
	if !r.Form.Has("sourceDep") {
		return parcel, fmt.Errorf("source department is required")
	}
	if !r.Form.Has("destinationDep") {
		return parcel, fmt.Errorf("destination department is required")
	}

	weight, err := strconv.ParseFloat(r.Form.Get("weight"), 10)
	if err != nil {
		return parcel, fmt.Errorf("weight must be number")
	}
	volume, err := strconv.ParseFloat(r.Form.Get("volume"), 10)
	if err != nil {
		return parcel, fmt.Errorf("volume must be number")
	}
	price, err := strconv.ParseFloat(r.Form.Get("price"), 10)
	if err != nil {
		return parcel, fmt.Errorf("price must be number")
	}
	sender, err := strconv.ParseUint(r.Form.Get("sender"), 10, 64)
	if err != nil {
		return parcel, fmt.Errorf("sender must be number")
	}
	receiver, err := strconv.ParseUint(r.Form.Get("receiver"), 10, 64)
	if err != nil {
		return parcel, fmt.Errorf("receiver must be number")
	}
	sourceDep, err := strconv.ParseUint(r.Form.Get("sourceDep"), 10, 64)
	if err != nil {
		return parcel, fmt.Errorf("sourceDep must be number")
	}
	destinationDep, err := strconv.ParseUint(r.Form.Get("destinationDep"), 10, 64)
	if err != nil {
		return parcel, fmt.Errorf("destinationDep must be number")
	}

	parcel.Name = r.Form.Get("name")
	parcel.Weight = weight
	parcel.Volume = volume
	parcel.Price = price
	parcel.SenderID = sender
	parcel.ReceiverID = receiver
	parcel.SourceDepartmentID = sourceDep
	parcel.DestinationDepartmentID = destinationDep
	return parcel, nil
}
