package apiDir

import (
	"net/http"
)

func (rest *Rest) getStudentsList(w http.ResponseWriter, r *http.Request) {
	allstuds, err := rest.service.Students.GetAllStudInfo(r.Context())
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, allstuds)
}
