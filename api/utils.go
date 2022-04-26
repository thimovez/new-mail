package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func readIDFromForm(r *http.Request) (id uint64, err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	if !r.Form.Has("id") {
		return 0, fmt.Errorf("not evaluate id param")
	}

	idString := r.Form.Get("id")

	return strconv.ParseUint(idString, 10, 64)
}
