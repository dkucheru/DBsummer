package postgres

import (
	"DBsummer/pdfReading"
	"context"
)

type RunnerMarksRepository struct {
	*Repository
}

func (r RunnerMarksRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r RunnerMarksRepository) Get(ctx context.Context) error {
	panic("implement me")
}

func (r RunnerMarksRepository) PostSheetMarksToDataBase(ctx context.Context, runnerID int, runnerMarks *pdfReading.StudInfoFromPDF) error {
	//	getStudentCipher := r.db.Rebind(`
	//		SELECT student_cipher
	//		FROM student
	//		WHERE firstname = ? AND last_name = ? AND (middle_name = ? OR middle_name IS NULL OR middle_name = '')
	//;
	//	`)
	//	row := r.db.QueryRowContext(ctx, getStudentCipher, sheetMarks.FirstName, sheetMarks.Lastname, sheetMarks.MiddleName)
	//	var studentID string
	//	err := row.Scan(&studentID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	query := r.db.Rebind(`
	//		INSERT into runner_marks(check_mark,runner_mark_number,national_mark,semester_mark,together_mark,
	//ects_mark,sheet_mark,runner) VALUES(?,?,?,?,?,?,?,?);
	//		`)
	//
	//	id := int(rand.Uint32())
	//
	//	_, err = r.db.Exec(query, runnerMarks.ControlMark,id,runnerMarks.NationalMark,runnerMarks.SemesterMark,
	//		runnerMarks.TogetherMark,runnerMarks.EctsMark,)
	//	if err != nil {
	//		return err
	//	}
	return nil
}
