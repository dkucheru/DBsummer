package postgres

import (
	"DBsummer/pdfReading"
	"context"
)

type RunnersRepository struct {
	*Repository
}

func (r RunnersRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r RunnersRepository) Get(ctx context.Context) error {
	panic("implement me")
}

func (r RunnersRepository) PostRunnerToDataBase(ctx context.Context, runner *pdfReading.ExtractedInformation, teacherId int) error {

	query := r.db.Rebind(`
			INSERT into runner(runner_number,date_of_compilation,date_of_expiration,postponing_reason,
		type_of_control,teacher) VALUES(?,?,?,?,?,?);
			`)

	_, err := r.db.Exec(query, runner.IdDocument, runner.Date, runner.ExpirationDate, runner.ReasonOfAbscence, runner.ControlType, teacherId)
	if err != nil {
		return err
	}

	return nil
}
