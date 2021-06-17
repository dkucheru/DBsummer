package postgres

import (
	"DBsummer/pdfReading"
	"context"
	"log"
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

func (r RunnerMarksRepository) PostRunnerMarksToDataBase(ctx context.Context, sheetMarkID int, runnerID int, runnerMarks *pdfReading.StudInfoFromPDF) error {
	query := r.db.Rebind(`
		INSERT into runner_marks(check_mark,national_mark,semester_mark,
	together_mark,ects_mark,sheet_mark,runner) VALUES(?,?,?,?,?,?,?);
		`)

	_, err := r.db.Exec(query, runnerMarks.ControlMark, runnerMarks.NationalMark, runnerMarks.SemesterMark,
		runnerMarks.TogetherMark, runnerMarks.EctsMark, sheetMarkID, runnerID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
