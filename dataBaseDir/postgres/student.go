package postgres

import (
	"DBsummer/structs"
	"context"
	"fmt"
	"log"
)

type StudentRepository struct {
	*Repository
}

func (r StudentRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r StudentRepository) Get(ctx context.Context, id int) (*structs.Student, error) {
	panic("implement me")
}

func (r StudentRepository) GetAllStudInfo(ctx context.Context) ([]*structs.AllStudInfo, error) {
	query := r.db.Rebind(`
		SELECT student_cipher, firstname || ' ' || last_name || ' ' || COALESCE(middle_name, '') AS pib, COALESCE(subjectname, '') AS subj, COALESCE(group_cipher, '') AS group, COALESCE(groupname,'') AS groupname
FROM (((student LEFT JOIN sheet_marks ON student_cipher = student ) LEFT JOIN sheet ON sheet = sheet.sheetid) LEFT JOIN groups_ ON group_cipher = cipher) LEFT JOIN subjects ON subject = subjectid;`)

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

	var allInfo []*structs.AllStudInfo
	for rows.Next() {
		var s structs.AllStudInfo
		err = rows.Scan(&s.RecordBook, &s.Pib, &s.SubjectName, &s.GroupCipher, &s.GroupName)
		if err != nil {
			return nil, err
		}
		fmt.Print("student.go file inside loop : ")
		fmt.Println(&s)
		allInfo = append(allInfo, &s)
	}
	fmt.Print("student.go after loop: ")
	fmt.Println(&allInfo)
	return allInfo, nil
}
