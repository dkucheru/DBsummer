package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (rest *Rest) getStudentMarksByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	idStud, err := strconv.Atoi(id)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	marks, err := rest.service.Students.GetAllStudentMarksByID(r.Context(), idStud)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, marks)
}
