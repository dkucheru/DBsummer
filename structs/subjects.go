package structs

type Subjects struct {
	SubjectId        int
	SubjectName      string
	EducationalLevel ed_level
	Faculty          string
}

type ed_level string

const (
	Bachelor ed_level = "Bachelor"
	Magistr  ed_level = "Magistr"
)

func (c ed_level) string() string {
	return string(c)
}
