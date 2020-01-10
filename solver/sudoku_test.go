package solver_test

import (
	"awesomeProject/solver"
	"testing"
)

func TestSudoku(t *testing.T) {
	t.Run("should parse correct string", func(t *testing.T) {
		given := "123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789"

		_, err := solver.NewSudoku(given)
		if err != nil {
			t.Fatal("should not have returned an error")
		}
	})

	t.Run("should not parse incorrect string", func(t *testing.T) {
		given := "123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789" +
			"123456789111"

		_, err := solver.NewSudoku(given)
		if err == nil {
			t.Fatal("should have returned an error")
		}
	})
}
