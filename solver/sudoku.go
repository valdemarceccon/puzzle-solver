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

	result.parse(cleanedString)

	if !result.IsValid() {
		return nil, errors.New("puzzle is not in a valid state")
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

func (s *Sudoku) parse(values string) {
	for row := 0; row < 9; row ++ {
		for col := 0; col < 9; col ++ {
			n, _ := strconv.Atoi(string(values[row*col + col]))
			s.tab[row][col] = uint8(n)
		}
	}
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

func (s *Sudoku) IsValid() bool {
	for i:= 0; i < 9; i++ {
		if !isLineValid(i) || !isColumnValid(i) || !isQuadrantValid(i) {
			return false
		}
	}

	return true
}

func isQuadrantValid(i int) bool {

	return true
}

func isColumnValid(i int) bool {

	return true
}

func isLineValid(i int) bool {

	return true
}

func (s *Sudoku) getColumn(col int) []uint8 {
	result := make([]uint8, 9)

	for linha := 0; linha < 9; linha++ {
		result[linha] = s.tab[linha][col]
	}

	return result
}

func (s *Sudoku) getRow(row int) []uint8 {
	result := make([]uint8, 9)

	for col := 0; col < 9; col++ {
		result[col] = s.tab[row][col]
	}

	return result
}

func (s *Sudoku) getQuadrant(q int) []uint8 {
	result := make([]uint8, 9)

	for quad := 0; quad < 9; quad++ {
		result[quad] = s.tab[q][quad]
	}

	return result
}
