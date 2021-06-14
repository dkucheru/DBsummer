package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (rest *Rest) getSubjectsByYear(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	year := params["year"]

	ye, err := strconv.Atoi(year)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	subjects, err := rest.service.Subjects.FindSubjectsWithYearParameter(r.Context(), ye)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, subjects)
}
