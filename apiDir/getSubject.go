package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (rest *Rest) getSubject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idString := params["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	subject, err := rest.service.Subjects.Get(r.Context(), id)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, subject)
}
