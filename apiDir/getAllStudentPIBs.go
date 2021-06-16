package apiDir

import "net/http"

func (rest *Rest) getAllStudentPIBs(w http.ResponseWriter, r *http.Request) {
	allPIBs, err := rest.service.Students.GetPIBAllStudents(r.Context())
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, allPIBs)
}
