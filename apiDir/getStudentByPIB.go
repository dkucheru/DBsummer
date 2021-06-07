package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (rest *Rest) getStudentByPIB(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	firstName := params["fn"]
	lastName := params["ln"]
	secondName := params["mn"]
	year := params["ye"]

	allstuds, err := rest.service.Students.GetStudentByPIB(r.Context(), firstName, lastName, secondName, year)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, allstuds)
}
