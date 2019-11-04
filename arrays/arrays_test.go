package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assert := func(got, expected int, t *testing.T) {
		t.Helper()
		if got != expected {
			t.Errorf("got %d expected %d", got, expected)
		}
	}

	t.Run("random number works", func(t *testing.T) {

		arr := []int{1, 2, 3, 4}
		got := Sum(arr)
		expected := 10

		assert(got, expected, t)
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5})
	expected := []int{6, 9}

	assert(got, expected, t)

}

func assert(got, expected []int, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {


	t.Run("noormaladd", func(t *testing.T) {

		got := SumAllTails([]int{1, 2, 3}, []int{4, 5})
		expected := []int{5, 5}

		assert(got, expected, t)
	})

	t.Run("emptyslice", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{4, 5})
		expected := []int{0, 5}

		assert(got, expected, t)
	})


}
