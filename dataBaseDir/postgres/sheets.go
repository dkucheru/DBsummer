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

func (r SheetsRepository) PostSheetToDataBase(ctx context.Context, sheet *pdfReading.ExtractedInformation, teacherId int, groupId int) (*pdfReading.ExtractedInformation, error) {

	query := r.db.Rebind(`
		INSERT into sheet(sheetid,number_of_attendees,number_of_absent,number_of_ineligible,
	type_of_control,date_of_compilation,teacher,group_cipher) VALUES(?,?,?,?,?,?,?,?);
		`)

	_, err := r.db.Exec(query, sheet.IdDocument, sheet.AmountPresent, sheet.AmountAbscent,
		sheet.AmountBanned, sheet.ControlType, sheet.Date, teacherId, groupId)
	if err != nil {
		log.Println(err)
		//if err.Error() == "pq: повторяющееся значение ключа нарушает ограничение уникальности \"sheet_pkey\"" {
		//	log.Println("Error defined")
		//}
		return sheet, err
	}

	return sheet, nil
}

func (r SheetsRepository) GetAVGSheetMark(ctx context.Context, sheetId int) (*float32, error) {
	query := r.db.Rebind(`
		SELECT AVG(together_mark)
		FROM sheet INNER JOIN sheet_marks ON sheetid =  sheet_marks.sheet
		WHERE sheetid = ?`)

	row := r.db.QueryRowContext(ctx, query, sheetId)
	var avgMark float32
	err := row.Scan(&avgMark)
	if err != nil {
		return nil, err
	}

	return &avgMark, nil
}

func (r SheetsRepository) DeleteAllData(ctx context.Context) error {
	query := r.db.Rebind(`
		TRUNCATE runner_marks ,
        sheet_marks ,
        sheet,
        groups_ ,
        runner,
        teachers,
        subjects,
        student RESTART IDENTITY;`)

	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
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
