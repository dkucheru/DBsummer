package postgres

import (
	"DBsummer/pdfReading"
	"context"
	"math/rand"
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

func (r SheetMarksRepository) PostSheetMarksToDataBase(ctx context.Context, sheetID int, sheetMarks *pdfReading.StudInfoFromPDF) error {
	getStudentCipher := r.db.Rebind(`
		SELECT student_cipher
		FROM student
		WHERE firstname = ? AND last_name = ? AND (middle_name = ? OR middle_name IS NULL OR middle_name = '')
;
	`)
	row := r.db.QueryRowContext(ctx, getStudentCipher, sheetMarks.FirstName, sheetMarks.Lastname, sheetMarks.MiddleName)
	var studentID string
	err := row.Scan(&studentID)
	if err != nil {
		return err
	}

	query := r.db.Rebind(`
		INSERT into sheet_marks(check_mark,mark_number,national_mark,semester_mark,
	together_mark,ects_mark,sheet,student) VALUES(?,?,?,?,?,?,?,?);
		`)

	id := int(rand.Uint32())

	_, err = r.db.Exec(query, sheetMarks.ControlMark, id, sheetMarks.NationalMark, sheetMarks.SemesterMark, sheetMarks.TogetherMark,
		sheetMarks.EctsMark, sheetID, studentID)
	if err != nil {
		return err
	}
	return nil
}
