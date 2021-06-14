package structs

type Subject struct {
	SubjectId        int
	SubjectName      string
	EducationalLevel Ed_level
	Faculty          string
}

type SubjectName struct {
	SubjectId   int
	SubjectName string
}

type Ed_level string

const (
	Bachelor Ed_level = "Bachelor"
	Magistr  Ed_level = "Magistr"
)

func (c Ed_level) String() string {
	return string(c)
}
