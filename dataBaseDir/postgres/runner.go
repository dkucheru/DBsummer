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
	//getTeacherCipher := r.db.Rebind(`
	//		SELECT teacher_cipher
	//		FROM teachers
	//		WHERE firstname = ? AND lastname = ? AND (middlename = ? OR middlename IS NULL OR middlename = '')
	//;
	//	`)
	//row := r.db.QueryRowContext(ctx, getTeacherCipher, runner.TeacherFirstName, runner.TeacherLastname, runner.TeacherMiddleName)
	//var teacherID string
	//err := row.Scan(&teacherID)
	//if err != nil {
	//	return err
	//}

	query := r.db.Rebind(`
			INSERT into runner(runner_number,date_of_compilation,date_of_expiration,postponing_reason,
		type_of_control,teacher) VALUES(?,?,?,?,?,?);
			`)

	_, err := r.db.Exec(query, runner.IdDocument, runner.Date, runner.Date, runner.ReasonOfAbscence, runner.ControlType, teacherId)
	if err != nil {
		return err
	}

	return nil
}
