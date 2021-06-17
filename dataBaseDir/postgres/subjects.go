package postgres

import (
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
	"fmt"
	"log"
	"math/rand"
)

type SubjectsRepository struct {
	*Repository
}

func (r SubjectsRepository) Create(ctx context.Context) (id int, err error) {
	query := r.db.Rebind(`
		INSERT into tablen(id_t) VALUES(?);
	`)

	id = int(rand.Uint32())

	_, err = r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r SubjectsRepository) FindSubjectsWithYearParameter(ctx context.Context, year int) ([]*structs.SubjectName, error) {
	query := r.db.Rebind(`
		SELECT subjectid,subjectname
		FROM subjects
		WHERE subjectid IN (
  			SELECT groups_.subject
  			FROM groups_
  			WHERE cipher IN (
    		SELECT group_cipher
    		FROM sheet
    		WHERE date_part('year', date_of_compilation) < ?
  )
);
	`)
	rows, err := r.db.QueryContext(ctx, query, year)
	if err != nil {
		fmt.Println("error performing sql query")
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.Println(e)
		}
	}()

	var subjects []*structs.SubjectName
	for rows.Next() {
		var s structs.SubjectName
		err = rows.Scan(&s.SubjectId, &s.SubjectName)
		if err != nil {
			fmt.Println("error in scan")
			return nil, err
		}

		subjects = append(subjects, &s)
	}

	return subjects, nil
}

func (r SubjectsRepository) Get(ctx context.Context, id int) (*structs.Subject, error) {
	query := r.db.Rebind(`
		SELECT subjectid, subjectname, educationallevel, faculty FROM subjects WHERE subjectid = ?;
	`)

	row := r.db.QueryRowContext(ctx, query, id)

	s := &Subjects{}
	err := row.Scan(&s.Subjectid, &s.Subjectname, &s.Educationallevel, &s.Faculty)
	if err != nil {
		return nil, err
	}

	return subjectFromDbx(ctx, s)
}

func (r SubjectsRepository) FindSubject(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error) {
	getSubjectId := r.db.Rebind(`
		SELECT subjectid
		FROM subjects
		WHERE subjectname = ? AND educationallevel = ? AND faculty = ?
;
	`)

	row := r.db.QueryRowContext(ctx, getSubjectId, sheet.Subject, sheet.EducationalLevel, sheet.Faculty)
	var subjectID int
	err := row.Scan(&subjectID)
	if err != nil {
		return nil, err
	}

	return &subjectID, nil
}

func (r SubjectsRepository) AddSubject(ctx context.Context, sheet *pdfReading.ExtractedInformation) (*int, error) {
	query := r.db.Rebind(`
		INSERT into subjects(subjectname,educationallevel,faculty,credits) VALUES(?,?,?,?);
		`)

	_, err := r.db.Exec(query, sheet.Subject, sheet.EducationalLevel, sheet.Faculty, sheet.Credits)

	if err != nil {
		return nil, err
	}
	fine := 1
	return &fine, nil
}

func (r SubjectsRepository) GetAll(ctx context.Context) ([]*structs.Subject, error) {
	query := r.db.Rebind(`
		SELECT subjectid, subjectname, educationallevel, faculty FROM subjects;
	`)

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

	var subjects []*Subjects
	for rows.Next() {
		var s Subjects
		err = rows.Scan(&s.Subjectid, &s.Subjectname, &s.Educationallevel, &s.Faculty)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, &s)
	}

	return subjectsFromDbx(ctx, subjects)
}

func subjectsFromDbx(ctx context.Context, models []*Subjects) ([]*structs.Subject, error) {
	var subjects []*structs.Subject

	for _, m := range models {
		s, err := subjectFromDbx(ctx, m)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, s)
	}
	return subjects, nil
}

func subjectFromDbx(ctx context.Context, dbx *Subjects) (*structs.Subject, error) {
	s := &structs.Subject{
		SubjectId:        dbx.Subjectid,
		SubjectName:      dbx.Subjectname,
		EducationalLevel: structs.Ed_level(dbx.Educationallevel),
		Faculty:          dbx.Faculty,
	}

	return s, nil
}
