package postgres

import (
	"DBsummer/pdfReading"
	"context"
	"math/rand"
)

type GroupsRepository struct {
	*Repository
}

func (r GroupsRepository) Create(ctx context.Context) (id int, err error) {
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

func (r GroupsRepository) FindGroup(ctx context.Context, sheet *pdfReading.ExtractedInformation, subjectId int) (*int, error) {
	getGroupCipher := r.db.Rebind(`
		SELECT cipher
		FROM groups_
		WHERE groupname = ? AND educationalyear = ? AND semester = ? AND course = ? AND subject = ?
;
	`)

	//fmt.Println(sheet.GroupName +" "+ fmt.Sprint(sheet.Date.Year()) + " " + sheet.Semester + sheet.EducationalYear + fmt.Sprint(subjectId))

	row := r.db.QueryRowContext(ctx, getGroupCipher, sheet.GroupName, sheet.Date.Year(), sheet.Semester,
		sheet.EducationalYear, subjectId)
	var groupID int
	err := row.Scan(&groupID)
	if err != nil {
		return nil, err
	}
	//fmt.Println(groupID)
	return &groupID, nil
}

func (r GroupsRepository) AddGroup(ctx context.Context, sheet *pdfReading.ExtractedInformation, subjectId int) (*int, error) {
	query := r.db.Rebind(`
		INSERT into groups_(groupname,educationalyear,semester,course,subject) VALUES(?,?,?,?,?);
		`)
	//fmt.Println(sheet.GroupName +" "+ fmt.Sprint(sheet.Date.Year()) + " " + sheet.Semester + sheet.EducationalYear + fmt.Sprint(subjectId))
	_, err := r.db.Exec(query, sheet.GroupName, sheet.Date.Year(), sheet.Semester,
		sheet.EducationalYear, subjectId)

	if err != nil {
		return nil, err
	}
	fine := 1
	return &fine, nil
}
