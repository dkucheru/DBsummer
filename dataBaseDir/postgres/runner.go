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

func (r RunnersRepository) PostRunnerToDataBase(ctx context.Context, runner *pdfReading.ExtractedInformation) error {
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
	//		return err
	//	}
	//
	//	getGroupCipher := r.db.Rebind(`
	//		SELECT cipher
	//		FROM groups_
	//		WHERE groupname = ? AND educationalyear = ? AND semester = ?;
	//	`)
	//	row = r.db.QueryRowContext(ctx, getGroupCipher, sheet.GroupName, sheet.EducationalYear, sheet.Semester)
	//	var groupID string
	//	err = row.Scan(&groupID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	query := r.db.Rebind(`
	//		INSERT into sheet(runner_number,date_of_compilation,date_of_expiration,postponing_reason,
	//	type_of_control,teacher) VALUES(?,?,?,?,?,?);
	//		`)
	//
	//	_, err = r.db.Exec(query, sheet.IdDocument, sheet.AmountPresent, sheet.AmountAbscent,
	//		sheet.AmountBanned, sheet.ControlType, sheet.Date, teacherID, groupID)
	//	if err != nil {
	//		return err
	//	}
	//
	return nil
}
