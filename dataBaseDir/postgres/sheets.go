package postgres

import (
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
