package postgres

import (
	"DBsummer/pdfReading"
	"context"
	"fmt"
	"log"
)

type SheetMarksRepository struct {
	*Repository
}

func (r SheetMarksRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r SheetMarksRepository) Get(ctx context.Context) error {
	panic("implement me")
}

func (r SheetMarksRepository) PostSheetMarksToDataBase(ctx context.Context, sheetID int, studentId int, sheetMarks *pdfReading.StudInfoFromPDF) (int, error) {

	query := r.db.Rebind(`
		INSERT into sheet_marks(check_mark,national_mark,semester_mark,
	together_mark,ects_mark,sheet,student) VALUES(?,?,?,?,?,?,?);
		`)

	_, err := r.db.Exec(query, sheetMarks.ControlMark, sheetMarks.NationalMark, sheetMarks.SemesterMark, sheetMarks.TogetherMark,
		sheetMarks.EctsMark, sheetID, studentId)
	if err != nil {
		log.Println(err)
		return sheetID, err
	}
	return sheetID, nil
}

func (r SheetMarksRepository) FindNezarahOrNezadov(ctx context.Context, studentId int, runner *pdfReading.ExtractedInformation) (*int, error) {

	query := r.db.Rebind(`
		SELECT mark_number as sheetMarkID
		FROM (((sheet_marks INNER JOIN student ON sheet_marks.student = student_cipher)
  				INNER JOIN sheet ON sheet_marks.sheet = sheetid)
  				INNER JOIN groups_ ON sheet.group_cipher = cipher)
  				INNER JOIN subjects ON subject = subjectid
    	WHERE student_cipher = ? AND subjectname = ? AND semester = ? AND educationalyear = ? AND
		(national_mark = 'не зараховано' OR national_mark = 'незадовільно');
		`)

	fmt.Println(studentId)
	fmt.Println(runner.Subject)
	fmt.Println(runner.Semester)
	fmt.Println(runner.EducationalYear)
	row := r.db.QueryRowContext(ctx, query, studentId, runner.Subject, runner.Semester, runner.EducationalYear)
	var sheetMarkID int
	err := row.Scan(&sheetMarkID)
	if err != nil {
		return nil, err
	}

	return &sheetMarkID, nil
}
