package postgres

import (
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
	"fmt"
	"log"
)

type SheetsRepository struct {
	*Repository
}

func (r SheetsRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r SheetsRepository) Get(ctx context.Context) error {
	panic("implement me")
}

func (r SheetsRepository) PostSheetToDataBase(ctx context.Context, sheet *pdfReading.ExtractedInformation, teacherId int) (*pdfReading.ExtractedInformation, error) {
	//	getTeacherCipher := r.db.Rebind(`
	//		SELECT teacher_cipher
	//		FROM teachers
	//		WHERE firstname = ? AND lastname = ? AND (middlename = ? OR middlename IS NULL OR middlename = '')
	//;
	//	`)
	//	row := r.db.QueryRowContext(ctx, getTeacherCipher, sheet.TeacherFirstName, sheet.TeacherLastname, sheet.TeacherMiddleName)
	//	var teacherID string
	//	err := row.Scan(&teacherID)
	//	if err != nil {
	//		return sheet, err
	//	}

	getGroupCipher := r.db.Rebind(`
		SELECT cipher
		FROM groups_
		WHERE groupname = ? AND educationalyear = ? AND semester = ?;
	`)
	row := r.db.QueryRowContext(ctx, getGroupCipher, sheet.GroupName, sheet.EducationalYear, sheet.Semester)
	//fmt.Println("group name : " + sheet.GroupName)
	//fmt.Println("year : " + sheet.EducationalYear)
	//fmt.Println("semester : "+ sheet.Semester)
	var groupID string
	err := row.Scan(&groupID)
	if err != nil {
		log.Println(err)
		return sheet, err
	}

	query := r.db.Rebind(`
		INSERT into sheet(sheetid,number_of_attendees,number_of_absent,number_of_ineligible,
	type_of_control,date_of_compilation,teacher,group_cipher) VALUES(?,?,?,?,?,?,?,?);
		`)

	_, err = r.db.Exec(query, sheet.IdDocument, sheet.AmountPresent, sheet.AmountAbscent,
		sheet.AmountBanned, sheet.ControlType, sheet.Date, teacherId, groupID)
	if err != nil {
		log.Println(err)
		if err.Error() == "pq: повторяющееся значение ключа нарушает ограничение уникальности \"sheet_pkey\"" {
			log.Println("Error defined")
		}
		return sheet, err
	}

	return sheet, nil
}

func (r SheetsRepository) GetSheetFromParameters(ctx context.Context, fn string, ln string, mn string, subj string, gr string, year string) ([]*structs.SheetByQuery, error) {
	query := r.db.Rebind(`
		SELECT sheetid, 
		student.firstname || ' ' || last_name || ' ' || COALESCE(middle_name, '') AS pibStud,
		semester_mark, COALESCE(check_mark, 0), COALESCE(together_mark, 0),
		national_mark,COALESCE(ects_mark, 'Undefined')
FROM ((((student JOIN sheet_marks ON student_cipher = student ) 
	  JOIN sheet ON sheet = sheet.sheetid) 
	  JOIN groups_ ON group_cipher = cipher) 
	  JOIN subjects ON subject = subjectid)
	  JOIN teachers ON teacher_cipher = sheet.teacher
WHERE lastname = ? AND teachers.firstname = ? AND (middlename = ? OR middlename IS NULL) 
		AND subjectname = ? AND educationalyear = ? AND groupname = ?;`)

	rows, err := r.db.QueryContext(ctx, query, ln, fn, mn, subj, year, gr)
	if err != nil {
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.Println(e)
		}
	}()

	var allInfo []*structs.SheetByQuery
	for rows.Next() {
		var s structs.SheetByQuery
		err = rows.Scan(&s.SheetId, &s.PibStudent, &s.SemesterMark, &s.CheckMark, &s.TogetherMark, &s.NationalMark, &s.EctsMark)
		if err != nil {
			return nil, err
		}
		fmt.Print("student.go file inside loop : ")
		fmt.Println(&s)
		allInfo = append(allInfo, &s)
	}
	fmt.Print("student.go after loop: ")
	fmt.Println(&allInfo)
	return allInfo, nil
}
