package pdfReading

import (
	"errors"
	"fmt"
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
}
var EctsMarks = map[string]bool{
	"A": true,
	"А": true,
	"B": true,
	"В": true,
	"C": true,
	"С": true,
	"D": true,
	"E": true,
	"Е": true,
	"F": true,
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
	Course            string
	GroupName         string
	Subject           string
	Semester          string
	ControlType       string
	Date              time.Time
	ExpirationDate    time.Time
	TeacherLastname   string
	TeacherFirstName  string
	TeacherMiddleName string
	AmountPresent     int
	AmountAbscent     int
	AmountBanned      int
	ReasonOfAbscence  string
	AcademicTitle     string
	ScientificDegree  string
	TeacherPost       string
	Credits           int
	ExtractedStudents []*StudInfoFromPDF
}

func ParsePDFfile(content string) (*ExtractedInformation, error) {

	numPresent := 0
	numNotAllowed := 0
	s := ExtractedInformation{}
	var allStudInfo []*StudInfoFromPDF
	regexWords := regexp.MustCompile(`([^_\s.:\-,«»]+)`)

	words := regexWords.FindAllStringSubmatch(content, -1)

	words = removeEmptyWords(&words)
	for i, v := range words {
		s1 := v[1]
		if s1 != "№" {
			s1 = strings.ToLower(formatWord(&s1))
		}
		//fmt.Println(s1)
		if s1 != "" {

			if s1 == "кандидат" {
				result := "кандидат"
				j := i
				for strings.ToLower(formatWord(&(words[j+1])[0])) != "наук" {
					result += " "
					result += strings.ToLower(formatWord(&(words[j+1])[0]))
					j++
				}
				result += " "
				result += "наук"
				s.ScientificDegree = result
			}

			if s1 == "доктор" {
				result := "доктор"
				j := i
				for strings.ToLower(formatWord(&(words[j+1])[0])) != "наук" {
					result += " "
					result += strings.ToLower(formatWord(&(words[j+1])[0]))
					j++
				}
				result += " "
				result += "наук"
				s.ScientificDegree = result
			}

			if strings.Contains(s1, "доцент") {
				s.AcademicTitle = "доцент"
			}
			if strings.Contains(s1, "професор") {
				s.AcademicTitle = "професор"
			}

			if s1 == "старший" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				if strings.Contains(s2, "дослідник") {
					result := s1 + " дослідник"
					s.AcademicTitle = result
				}
				if strings.Contains(s2, "викладач") {
					result := s1 + " викладач"
					s.TeacherPost = result
				}
			}

			if s1 == "молодший" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				if strings.Contains(s2, "дослідник") {
					result := s1 + " дослідник"
					s.AcademicTitle = result
				}
				if strings.Contains(s2, "викладач") {
					result := s1 + " викладач"
					s.TeacherPost = result
				}
			}

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
				//next line is sh*t code but it`s the only way to solve this case
				s2_clean := strings.ReplaceAll(s2, "освітній", "")

				if s2 != s2_clean {
					return nil, errors.New("не валідний файл, використовуйте шаблон з сайту")
				}

				id, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
					return nil, errors.New("помилка при зчитуванні номера документу")
				}
				s.IdDocument = id
			}
			if s1 == "рівень" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))

				s2 = strings.ReplaceAll(s2, "факультет", "")
				s.EducationalLevel = s2
			}
			if s1 == "факультет" || s1 == "бакалаврфакультет" || s1 == "магістрфакультет" {
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
				s.Course = s2
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
				s2 = strings.ReplaceAll(s2, "дисципліна", "")
				s.GroupName = s2
			}
			if s1 == "бали" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				s2 = strings.ReplaceAll(s2, "форма", "")
				amount, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
					return nil, errors.New("помилка при зчитуванні залікових балів")
				}
				s.Credits = amount
			}

			if s1 == "дисципліна" || strings.Contains(s1, "дисципліна") {
				result := strings.ToLower(formatWord(&(words[i+1])[0]))
				j := i + 1
				for !strings.Contains(strings.ToLower(formatWord(&(words[j+1])[0])), "семестр") {
					result += " "
					result += strings.ToLower(formatWord(&(words[j+1])[0]))
					j++
				}

				check := strings.ReplaceAll(strings.ToLower(formatWord(&(words[j+1])[0])), "семестр", "")
				if check != "" {
					result += " "
					result += check
				}
				s.Subject = result
			}
			if s1 == "семестр" || strings.Contains(s1, "семестр") {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				s.Semester = s2
			}
			if s1 == "контролю" {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				s.ControlType = s2
			}
			if s1 == "дійсне" && strings.ToLower(formatWord(&(words[i+1])[0])) == "до" &&
				isNumber(formatWord(&(words[i+2])[0])) {
				s2 := strings.ToLower(formatWord(&(words[i+2])[0]))
				day, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
					return nil, errors.New("помилка при зчитуванні дня дати")
				}
				s2 = strings.ToLower(formatWord(&(words[i+3])[0]))
				month := getMonthNumber(s2)
				s2 = strings.ToLower(formatWord(&(words[i+4])[0]))
				year, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
					return nil, errors.New("помилка при зчитуванні року дати")
				}
				s.ExpirationDate = transformStringDate(day, month, year)
			}
			if s1 == "дата" && (strings.ToLower(formatWord(&(words[i-1])[0])) == "екзамен" ||
				strings.ToLower(formatWord(&(words[i-1])[0])) == "залік" ||
				strings.ToLower(formatWord(&(words[i-1])[0])) == "іспит") {
				s2 := strings.ToLower(formatWord(&(words[i+1])[0]))
				day, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
					return nil, errors.New("помилка при зчитуванні дня дати")
				}
				s2 = strings.ToLower(formatWord(&(words[i+2])[0]))
				month := getMonthNumber(s2)
				s2 = strings.ToLower(formatWord(&(words[i+3])[0]))
				year, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
					return nil, errors.New("помилка при зчитуванні року дати")
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
					return nil, errors.New("помилка при зчитуванні кількості присутніх")
				}
				s.AmountPresent = num
			}
			if s1 == "екзамен" && strings.ToLower(formatWord(&(words[i-1])[0])) == "на" {
				s2 := strings.ToLower(formatWord(&(words[i+3])[0]))
				num, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
					return nil, errors.New("помилка при зчитуванні кількості відсутніх")
				}
				s.AmountAbscent = num
			}
			if s1 == "екзамену" {
				s2 := strings.ToLower(formatWord(&(words[i+3])[0]))
				num, err := strconv.Atoi(s2)
				if err != nil {
					log.Println(err)
					return nil, errors.New("помилка при зчитуванні кількості не допущених")
				}
				s.AmountBanned = num
			}

			if strings.Contains(s1, "бп") || strings.Contains(s1, "мп") {
				s2 := strings.ReplaceAll(s1, "бп", "")
				s2 = strings.ReplaceAll(s2, "мп", "")

				if s2 != "" {
					return nil, errors.New("не валідний файл, використовуйте шаблон з сайту")
				}
			}
			if s1 == "бп" || s1 == "мп" {
				var stud StudInfoFromPDF

				//get recordbook info
				recordbook := strings.ToLower(formatWord(&(words[i-2])[0]))
				recordbook += " "
				recordbook += strings.ToLower(formatWord(&(words[i-1])[0]))
				recordbook += " "
				recordbook += s1
				stud.RecordBook = recordbook

				//now parse scores and marks
				i_plus_1 := strings.ToLower(formatWord(&(words[i+1])[0]))
				i_plus_2 := strings.ToLower(formatWord(&(words[i+2])[0]))
				i_plus_3 := strings.ToLower(formatWord(&(words[i+3])[0]))
				i_plus_4 := strings.ToLower(formatWord(&(words[i+4])[0]))
				i_plus_5 := strings.ToLower(formatWord(&(words[i+5])[0]))
				i_plus_6 := strings.ToLower(formatWord(&(words[i+6])[0]))

				if i_plus_1 == "не" && strings.Contains(i_plus_2, "відв") {
					stud.SemesterMark = 0
					stud.ControlMark = 0
					stud.TogetherMark = 0
					stud.NationalMark = i_plus_1 + " " + i_plus_2
					stud.EctsMark = "Undefined"

				} else if strings.Contains(i_plus_1, "невід") {
					stud.SemesterMark = 0
					stud.ControlMark = 0
					stud.TogetherMark = 0
					stud.NationalMark = "не " + strings.ReplaceAll(i_plus_1, "не", "")
					stud.EctsMark = "Undefined"

				} else if (i_plus_2 == "не" && strings.Contains(i_plus_3, "відв")) ||
					(i_plus_3 == "не" && strings.Contains(i_plus_4, "відв")) ||
					(i_plus_4 == "не" && strings.Contains(i_plus_5, "відв")) ||
					strings.Contains(i_plus_2, "невід") ||
					strings.Contains(i_plus_3, "невід") ||
					strings.Contains(i_plus_4, "невід") {
					return nil, errors.New("у студента, який не відвідував не має бути оцінок")

				} else if i_plus_4 == "не" && strings.Contains(i_plus_5, "допущ") ||
					(strings.Contains(i_plus_4, "недоп")) {

					intCotrol, err := strconv.Atoi(i_plus_2)
					if err != nil {
						return nil, err
					}
					if intCotrol != 0 {
						return nil, errors.New("не допущений студент не може мати оцінки>0 за контроль знань")
					}
					intsemester, err := strconv.Atoi(i_plus_1)
					if err != nil {
						return nil, err
					}
					inttogether, err := strconv.Atoi(i_plus_3)
					if err != nil {
						return nil, err
					}
					if intsemester != inttogether {
						return nil, errors.New("оцінка за триместр та оцінка разом мають бути однакові для студента, що не допущений")
					}
					stud.SemesterMark = intsemester
					stud.ControlMark = 0
					stud.TogetherMark = inttogether

					if i_plus_4 == "не" && strings.Contains(i_plus_5, "допущ") {
						if EctsMarks[formatWord(&(words[i+6])[0])] || isNumber(i_plus_6) {
							stud.NationalMark = i_plus_4 + " " + i_plus_5
						} else if EctsMarks[formatWord(&(words[i+7])[0])] || isNumber(formatWord(&(words[i+7])[0])) {
							stud.NationalMark = i_plus_4 + " " + i_plus_5 + i_plus_6
						} else {
							return nil, errors.New("помилка зчитування нацональної оцінки не допущено")
						}
					} else if strings.Contains(i_plus_4, "недоп") {
						if EctsMarks[formatWord(&(words[i+5])[0])] || isNumber(i_plus_5) {
							stud.NationalMark = "не " + strings.ReplaceAll(i_plus_4, "не", "")
						} else if EctsMarks[formatWord(&(words[i+6])[0])] || isNumber(i_plus_6) {
							stud.NationalMark = "не " + strings.ReplaceAll(i_plus_4, "не", "") + i_plus_5
						} else {
							return nil, errors.New("помилка зчитування нацональної оцінки не допущено")
						}
					}
					stud.EctsMark = "F"
					numNotAllowed += 1

				} else if (i_plus_3 == "не" && strings.Contains(i_plus_4, "допущ")) ||
					strings.Contains(i_plus_3, "недоп") {

					intsemester, err := strconv.Atoi(i_plus_1)
					if err != nil {
						return nil, err
					}
					inttogether, err := strconv.Atoi(i_plus_2)
					if err != nil {
						return nil, err
					}
					if intsemester != inttogether {
						return nil, errors.New("оцінка за триместр та оцінка разом мають бути однакові для студента, що не допущений")
					}
					stud.SemesterMark = intsemester
					stud.ControlMark = 0
					stud.TogetherMark = inttogether

					if i_plus_3 == "не" && strings.Contains(i_plus_4, "допущ") {
						if EctsMarks[formatWord(&(words[i+5])[0])] || isNumber(i_plus_5) {
							stud.NationalMark = i_plus_3 + " " + i_plus_4
						} else if EctsMarks[formatWord(&(words[i+6])[0])] || isNumber(i_plus_6) {
							stud.NationalMark = i_plus_3 + " " + i_plus_4 + i_plus_5
						} else {
							return nil, errors.New("помилка зчитування нацональної оцінки не допущено")
						}
					} else if strings.Contains(i_plus_3, "недоп") {
						if EctsMarks[formatWord(&(words[i+4])[0])] || isNumber(i_plus_4) {
							stud.NationalMark = "не " + strings.ReplaceAll(i_plus_3, "не", "")
						} else if EctsMarks[formatWord(&(words[i+5])[0])] || isNumber(i_plus_5) {
							stud.NationalMark = "не " + strings.ReplaceAll(i_plus_3, "не", "") + i_plus_4
						} else {
							return nil, errors.New("помилка зчитування нацональної оцінки не допущено")
						}
					}
					stud.EctsMark = "F"
					numNotAllowed += 1
				} else if (i_plus_2 == "не" && strings.Contains(i_plus_3, "допущ")) ||
					strings.Contains(i_plus_2, "недоп") {

					intsemester, err := strconv.Atoi(i_plus_1)
					if err != nil {
						return nil, err
					}

					stud.SemesterMark = intsemester
					stud.ControlMark = 0
					stud.TogetherMark = intsemester

					if i_plus_2 == "не" && strings.Contains(i_plus_3, "допущ") {
						if EctsMarks[formatWord(&(words[i+4])[0])] || isNumber(i_plus_4) {
							stud.NationalMark = i_plus_2 + " " + i_plus_3
						} else if EctsMarks[formatWord(&(words[i+5])[0])] || isNumber(i_plus_5) {
							stud.NationalMark = i_plus_2 + " " + i_plus_3 + i_plus_4
						} else {
							return nil, errors.New("помилка зчитування нацональної оцінки не допущено")
						}
					} else if strings.Contains(i_plus_2, "недоп") {
						if EctsMarks[formatWord(&(words[i+3])[0])] || isNumber(i_plus_3) {
							stud.NationalMark = "не " + strings.ReplaceAll(i_plus_2, "не", "")
						} else if EctsMarks[formatWord(&(words[i+4])[0])] || isNumber(i_plus_4) {
							stud.NationalMark = "не " + strings.ReplaceAll(i_plus_2, "не", "") + i_plus_3
						} else {
							return nil, errors.New("помилка зчитування нацональної оцінки не допущено")
						}
					}
					stud.EctsMark = "F"
					numNotAllowed += 1

				} else if (i_plus_1 == "не" && strings.Contains(i_plus_2, "допущ")) ||
					strings.Contains(i_plus_1, "недоп") {
					return nil, errors.New("не допущений студент повинен мати оцінку за триместр")

				} else if NationalMarks[i_plus_1] || NationalMarks[i_plus_2] || NationalMarks[i_plus_3] ||
					i_plus_1 == "незараховано" || i_plus_2 == "незараховано" || i_plus_3 == "незараховано" ||
					strings.Contains(i_plus_1, "незар") || strings.Contains(i_plus_2, "незар") || strings.Contains(i_plus_3, "незар") ||
					(i_plus_1 == "не" && (i_plus_2 == "зараховано" || strings.Contains(i_plus_2, "зар"))) ||
					(i_plus_2 == "не" && (i_plus_3 == "зараховано" || strings.Contains(i_plus_3, "зар"))) ||
					(i_plus_3 == "не" && (i_plus_4 == "зараховано" || strings.Contains(i_plus_4, "зар"))) {
					return nil, errors.New("допущена до роботи людина повинна мати 3 оцінки : триместр, контроль, разом")

				} else if NationalMarks[i_plus_4] ||
					strings.Contains(i_plus_4, "незар") ||
					(i_plus_4 == "не" && strings.Contains(i_plus_5, "зар")) ||
					strings.Contains(i_plus_4, "незад") ||
					(i_plus_4 == "не" && strings.Contains(i_plus_5, "зад")) {

					intsemester, err := strconv.Atoi(i_plus_1)
					if err != nil {
						return nil, err
					}
					intcontrol, err := strconv.Atoi(i_plus_2)
					if err != nil {
						return nil, err
					}
					inttogether, err := strconv.Atoi(i_plus_3)
					if err != nil {
						return nil, err
					}
					if (intsemester + intcontrol) != inttogether {
						return nil, errors.New("неправильно порахована оцінка разом : " + fmt.Sprint(inttogether))
					}
					stud.SemesterMark = intsemester
					stud.ControlMark = intcontrol
					stud.TogetherMark = inttogether

					if NationalMarks[i_plus_4] {
						stud.NationalMark = i_plus_4
						if EctsMarks[formatWord(&(words[i+5])[0])] {
							stud.EctsMark = formatWord(&(words[i+5])[0])
							if !checkMarkAccordance(stud.EctsMark, stud.NationalMark, stud.TogetherMark) {
								return nil, errors.New("перевірте у студента [" + stud.RecordBook + "] оцінку національну, єктс та разом ")
							}
						} else if isNumber(i_plus_5) {
							return nil, errors.New("у студента що складав роботу має бути оцінка ЄКТС ")
						}
						numPresent += 1
					} else if strings.Contains(i_plus_4, "незад") || strings.Contains(i_plus_4, "незар") {

						if strings.Contains(i_plus_4, "незад") {
							stud.NationalMark = "незадовільно"
						} else {
							stud.NationalMark = "не зараховано"
						}

						if EctsMarks[formatWord(&(words[i+5])[0])] {
							stud.EctsMark = formatWord(&(words[i+5])[0])
							if !checkMarkAccordance(stud.EctsMark, stud.NationalMark, stud.TogetherMark) {
								return nil, errors.New("перевірте у студента [" + stud.RecordBook + "] оцінку національну, єктс та разом ")
							}
						} else if isNumber(i_plus_5) {
							return nil, errors.New("у студента що складав роботу має бути оцінка ЄКТС ")
						} else if isNumber(i_plus_6) {
							if !EctsMarks[formatWord(&(words[i+5])[0])] {
								return nil, errors.New("у студента що складав роботу має бути оцінка ЄКТС ")
							}
						} else if isNumber(formatWord(&(words[i+7])[0])) {
							if !EctsMarks[formatWord(&(words[i+6])[0])] {
								return nil, errors.New("у студента що складав роботу має бути оцінка ЄКТС ")
							} else {
								stud.EctsMark = formatWord(&(words[i+6])[0])
								if !checkMarkAccordance(stud.EctsMark, stud.NationalMark, stud.TogetherMark) {
									return nil, errors.New("перевірте у студента [" + stud.RecordBook + "] оцінку національну, єктс та разом ")
								}
							}
						}
						numPresent += 1
					} else if (i_plus_4 == "не" && strings.Contains(i_plus_5, "зад")) ||
						(i_plus_4 == "не" && strings.Contains(i_plus_5, "зар")) {

						if strings.Contains(i_plus_5, "зад") {
							stud.NationalMark = "незадовільно"
						} else {
							stud.NationalMark = "не зараховано"
						}

						if EctsMarks[formatWord(&(words[i+6])[0])] {
							stud.EctsMark = formatWord(&(words[i+6])[0])
							if !checkMarkAccordance(stud.EctsMark, stud.NationalMark, stud.TogetherMark) {
								return nil, errors.New("перевірте у студента [" + stud.RecordBook + "] оцінку національну, єктс та разом ")
							}
						} else if isNumber(i_plus_6) {
							return nil, errors.New("у студента що складав роботу має бути оцінка ЄКТС ")
						} else if isNumber(formatWord(&(words[i+7])[0])) {
							if !EctsMarks[formatWord(&(words[i+6])[0])] {
								return nil, errors.New("у студента що складав роботу має бути оцінка ЄКТС ")
							}
						} else if isNumber(formatWord(&(words[i+8])[0])) {
							if !EctsMarks[formatWord(&(words[i+7])[0])] {
								return nil, errors.New("у студента що складав роботу має бути оцінка ЄКТС ")
							} else {
								stud.EctsMark = formatWord(&(words[i+7])[0])
								if !checkMarkAccordance(stud.EctsMark, stud.NationalMark, stud.TogetherMark) {
									return nil, errors.New("перевірте у студента [" + stud.RecordBook + "] оцінку національну, єктс та разом ")
								}
							}
						}

						numPresent += 1
					} else {
						return nil, errors.New("error with student [" + stud.RecordBook + "]")
					}
				} else {
					return nil, errors.New("відомість не відповідає потрібному формату ")
				}

				//now we have to get students full name (check if he has middlename)
				isNumberIfStudentHasMiddleName := formatWord(&(words[i-6])[0])
				_, err := strconv.Atoi(isNumberIfStudentHasMiddleName)
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
	s.EducationalYear = fmt.Sprint(s.Date.Year())
	s.ExtractedStudents = allStudInfo
	return &s, nil
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

func removeEmptyWords(arr *[][]string) [][]string {
	var result [][]string

	for _, v := range *arr {
		s1 := v[1]
		if s1 != "" {
			result = append(result, v)
		}
	}

	return result
}

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}

func checkMarkAccordance(ects string, national string, together int) bool {
	if (ects == "A" || ects == "А") && (national == "відмінно" || national == "зараховано") && (together > 90 && together < 101) {
		return true
	}
	if (ects == "B" || ects == "В") && (national == "добре" || national == "зараховано") && (together > 80 && together < 91) {
		return true
	}
	if (ects == "C" || ects == "С") && (national == "добре" || national == "зараховано") && (together > 70 && together < 81) {
		return true
	}
	if ects == "D" && (national == "задовільно" || national == "зараховано") && (together > 65 && together < 71) {
		return true
	}
	if (ects == "E" || ects == "Е") && national == "задовільно" && together >= 61 && together <= 65 {
		return true
	}
	if (ects == "E" || ects == "Е") && national == "зараховано" && together >= 60 && together <= 65 {
		return true
	}
	if ects == "F" && national == "незадовільно" && together >= 0 && together <= 60 {
		return true
	}
	if ects == "F" && national == "не зараховано" && together >= 0 && together <= 59 {
		return true
	}

	return false
}
