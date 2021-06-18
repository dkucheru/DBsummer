package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (rest *Rest) getRunnerInfoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idString := params["runnerId"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	subject, err := rest.service.Runners.GetRunnerInfoByID(r.Context(), id)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, subject)
}
