package postgres

import (
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
	"fmt"
	"log"
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

func (r GroupsRepository) FindGroupsOfScientist(ctx context.Context, scientificDegree string) ([]*structs.GroupOfScientist, error) {
	query := r.db.Rebind(`
		SELECT G.cipher,G.group_name,T.last_name,T.first_name
FROM (sheet Sh INNER JOIN Groups_ G ON G.cipher=Sh.Group_cipher) INNER JOIN Teachers T ON Sh.teacher=T.teacher_cipher
WHERE scientific_degree=?;

	`)

	rows, err := r.db.QueryContext(ctx, query, scientificDegree)
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

	var groups []*structs.GroupOfScientist
	for rows.Next() {
		var s structs.GroupOfScientist
		err = rows.Scan(&s.Cipher, &s.Groupname, &s.TeacherLastname, &s.TeacherFirstname)
		if err != nil {
			fmt.Println("error in scan")
			return nil, err
		}

		groups = append(groups, &s)
	}

	return groups, nil
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
