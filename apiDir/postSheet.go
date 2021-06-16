package apiDir

import (
	"DBsummer/pdfReading"
	"errors"
	"fmt"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	"log"
	"net/http"
	"os"
)

func OutputPdfText(inputPath string) (*string, error) {

	result := ""

	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return nil, err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return nil, err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return nil, err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return nil, err
		}

		result += text
	}

	return &result, nil
}

func (rest *Rest) postSheet(w http.ResponseWriter, r *http.Request) {
	err := license.SetMeteredKey(`10bc0bc50f7829dffd2a7e8b87d38a8ab09775e890aa1a92db4a8f2d70c695a8`)
	if err != nil {
		rest.sendError(w, err)
		return
	}

	//params := mux.Vars(r)
	//path := params["myurl"]
	//fmt.Println(path)
	receivedString, err := OutputPdfText("./collection/Vidomosti_pdf/From4Group_pdfs/infoposhuk_good.pdf")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		//os.Exit(1)
		rest.sendError(w, err)
		return
	}

	doc, err := pdfReading.ParsePDFfile(*receivedString)
	if err != nil {
		log.Println(err)
		rest.sendError(w, err)
		return
	}

	if doc.ControlType == "екзамен" {
		//rest.sendError(w, errors.New("У відомості зазначено слово екзамен, у базу вноситься іспит"))
		_, err = w.Write([]byte("У відомості зазначено слово екзамен, у базу вноситься іспит"))
		if err != nil {
			log.Println(err)
		}
		doc.ControlType = "іспит"
	}

	if doc.Faculty == "" {
		//rest.sendError(w, errors.New("У відомості зазначено слово екзамен, у базу вноситься іспит"))
		_, err = w.Write([]byte("У відомості не зазначено факультет"))
		if err != nil {
			log.Println(err)
		}
		rest.sendError(w, errors.New("перепешіть відомість"))
		return
	}
	//pdf.DebugOn = true
	//content, err := pdfReading.ReadPdf("./collection/phil_new_version.pdf") // Read local pdf file
	//if err != nil {
	//	rest.sendError(w, err)
	//	return
	//}
	//if content == "" {
	//	rest.sendError(w, errors.New("не валідний файл, використовуйте шаблон з сайту"))
	//	return
	//}
	//
	//doc, err := pdfReading.ParsePDFfile(content)
	//if err != nil {
	//	fmt.Println("parser")
	//	fmt.Println(err)
	//	rest.sendError(w, err)
	//	return
	//}

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
