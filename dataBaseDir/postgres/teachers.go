package postgres

import (
	"DBsummer/pdfReading"
	"context"
)

type TeachersRepository struct {
	*Repository
}

func (r TeachersRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r TeachersRepository) Get(ctx context.Context) error {
	panic("implement me")
}

func (r TeachersRepository) FindTeacher(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error) {
	getTeacherCipher := r.db.Rebind(`
		SELECT teacher_cipher
		FROM teachers
		WHERE firstname = ? AND lastname = ? AND (middlename = ? OR middlename IS NULL OR middlename = '')
;
	`)

	row := r.db.QueryRowContext(ctx, getTeacherCipher, sheet.TeacherFirstName, sheet.TeacherLastname, sheet.TeacherMiddleName)
	var teacherID int
	err := row.Scan(&teacherID)
	if err != nil {
		return nil, err
	}

	return &teacherID, nil
}

func (r TeachersRepository) AddTeacher(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error) {
	query := r.db.Rebind(`
		INSERT into teachers(firstname,lastname,middlename,scientificdegree,academictitles,post) VALUES(?,?,?,?,?,?);
		`)

	_, err := r.db.Exec(query, sheet.TeacherFirstName, sheet.TeacherLastname, sheet.TeacherMiddleName, sheet.ScientificDegree, sheet.AcademicTitle,
		sheet.TeacherPost)

	if err != nil {
		return nil, err
	}
	fine := 1
	return &fine, nil
}
