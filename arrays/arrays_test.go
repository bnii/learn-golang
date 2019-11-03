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

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}
