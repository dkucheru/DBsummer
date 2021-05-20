package postgres

import (
	"DBsummer/structs"
	"context"
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

func (r SubjectsRepository) Get(ctx context.Context, id int) (*structs.Subject, error) {
	query := r.db.Rebind(`
		SELECT subjectid, subjectname, educationallevel, faculty FROM subjects WHERE subjectid = ?;
	`)

	row := r.db.QueryRowContext(ctx, query, id)

	s := Subjects{}
	err := row.Scan(&s.Subjectid, &s.Subjectname, &s.Educationallevel, &s.Faculty)
	if err != nil {
		return nil, err
	}

	return subjectFromDbx(ctx, s)
}

func subjectFromDbx(ctx context.Context, dbx Subjects) (*structs.Subject, error) {
	s := &structs.Subject{
		SubjectId:        dbx.Subjectid,
		SubjectName:      dbx.Subjectname,
		EducationalLevel: structs.Ed_level(dbx.Educationallevel),
		Faculty:          dbx.Faculty,
	}

	return s, nil
}
