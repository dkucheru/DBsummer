package postgres

import (
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
	"log"
)

type RunnerMarksRepository struct {
	*Repository
}

func (r RunnerMarksRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r RunnerMarksRepository) Get(ctx context.Context) error {
	panic("implement me")
}

func (r RunnerMarksRepository) GetRatingStudentWithRunners(ctx context.Context, sem string, ed_y string) ([]*structs.RatingWithRunners, error) {
	createParametrizedView := r.db.Rebind(`
		CREATE VIEW weighted_scores_runners AS
SELECT student_cipher,record_book_number, 
    last_name || ' ' || firstname || ' ' || middle_name AS pib_student,
      CASE
      WHEN runner_marks.together_mark IS NULL
        THEN sheet_marks.together_mark * credits
      ELSE  runner_marks.together_mark* credits
    END weight, credits
  
FROM ((((student INNER JOIN sheet_marks ON student_cipher = sheet_marks.student)
      INNER JOIN sheet ON sheet_marks.sheet = sheetid)
      INNER JOIN groups_ ON cipher = group_cipher)
      LEFT JOIN runner_marks ON mark_number = runner_marks.sheet_mark)
    LEFT JOIN subjects ON subjectid = groups_.subject
    
WHERE student_cipher IN (
    SELECT student_cipher
    FROM (student INNER JOIN sheet_marks ON student_cipher = sheet_marks.student)
        LEFT JOIN runner_marks ON mark_number = runner_marks.sheet_mark
    WHERE runner_marks.check_mark IS NOT NULL
  )
;`)
	_, err := r.db.Exec(createParametrizedView)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := r.db.Rebind(`
		SELECT record_book_number,pib_student, SUM(weight)::numeric/SUM(credits) AS rating
		FROM weighted_scores_runners
		WHERE semester = ? AND educationalyear = ?
		GROUP BY student_cipher,record_book_number,pib_student;`)

	rows, err := r.db.QueryContext(ctx, query, sem, ed_y)
	if err != nil {
		return nil, err
	}
	defer func() {
		e := rows.Close()
		if e != nil {
			log.Println(e)
		}
	}()

	var allInfo []*structs.RatingWithRunners
	for rows.Next() {
		var s structs.RatingWithRunners
		err = rows.Scan(&s.RecordBookNumber, &s.PibStudent, &s.Rating)
		if err != nil {
			return nil, err
		}
		allInfo = append(allInfo, &s)
	}

	dropParametrizedView := r.db.Rebind(`DROP VIEW weighted_scores_runners ;`)
	_, err = r.db.Exec(dropParametrizedView)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return allInfo, nil
}

func (r RunnerMarksRepository) PostRunnerMarksToDataBase(ctx context.Context, sheetMarkID int, runnerID int, runnerMarks *pdfReading.StudInfoFromPDF) error {
	query := r.db.Rebind(`
		INSERT into runner_marks(check_mark,national_mark,semester_mark,
	together_mark,ects_mark,sheet_mark,runner) VALUES(?,?,?,?,?,?,?);
		`)

	_, err := r.db.Exec(query, runnerMarks.ControlMark, runnerMarks.NationalMark, runnerMarks.SemesterMark,
		runnerMarks.TogetherMark, runnerMarks.EctsMark, sheetMarkID, runnerID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
