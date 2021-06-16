package apiDir

import "net/http"

func (rest *Rest) deleteAllData(w http.ResponseWriter, r *http.Request) {
	err := rest.service.Sheets.DeleteAllData(r.Context())
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, "DB data deleted successfully")
}
