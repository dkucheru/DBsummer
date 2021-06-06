package apiDir

import "net/http"

func (rest *Rest) getAllBorjniki(w http.ResponseWriter, r *http.Request) {
	allstuds, err := rest.service.Students.GetAllBorjniki(r.Context())
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, allstuds)
}
