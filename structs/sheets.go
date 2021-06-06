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
