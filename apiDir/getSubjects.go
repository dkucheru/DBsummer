package apiDir

import (
	"net/http"
)

func (rest *Rest) getSubjects(w http.ResponseWriter, r *http.Request) {
	subject, err := rest.service.Subjects.GetAll(r.Context())
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, subject)
}
