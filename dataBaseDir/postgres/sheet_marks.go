package postgres

import (
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
	"fmt"
	"log"
)

type SheetMarksRepository struct {
	*Repository
}

func (r SheetMarksRepository) Create(ctx context.Context) (id int, err error) {
	panic("implement me")
}

func (r SheetMarksRepository) Get(ctx context.Context) error {
	panic("implement me")
}

func (r SheetMarksRepository) PostSheetMarksToDataBase(ctx context.Context, sheetID int, studentId int, sheetMarks *pdfReading.StudInfoFromPDF) (int, error) {

	query := r.db.Rebind(`
		INSERT into sheet_marks(check_mark,national_mark,semester_mark,
	together_mark,ects_mark,sheet,student) VALUES(?,?,?,?,?,?,?);
		`)

	_, err := r.db.Exec(query, sheetMarks.ControlMark, sheetMarks.NationalMark, sheetMarks.SemesterMark, sheetMarks.TogetherMark,
		sheetMarks.EctsMark, sheetID, studentId)
	if err != nil {
		log.Println(err)
		return sheetID, err
	}
	return sheetID, nil
}

func (r SheetMarksRepository) FindNezarahOrNezadov(ctx context.Context, studentId int, runner *pdfReading.ExtractedInformation) (*int, error) {

	query := r.db.Rebind(`
		SELECT mark_number as sheetMarkID
		FROM (((sheet_marks INNER JOIN student ON sheet_marks.student = student_cipher)
  				INNER JOIN sheet ON sheet_marks.sheet = sheetid)
  				INNER JOIN groups_ ON sheet.group_cipher = cipher)
  				INNER JOIN subjects ON subject = subjectid
    	WHERE student_cipher = ? AND subjectname = ? AND semester = ? AND educationalyear = ? AND
		(national_mark = 'не зараховано' OR national_mark = 'незадовільно');
		`)

	fmt.Println(studentId)
	fmt.Println(runner.Subject)
	fmt.Println(runner.Semester)
	fmt.Println(runner.EducationalYear)
	row := r.db.QueryRowContext(ctx, query, studentId, runner.Subject, runner.Semester, runner.EducationalYear)
	var sheetMarkID int
	err := row.Scan(&sheetMarkID)
	if err != nil {
		return nil, err
	}

	return &sheetMarkID, nil
}

func (r SheetMarksRepository) GetRatingStudents(ctx context.Context, sem string, ed_y string) ([]*structs.RatingWithRunners, error) {
	createParametrizedView := r.db.Rebind(`
CREATE VIEW weighted_scores_sheets AS
SELECT student_cipher,record_book_number, 
    last_name || ' ' || firstname || ' ' || middle_name AS pib_student,
      sheet_marks.together_mark * credits AS weight, credits
  
FROM ((((student INNER JOIN sheet_marks ON student_cipher = sheet_marks.student)
      INNER JOIN sheet ON sheet_marks.sheet = sheetid)
      INNER JOIN groups_ ON cipher = group_cipher)
      LEFT JOIN runner_marks ON mark_number = runner_marks.sheet_mark)
    LEFT JOIN subjects ON subjectid = groups_.subject
    
WHERE student_cipher NOT IN (
    SELECT student_cipher
    FROM (student INNER JOIN sheet_marks ON student_cipher = sheet_marks.student)
        LEFT JOIN runner_marks ON mark_number = runner_marks.sheet_mark
    WHERE runner_marks.check_mark IS NOT NULL
  )
AND semester = ? AND educationalyear = ?
;`)
	_, err := r.db.Exec(createParametrizedView, sem, ed_y)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := r.db.Rebind(`
		SELECT record_book_number,pib_student, SUM(weight)::numeric/SUM(credits) AS rating
		FROM weighted_scores_sheets
		GROUP BY student_cipher,record_book_number,pib_student;
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

	var allInfo []*structs.RatingWithRunners
	for rows.Next() {
		var s structs.RatingWithRunners
		err = rows.Scan(&s.RecordBookNumber, &s.PibStudent, &s.Rating)
		if err != nil {
			return nil, err
		}
		allInfo = append(allInfo, &s)
	}

	dropParametrizedView := r.db.Rebind(`DROP VIEW weighted_scores_sheets ;`)
	_, err = r.db.Exec(dropParametrizedView, sem, ed_y)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return allInfo, nil
}
