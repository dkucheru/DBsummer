package apiDir

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (rest *Rest) getTeacherPasses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	passedOrNot := params["pass"]

	teacherStatistics, err := rest.service.Teachers.GetTeacherPassStatistics(r.Context(), passedOrNot)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	rest.sendData(w, teacherStatistics)
}
