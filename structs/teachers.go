package structs

type Teacher struct {
	TeacherCipher    int
	FirstName        string
	MiddleName       string
	ScientificDegree string
	AcademicTitles   string
	Post             string
}

type TeacherPassStatistics struct {
	PIB            string
	TeacherCipher  int
	PassStatistics int
}

type TeacherPIB struct {
	Pib           string
	TeacherCipher int
}
