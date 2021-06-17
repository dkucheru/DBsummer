package pdfReading

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ReadInfoAboutStudentWithNoRecordBook(words [][]string, i int, numPresent *int, numNotAllowed *int) (*StudInfoFromPDF, error) {

	var stud StudInfoFromPDF
	c := 0 //middlename case

	//recordbook info
	stud.RecordBook = ""

	i_plus_1 := strings.ToLower(formatWord(&(words[i+1])[0]))
	i_plus_2 := strings.ToLower(formatWord(&(words[i+2])[0]))
	i_plus_3 := strings.ToLower(formatWord(&(words[i+3])[0]))

	//parse name
	if !isNumber(i_plus_1) {
		stud.Lastname = i_plus_1
	}
	if !isNumber(i_plus_2) {
		stud.FirstName = i_plus_2
	}
	if !isNumber(i_plus_3) {
		stud.MiddleName = i_plus_3
		c = 1
	}

	i += 2 + c
	//helping variables
	i_plus_1 = strings.ToLower(formatWord(&(words[i+1])[0]))
	i_plus_2 = strings.ToLower(formatWord(&(words[i+2])[0]))
	i_plus_3 = strings.ToLower(formatWord(&(words[i+3])[0]))
	i_plus_4 := strings.ToLower(formatWord(&(words[i+4])[0]))
	i_plus_5 := strings.ToLower(formatWord(&(words[i+5])[0]))
	i_plus_6 := strings.ToLower(formatWord(&(words[i+6])[0]))

	if strings.ToLower(formatWord(&(words[i+1])[0])) == "не" &&
		strings.Contains(i_plus_2, "відв") {
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
		*numNotAllowed += 1

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
		*numNotAllowed += 1
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
		*numNotAllowed += 1
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
			*numPresent += 1
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
			*numPresent += 1
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
			*numPresent += 1
		} else {
			return nil, errors.New("error with student [" + stud.RecordBook + "]")
		}
	} else {
		return nil, errors.New("відомість не відповідає потрібному формату [no record book] ")
	}

	return &stud, nil
}
