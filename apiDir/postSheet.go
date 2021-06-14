package apiDir

import (
	"DBsummer/pdfReading"
	"errors"
	"fmt"
	"github.com/ledongthuc/pdf"
	"net/http"
)

func (rest *Rest) postSheet(w http.ResponseWriter, r *http.Request) {

	pdf.DebugOn = true
	content, err := pdfReading.ReadPdf("./collection/phil_new_version.pdf") // Read local pdf file
	if err != nil {
		rest.sendError(w, err)
		return
	}
	if content == "" {
		rest.sendError(w, errors.New("не валідний файл, використовуйте шаблон з сайту"))
		return
	}

	doc, err := pdfReading.ParsePDFfile(content)
	if err != nil {
		fmt.Println("parser")
		fmt.Println(err)
		rest.sendError(w, err)
		return
	}

	//Finding in DB the teacher from a PDF
	idTeacher, err := rest.service.Teachers.FindTeacher(r.Context(), doc)
	if err != nil {
		//If select SQL query returned error, it means that teacher does not exist in the DB
		//Therefore we need to add him
		idTeacher, err = rest.service.Teachers.AddTeacher(r.Context(), doc)

		if err != nil {
			fmt.Println("error sql-adding teacher")
			fmt.Println(err)
			rest.sendError(w, err)
			return
		}
		//now,when teacher was added we can get his id
		idTeacher, err = rest.service.Teachers.FindTeacher(r.Context(), doc)
		if err != nil {
			rest.sendError(w, err)
			return
		}
	}

	//Finding in DB the subject from a PDF
	idSubject, err := rest.service.Subjects.FindSubject(r.Context(), doc)
	if err != nil {
		//If select SQL query returned error, it means that subject does not exist in the DB
		//Therefore we need to add it
		idSubject, err = rest.service.Subjects.AddSubject(r.Context(), doc)

		if err != nil {
			fmt.Println("error adding subject")
			fmt.Println(err)
			rest.sendError(w, err)
			return
		}

		//now,when subject was added we can get his id
		idSubject, err = rest.service.Subjects.FindSubject(r.Context(), doc)
		if err != nil {
			rest.sendError(w, err)
			fmt.Println("error finding subject")
			fmt.Println(err)
			return
		}
	}

	//Finding in DB the group from a PDF
	idGroup, err := rest.service.Groups.FindGroup(r.Context(), doc, *idSubject)
	if err != nil {
		//If select SQL query returned error, it means that group does not exist in the DB
		//Therefore we need to add it
		idGroup, err = rest.service.Groups.AddGroup(r.Context(), doc, *idSubject)

		if err != nil {
			fmt.Println("error adding group")
			fmt.Println(err)
			rest.sendError(w, err)
			return
		}

		//now,when group was added we can get his id
		idGroup, err = rest.service.Groups.FindGroup(r.Context(), doc, *idSubject)
		if err != nil {
			rest.sendError(w, err)
			return
		}
	}

	//Adding Sheet instance to DB after checking teacher,subject and group
	ret, err := rest.service.Sheets.PostSheetToDataBase(r.Context(), doc, *idTeacher, *idGroup)
	if err != nil {
		fmt.Println("error adding sheet")
		fmt.Println(err)
		rest.sendError(w, err)
		return
	}

	for _, v := range doc.ExtractedStudents {
		idStudent, err := rest.service.Students.FindStudent(r.Context(), v)
		if err != nil {
			//If select SQL query returned error, it means that group does not exist in the DB
			//Therefore we need to add it
			idStudent, err = rest.service.Students.AddStudent(r.Context(), v)
			if err != nil {
				fmt.Println("error adding student")
				fmt.Println(err)
				rest.sendError(w, err)
				return
			}

			//now,when group was added we can get his id
			idStudent, err = rest.service.Students.FindStudent(r.Context(), v)
			if err != nil {
				rest.sendError(w, err)
				return
			}
		}

		id, err := rest.service.SheetMarks.PostSheetMarksToDataBase(r.Context(), doc.IdDocument, *idStudent, v)
		if err != nil {
			fmt.Println("error adding sheetmarks")
			fmt.Println(id)
			rest.sendError(w, err)
			return
		}
	}

	rest.sendData(w, ret)
}
