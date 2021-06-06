package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (rest *Rest) getSheetFromParams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	firstName := params["fn"]
	lastName := params["ln"]
	middleName := params["mn"]
	subj := params["subj"]
	gr := params["gr"]
	year := params["ye"]

	fullSheet, err := rest.service.Sheets.GetSheetFromParameters(r.Context(), firstName, lastName, middleName, subj, gr, year)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, fullSheet)
}
