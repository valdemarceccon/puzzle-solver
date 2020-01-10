package solver_test

import (
	"awesomeProject/solver"
	"testing"
)

func TestNewSudoku(t *testing.T) {
	givenSolved := "957613284" +
		"483257196" +
		"612849537" +
		"178364952" +
		"524971368" +
		"369528741" +
		"845792613" +
		"291436875" +
		"736185429"

	givenUnsolved := "000000200" +
		"080007090" +
		"602000500" +
		"070060000" +
		"000901000" +
		"000020040" +
		"005000603" +
		"090400070" +
		"006000000"

	_ = givenUnsolved

	t.Run("should parse correct string", func(t *testing.T) {
		sudoku, err := solver.NewSudoku(givenSolved)
		if err != nil || sudoku == nil {
			t.Fatal("should not have returned an error, and the parse should have be done")
		}
	})

	t.Run("should not parse incorrect string", func(t *testing.T) {
		given := givenSolved + "1135456454"

		_, err := solver.NewSudoku(given)
		if err == nil {
			t.Fatal("should have returned an error")
		}
	})
}

func TestSolveSudoku(t *testing.T) {

}
