package postgres

import (
	"DBsummer/pdfReading"
	"context"
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
	//	getStudentCipher := r.db.Rebind(`
	//		SELECT student_cipher
	//		FROM student
	//		WHERE firstname = ? AND last_name = ? AND (middle_name = ? OR middle_name IS NULL OR middle_name = '')
	//;
	//	`)
	//	row := r.db.QueryRowContext(ctx, getStudentCipher, sheetMarks.FirstName, sheetMarks.Lastname, sheetMarks.MiddleName)
	//	fmt.Println("group name : " + sheetMarks.FirstName)
	//	fmt.Println("year : " + sheetMarks.Lastname)
	//	fmt.Println("semester : " + sheetMarks.MiddleName)
	//
	//var studentID string
	//err := row.Scan(&studentID)
	//if err != nil {
	//	log.Println(err)
	//	return sheetID, err
	//}

	query := r.db.Rebind(`
		INSERT into sheet_marks(check_mark,national_mark,semester_mark,
	together_mark,ects_mark,sheet,student) VALUES(?,?,?,?,?,?,?);
		`)

	//rand.Seed(time.Now().UnixNano())
	//min := 1
	//max := 1000
	//random := rand.Intn(max-min+1) + min
	//random2 := rand.Intn(max-min+1) + min
	//id := int(sheetID + sheetMarks.SemesterMark + random + random2)
	//fmt.Println(id)

	_, err := r.db.Exec(query, sheetMarks.ControlMark, sheetMarks.NationalMark, sheetMarks.SemesterMark, sheetMarks.TogetherMark,
		sheetMarks.EctsMark, sheetID, studentId)
	if err != nil {
		log.Println(err)
		return sheetID, err
	}
	return sheetID, nil
}
