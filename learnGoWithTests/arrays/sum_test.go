package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Calculate sum of N integer array", func(t *testing.T) {

		numbers := []int{1, 2, 3}
		actual := Sum(numbers)
		expected := 6

		assertCorrectSum(t, actual, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Calculate sums of slices", func(t *testing.T) {

		x := []int{1, 2, 3}
		y := []int{1, 2, 3, 4, 5}
		z := []int{1, 2}
		actual := SumAll(x, y, z)
		expected := []int{6, 15, 3}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v want %v", actual, expected)
		}

	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(actual []int, expected []int, t *testing.T) {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v want %v", actual, expected)
		}
	}
	t.Run("Calculate sums of slice tails", func(t *testing.T) {

		x := []int{1, 2, 3}
		y := []int{1, 2, 3, 4, 5}
		z := []int{1, 2}
		actual := SumAllTails(x, y, z)
		expected := []int{5, 14, 2}

		checkSums(actual, expected, t)

	})

	t.Run("Saftely calculate sums of empty slice tails", func(t *testing.T) {

		x := []int{}
		y := []int{1, 2, 3, 4, 5}
		z := []int{1, 2}
		actual := SumAllTails(x, y, z)
		expected := []int{0, 14, 2}

		checkSums(actual, expected, t)

	})
}

func assertCorrectSum(t testing.TB, actual, expected int, numbers []int) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected '%d' but got '%d' given, '%v' ", expected, actual, numbers)

	}
}
