package structs

type Student struct {
	StudentCipher    string
	FirstName        string
	LastName         string
	MiddleName       string
	RecordBookNumber string
}

type StudentMarks struct {
	SubjectName  string
	MarkTogether int
	EctsMark     string
	TeacherPIB   string
}
