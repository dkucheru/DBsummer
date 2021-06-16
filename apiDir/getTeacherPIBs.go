package apiDir

import "net/http"

func (rest *Rest) getTeacherPIBs(w http.ResponseWriter, r *http.Request) {
	allPIBs, err := rest.service.Teachers.GetTeacherPIBs(r.Context())
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, allPIBs)
}
