package solver

import (
	"errors"
	"strconv"
)

type Sudoku struct {
	tab [9][9]uint8
}

func NewSudoku(values string) (result *Sudoku, err error) {
	result = &Sudoku{}
	cleanedString := getAllNumbers(values)

	if len(cleanedString) != 81 {
		return nil, errors.New("Cannot parse string")
	}

	return
}

func getAllNumbers(values string) string {
	cleanedString := ""

	for _, v := range values {
		_, err := strconv.Atoi(string(v))

		if err == nil {
			cleanedString += string(v)
		}
	}
	return cleanedString
}

func (s *Sudoku) String() string {
	result := ""

	for _, lin := range s.tab {
		for _, col := range lin {
			result = result + strconv.Itoa(int(col)) + " "

		}

		result += "\n"
	}

	return result
}
