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

	doc := pdfReading.ParsePDFfile(content)
	//ret,err := rest.service.Sheets.PostSheetToDataBase(r.Context(),doc)
	//if err != nil {
	//	rest.sendError(w, err)
	//	return
	//}

	for _, v := range doc.ExtractedStudents {
		id, err := rest.service.SheetMarks.PostSheetMarksToDataBase(r.Context(), doc.IdDocument, v)
		if err != nil {
			fmt.Println(id)
			rest.sendError(w, err)
			return
		}
	}

	rest.sendData(w, doc)
}
