package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (rest *Rest) getRunnerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idString := params["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	runner, err := rest.service.Runners.GetRunnerByID(r.Context(), id)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, runner)
}
