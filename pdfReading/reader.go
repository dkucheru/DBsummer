package pdfReading

import (
	"bytes"
	"github.com/ledongthuc/pdf"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var NationalMarks = map[string]bool{
	"відмінно":     true,
	"добре":        true,
	"задовільно":   true,
	"незадовільно": true,
	"зараховано":   true,
	"не":           true,
}

type StudInfoFromPDF struct {
	Lastname     string
	FirstName    string
	MiddleName   string
	RecordBook   string
	SemesterMark int
	ControlMark  int
	TogetherMark int
	NationalMark string
	EctsMark     string
}
type ExtractedInformation struct {
	DocumentType      string
	IdDocument        int
	EducationalLevel  string
	Faculty           string
	EducationalYear   string
	GroupName         string
	Subject           string
	Semester          string
	ControlType       string
	Date              time.Time
	TeacherLastname   string
	TeacherFirstName  string
	TeacherMiddleName string
	AmountPresent     int
	AmountAbscent     int
	AmountBanned      int
	ReasonOfAbscence  string
	ExtractedStudents []*StudInfoFromPDF
}

func ParsePDFfile(content string) *ExtractedInformation {
	s := ExtractedInformation{}
	var allStudInfo []*StudInfoFromPDF
	regexWords := regexp.MustCompile(`([^_\s.:\-,]+)`)

	words := regexWords.FindAllStringSubmatch(content, -1)
	for i, v := range words {
		s1 := v[1]
		if s1 != "№" {
			s1 = strings.ToLower(formatWord(&s1))
		}
		//fmt.Println(s1)
		if s1 != "" {
			if s1 == "група" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				if s2 == "бігунець" {
					s.DocumentType = s2
				} else {
					s.DocumentType = "відомість"
				}
			}
			if s1 == "№" && (strings.ToLower(formatWord(&(words[i-1])[0])) == "відомість" ||
				strings.ToLower(formatWord(&(words[i-1])[0])) == "листок") {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				id, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
				}
				s.IdDocument = id
			}
			if s1 == "рівень" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				s.EducationalLevel = s2
			}
			if s1 == "факультет" {
				result := "факультет"
				j := i
				for strings.ToLower(formatWord(&(words[j+1])[0])) != "рік" {
					result += " "
					result += strings.ToLower(formatWord(&(words[j+1])[0]))
					j++
				}
				s.Faculty = result
			}
			if s1 == "навчання" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				s.EducationalYear = s2
			}
			if s1 == "перенесення" {
				result := strings.ToLower(formatWord(&(words[i+1])[0]))
				j := i + 1
				for strings.ToLower(formatWord(&(words[j+1])[0])) != "форма" {
					result += " "
					result += strings.ToLower(formatWord(&(words[j+1])[0]))
					j++
				}
				s.ReasonOfAbscence = result
			}
			if s1 == "група" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				s.GroupName = s2
			}
			if s1 == "дисципліна" {
				result := strings.ToLower(formatWord(&(words[i+1])[0]))
				j := i + 1
				for strings.ToLower(formatWord(&(words[j+1])[0])) != "семестр" {
					result += " "
					result += strings.ToLower(formatWord(&(words[j+1])[0]))
					j++
				}
				s.Subject = result
			}
			if s1 == "семестр" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				s.Semester = s2
			}
			if s1 == "контролю" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				s.ControlType = s2
			}
			if s1 == "дата" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				day, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
				}
				s2 = strings.ToLower(formatWord(&(words[i+2])[0]))
				month := getMonthNumber(s2)
				s2 = strings.ToLower(formatWord(&(words[i+3])[0]))
				year, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
				}
				s.Date = transformStringDate(day, month, year)
			}
			if s1 == "р" {
				lastname := strings.ToLower(formatWord(&(words[i+1])[0]))
				firstname := strings.ToLower(formatWord(&(words[i+2])[0]))
				middlename := strings.ToLower(formatWord(&(words[i+3])[0]))

				s.TeacherLastname = lastname
				s.TeacherFirstName = firstname
				s.TeacherMiddleName = middlename
			}

			if s1 == "екзамені" {
				s2 := strings.ToLower(formatWord(&(words[i+3])[0]))
				num, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
				}
				s.AmountPresent = num
			}
			if s1 == "екзамен" {
				s2 := strings.ToLower(formatWord(&(words[i+3])[0]))
				num, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
				}
				s.AmountAbscent = num
			}
			if s1 == "екзамену" {
				s2 := strings.ToLower(formatWord(&(words[i+3])[0]))
				num, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
				}
				s.AmountBanned = num
			}
			if s1 == "бп" {
				var stud StudInfoFromPDF

				recordbook := strings.ToLower(formatWord(&(words[i-2])[0]))
				recordbook += " "
				recordbook += strings.ToLower(formatWord(&(words[i-1])[0]))
				recordbook += " "
				recordbook += s1

				stud.RecordBook = recordbook

				semesterScore := strings.ToLower(formatWord(&(words[i+1])[0]))
				num, err := strconv.Atoi(semesterScore)
				if err != nil {
					log.Println(err)
				}
				stud.SemesterMark = num
				//next if statement checks if student has skipped the control
				//the if statement works when a student didn`t show up to exam
				s2 := strings.ToLower(formatWord(&(words[i+2])[0]))
				if NationalMarks[s2] == true {
					if s2 == "не" {
						national := s2 + " " + strings.ToLower(formatWord(&(words[i+3])[0]))
						stud.NationalMark = national
					} else {
						stud.NationalMark = s2
					}
					stud.ControlMark = 0
					stud.TogetherMark = 0
					stud.EctsMark = "Undefined"
				} else { //if we pass down to here then student has control marks and ects mark
					controlScore, err := strconv.Atoi(s2)
					if err != nil {
						log.Println(err)
					}
					stud.ControlMark = controlScore

					s2 = strings.ToLower(formatWord(&(words[i+3])[0]))
					togetherScore, err := strconv.Atoi(s2)
					if err != nil {
						log.Println(err)
					}
					stud.TogetherMark = togetherScore

					s2 = strings.ToLower(formatWord(&(words[i+4])[0]))
					if s2 == "не" {
						national := s2 + " " + strings.ToLower(formatWord(&(words[i+5])[0]))
						stud.NationalMark = national

						stud.EctsMark = formatWord(&(words[i+6])[0])
					} else {
						stud.NationalMark = s2
						stud.EctsMark = formatWord(&(words[i+5])[0])
					}
				}

				//now we have to get students full name (check if he has middlename)
				isNumberIfStudentHasMiddleName := formatWord(&(words[i-6])[0])
				_, err = strconv.Atoi(isNumberIfStudentHasMiddleName)
				if err != nil {
					//not a number => student doesn`t have a middle name
					//log.Println(err )
					log.Print(" A person doesn`t have a middle name")
					stud.FirstName = strings.ToLower(formatWord(&(words[i-3])[0]))
					stud.Lastname = strings.ToLower(formatWord(&(words[i-4])[0]))
				} else {
					stud.MiddleName = strings.ToLower(formatWord(&(words[i-3])[0]))
					stud.FirstName = strings.ToLower(formatWord(&(words[i-4])[0]))
					stud.Lastname = strings.ToLower(formatWord(&(words[i-5])[0]))
				}

				allStudInfo = append(allStudInfo, &stud)
			}
		}

	}
	s.ExtractedStudents = allStudInfo
	return &s
}

func formatWord(str *string) string {
	str1 := strings.TrimFunc(*str, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	return str1
}

func transformStringDate(day int, month int, year int) time.Time {
	return time.Date(year, time.Month(month), day, 12, 30, 0, 0, time.UTC)
}

func getMonthNumber(mon string) int {
	switch mon {
	case "січня":
		return 1
	case "лютого":
		return 2
	case "березня":
		return 3
	case "квітня":
		return 4
	case "травня":
		return 5
	case "червень":
		return 6
	case "липень":
		return 7
	case "Серпень":
		return 8
	case "Вересень":
		return 9
	case "Жовтень":
		return 10
	case "Листопад":
		return 11
	case "Грудень":
		return 12
	}
	return 1
}

func ReadPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
