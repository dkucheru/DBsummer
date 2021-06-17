package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (rest *Rest) getRatingRunners(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	semester := params["sem"]
	year := params["year"]

	allPIBs, err := rest.service.RunnerMarks.GetRatingStudentWithRunners(r.Context(), semester, year)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, allPIBs)
}
