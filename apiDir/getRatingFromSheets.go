package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (rest *Rest) getRatingSheets(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	semester := params["sem"]
	year := params["year"]

	allPIBs, err := rest.service.SheetMarks.GetRatingStudents(r.Context(), semester, year)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, allPIBs)
}
