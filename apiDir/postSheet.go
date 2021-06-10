package apiDir

import (
	"DBsummer/pdfReading"
	"fmt"
	"github.com/ledongthuc/pdf"
	"net/http"
)

func (rest *Rest) postSheet(w http.ResponseWriter, r *http.Request) {

	pdf.DebugOn = true
	content, err := pdfReading.ReadPdf("./collection/phil.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}

	//Finding in DB the teacher from a PDF
	doc := pdfReading.ParsePDFfile(content)
	idTeacher, err := rest.service.Teachers.FindTeacher(r.Context(), doc)
	if err != nil {
		//rest.sendError(w, err)
		idTeacher, err = rest.service.Teachers.AddTeacher(r.Context(), doc)

		if err != nil {
			rest.sendError(w, err)
			return
		}
		//return

	}

	//Adding Sheet instance to DB after checking teacher
	ret, err := rest.service.Sheets.PostSheetToDataBase(r.Context(), doc, *idTeacher)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	for _, v := range doc.ExtractedStudents {
		id, err := rest.service.SheetMarks.PostSheetMarksToDataBase(r.Context(), doc.IdDocument, v)
		if err != nil {
			fmt.Println(id)
			rest.sendError(w, err)
			return
		}
	}

	rest.sendData(w, ret)
}
