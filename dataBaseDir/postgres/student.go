package postgres

import (
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
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
func (r StudentRepository) FindStudent(ctx context.Context, sheetMarks *pdfReading.StudInfoFromPDF) (*int, error) {
	getStudentId := r.db.Rebind(`
		SELECT student_cipher
		FROM student
		WHERE firstname = ? AND last_name = ? AND (middle_name = ? OR middle_name IS NULL OR middle_name = '') AND record_book_number = ?;`)

	row := r.db.QueryRowContext(ctx, getStudentId, sheetMarks.FirstName, sheetMarks.Lastname, sheetMarks.MiddleName, sheetMarks.RecordBook)
	var subjectID int
	err := row.Scan(&subjectID)
	if err != nil {
		return nil, err
	}

	return &subjectID, nil
}

func (r StudentRepository) GetPIBAllStudents(ctx context.Context, q string) ([]*structs.StudentPIB, error) {
	getStudentsPIBs := r.db.Rebind(`
		SELECT DISTINCT (last_name || ' ' ||firstname || ' ' || COALESCE(middle_name,'')) AS pib_student, student_cipher AS StudentCipher, record_book_number AS RecordNumber
		FROM student WHERE (last_name || ' ' ||firstname || ' ' || COALESCE(middle_name,'')) ILIKE ?
		;`)

	rows, err := r.db.QueryContext(ctx, getStudentsPIBs, q+"%")
	if err != nil {
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.Println(e)
		}
	}()

	var allPIBs []*structs.StudentPIB
	for rows.Next() {
		var s structs.StudentPIB
		err = rows.Scan(&s.Pib, &s.StudentCipher, &s.RecordNumber)
		if err != nil {
			return nil, err
		}
		allPIBs = append(allPIBs, &s)
	}

	return allPIBs, nil
}

func (r StudentRepository) AddStudent(ctx context.Context, sheetMarks *pdfReading.StudInfoFromPDF) (*int, error) {
	query := r.db.Rebind(`
		INSERT into student(firstname,last_name,middle_name,record_book_number) VALUES(?,?,?,?);
		`)

	_, err := r.db.Exec(query, sheetMarks.FirstName, sheetMarks.Lastname, sheetMarks.MiddleName, sheetMarks.RecordBook)

	if err != nil {
		return nil, err
	}
	fine := 1
	return &fine, nil
}

func (r StudentRepository) GetStudentByPIB(ctx context.Context, fn string, ln string, mn string, year string) ([]*structs.StudentMarks, error) {
	query := r.db.Rebind(`
		SELECT
		COALESCE(subjectname, '') AS subj,
		together_mark, ects_mark,
		teachers.firstname || ' ' || lastname || ' ' || COALESCE(middlename, '') AS pibteach
FROM ((((student JOIN sheet_marks ON student_cipher = student ) 
	    JOIN sheet ON sheet = sheet.sheetid) 
	   JOIN groups_ ON group_cipher = cipher) 
	   JOIN subjects ON subject = subjectid)
	   JOIN teachers ON teacher_cipher = sheet.teacher
WHERE educationalyear = ? AND student.firstname = ? AND last_name = ? AND (middle_name = ? OR middle_name IS NULL);`)

	rows, err := r.db.QueryContext(ctx, query, year, fn, ln, mn)
	if err != nil {
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.Println(e)
		}
	}()

	var allInfo []*structs.StudentMarks
	for rows.Next() {
		var s structs.StudentMarks
		err = rows.Scan(&s.SubjectName, &s.MarkTogether, &s.EctsMark, &s.TeacherPIB)
		if err != nil {
			return nil, err
		}
		allInfo = append(allInfo, &s)
	}
	return allInfo, nil
}

func (r StudentRepository) GetAllStudInfo(ctx context.Context) ([]*structs.AllStudInfo, error) {
	query := r.db.Rebind(`
		SELECT record_book_number, last_name || ' ' || firstname || ' ' || middle_name AS pib_stud,
		subjectname,groupname,
    CASE 
      WHEN runner_marks.together_mark ISNULL
        THEN sheet_marks.together_mark
      WHEN runner_marks.together_mark NOTNULL
        THEN runner_marks.together_mark
    END mark,
    sheetid,
	COALESCE(NULLIF(COALESCE(runner_number,0),0)::character varying(15),'')
    AS runner,
    educationalyear,
    semester
FROM (((((student INNER JOIN sheet_marks ON student_cipher = sheet_marks.student)
		INNER JOIN sheet ON sheetid = sheet_marks.sheet)
		INNER JOIN groups_ ON cipher = group_cipher)
		INNER JOIN subjects ON subjectid = groups_.subject)
		LEFT JOIN runner_marks ON sheet_mark = mark_number)
		LEFT JOIN runner ON runner_number = runner_marks.runner;`)

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
		err = rows.Scan(&s.RecordBook, &s.PibStud, &s.SubjectName, &s.GroupName, &s.Mark, &s.SheetId, &s.RunnerId, &s.EducationalYear,
			&s.Semester)
		if err != nil {
			return nil, err
		}
		allInfo = append(allInfo, &s)
	}
	return allInfo, nil
}

func (r StudentRepository) GetAllBorjniki(ctx context.Context) ([]*structs.AllStudInfo, error) {
	query := r.db.Rebind(`
		SELECT COALESCE(record_book_number,'вільний слухач'), 
		student.firstname || ' ' || last_name || ' ' || COALESCE(middle_name, '') AS pibstud, 
		COALESCE(subjectname, '') AS subj,
		COALESCE(group_cipher, '') AS group, 
		COALESCE(groupname,'') AS groupname,
		COALESCE(teachers.firstname,'') || ' ' || COALESCE(lastname,'') || ' ' || COALESCE(middlename, '') AS pibteach
FROM ((((student LEFT JOIN sheet_marks ON student_cipher = student ) 
	   LEFT JOIN sheet ON sheet = sheet.sheetid) 
	  LEFT JOIN groups_ ON group_cipher = cipher) 
	  LEFT JOIN subjects ON subject = subjectid)
	  LEFT JOIN teachers ON teacher_cipher = sheet.teacher
WHERE check_mark IS NULL AND subjectname IS NOT NULL;`)

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
		err = rows.Scan(&s.RecordBook, &s.PibStud, &s.SubjectName, &s.GroupName)
		if err != nil {
			return nil, err
		}
		allInfo = append(allInfo, &s)
	}
	return allInfo, nil
}

func (r StudentRepository) GetAllStudentMarksByID(ctx context.Context, id int) ([]*structs.StudentAllMarks, error) {
	query := r.db.Rebind(`
		SELECT subjectname,sheet_marks.together_mark,sheetid,
		COALESCE(NULLIF(COALESCE(runner_marks.together_mark,0),0)::character varying(15),'0') AS runner_mark,
		COALESCE(NULLIF(COALESCE(runner_number,0),0)::character varying(15),'0') AS runnerid,
  		semester,educationalyear
FROM (((((student INNER JOIN sheet_marks ON student_cipher = sheet_marks.student)
      INNER JOIN sheet ON sheet_marks.sheet = sheetid)
      INNER JOIN groups_ ON cipher = group_cipher)
      LEFT JOIN runner_marks ON mark_number = runner_marks.sheet_mark)
    LEFT JOIN subjects ON subjectid = groups_.subject)
	LEFT JOIN runner ON runner_marks.runner = runner_number
WHERE student_cipher = ?;`)

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.Println(e)
		}
	}()

	var allInfo []*structs.StudentAllMarks
	for rows.Next() {
		var s structs.StudentAllMarks
		err = rows.Scan(&s.SubjectName, &s.SheetMark, &s.SheetID, &s.RunnerMark, &s.RunnerID, &s.Semester, &s.EducationalYear)
		if err != nil {
			return nil, err
		}
		allInfo = append(allInfo, &s)
	}
	return allInfo, nil

}
