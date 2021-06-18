package structs

import "time"

type Sheet struct {
	SheetId            int
	NumberOfAttendees  int
	NumberOfAbsent     int
	NumberOfIneligible int
	TypeOfControl      string
	DateOfCompilation  time.Time
	Teacher            string
	GroupCipher        string
}

type SheetByQuery struct {
	SheetId      int
	PibStudent   string
	SemesterMark int
	CheckMark    int
	TogetherMark int
	NationalMark string
	EctsMark     string
}

type AvgSheetMarkByID struct {
	PibTeacher  string
	SubjectName string
	GroupName   string
	AvgMark     float32
}

type SheetByID struct {
	RecordBook   string
	PibStudent   string
	SemesterMark int
	ControlMark  int
	TogetherMark int
	NationalMark string
	ECTS         string
}
