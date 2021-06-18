package apiDir

import (
	"net/http"
)

func (rest *Rest) getAllStudentPIBs(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()["q"][0]

	allPIBs, err := rest.service.Students.GetPIBAllStudents(r.Context(), q)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, allPIBs)
}
