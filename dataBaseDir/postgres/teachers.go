package postgres

import (
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
	"fmt"
	"log"
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

func (r TeachersRepository) GetTeacherPassStatistics(ctx context.Context, passedOrNot string) ([]*structs.TeacherPassStatistics, error) {
	query := r.db.Rebind(`
		SELECT teacher_cipher, firstname || ' ' || lastname || ' ' || COALESCE(middlename, '') AS pibteach, Count(sheet_marks.mark_number) AS Amount_of_F
    FROM (teachers INNER JOIN sheet ON teachers.teacher_cipher=sheet.teacher) INNER JOIN sheet_marks ON sheet_marks.sheet=sheet.sheetid
    WHERE sheet_marks.national_mark=?
    Group BY teachers.teacher_cipher
    Having Count(sheet_marks.mark_number)>1;`)

	rows, err := r.db.QueryContext(ctx, query, passedOrNot)
	if err != nil {
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.Println(e)
		}
	}()
	var teachers []*structs.TeacherPassStatistics
	for rows.Next() {
		var s structs.TeacherPassStatistics
		err = rows.Scan(&s.TeacherCipher, &s.PIB, &s.PassStatistics)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, &s)
	}
	return teachers, nil
}

func (r TeachersRepository) GetTeacherPIBs(ctx context.Context) ([]*structs.TeacherPIB, error) {
	query := r.db.Rebind(`
		SELECT DISTINCT(firstname || ' ' || lastname || ' ' || COALESCE(middlename, '')) AS pibteach,
			teacher_cipher AS TeacherCipher
		FROM teachers;`)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.Println(e)
		}
	}()
	var pibs []*structs.TeacherPIB
	for rows.Next() {
		var s structs.TeacherPIB
		err = rows.Scan(&s.Pib, &s.TeacherCipher)
		if err != nil {
			return nil, err
		}
		pibs = append(pibs, &s)
	}
	fmt.Println(pibs)
	return pibs, nil
}
