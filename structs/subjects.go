package structs

type Subject struct {
	SubjectId        int
	SubjectName      string
	EducationalLevel Ed_level
	Faculty          string
}

type Ed_level string

const (
	Bachelor Ed_level = "Bachelor"
	Magistr  Ed_level = "Magistr"
)

func (c Ed_level) String() string {
	return string(c)
}

type Group struct {
	Cipher          string
	groupname       string
	educationalyear string
	semester        string
	course          string
	subject         int
}
