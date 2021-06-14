package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (rest *Rest) getGroupsOfScientist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	degree := params["degree"]

	groups, err := rest.service.Groups.FindGroupsOfScientist(r.Context(), degree)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, groups)
}
