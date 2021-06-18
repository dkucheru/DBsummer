package postgres

import (
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
	"log"
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

func (r RunnersRepository) GetRunnerInfoByID(ctx context.Context, id int) (*string, error) {
	query := r.db.Rebind(`
		SELECT  DISTINCT(lastname || ' ' || firstname || ' ' || middlename) AS pib,
		subjectname,groupname
		FROM (((((runner INNER JOIN runner_marks ON runner_marks.runner = runner_number)		
			INNER JOIN sheet_marks ON mark_number=sheet_mark)
			INNER JOIN sheet ON sheetid =  sheet_marks.sheet)
			INNER JOIN teachers ON teacher_cipher = sheet.teacher)
			INNER JOIN groups_ ON group_cipher = cipher)
			INNER JOIN subjects ON subjectid = groups_.subject
		WHERE runner_number = ?;`)

	row := r.db.QueryRowContext(ctx, query, id)
	var avgMark structs.SheetInfo
	err := row.Scan(&avgMark.PibTeacher, &avgMark.SubjectName, &avgMark.GroupName)
	if err != nil {
		return nil, err
	}

	result := "ПІБ викладача: " + avgMark.PibTeacher + "    Дисципліна: " + avgMark.SubjectName + "    Назва групи: " + avgMark.GroupName
	return &result, nil
}

func (r RunnersRepository) GetRunnerByID(ctx context.Context, id int) ([]*structs.SheetByID, error) {
	query := r.db.Rebind(`
		SELECT record_book_number,
		last_name || ' ' || firstname || ' ' || middle_name AS pib_student,
		runner_marks.semester_mark, runner_marks.check_mark,
		runner_marks.together_mark,runner_marks.national_mark,runner_marks.ects_mark
		FROM ((runner INNER JOIN runner_marks ON runner_marks.runner = runner_number)		
			INNER JOIN sheet_marks ON mark_number=sheet_mark)
			INNER JOIN student ON student_cipher=sheet_marks.student
		WHERE runner_number = ?;`)

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

	var allInfo []*structs.SheetByID
	for rows.Next() {
		var s structs.SheetByID
		err = rows.Scan(&s.RecordBook, &s.PibStudent, &s.SemesterMark, &s.ControlMark, &s.TogetherMark, &s.NationalMark, &s.ECTS)
		if err != nil {
			return nil, err
		}
		allInfo = append(allInfo, &s)
	}
	return allInfo, nil
}

func (r RunnersRepository) PostRunnerToDataBase(ctx context.Context, runner *pdfReading.ExtractedInformation, teacherId int) error {

	query := r.db.Rebind(`
			INSERT into runner(runner_number,date_of_compilation,date_of_expiration,postponing_reason,
		type_of_control,teacher) VALUES(?,?,?,?,?,?);
			`)

	_, err := r.db.Exec(query, runner.IdDocument, runner.Date, runner.ExpirationDate, runner.ReasonOfAbscence, runner.ControlType, teacherId)
	if err != nil {
		return err
	}

	return nil
}
