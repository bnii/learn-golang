package arrays

import "testing"

func TestSum(t *testing.T) {

	assert := func(got, expected int, t *testing.T) {
		t.Helper()
		if got != expected {
			t.Errorf("got %d expected %d", got, expected)
		}
	}

	t.Run("random number woorks", func(t *testing.T) {

		arr := []int{1, 2, 3, 4}
		got := Sum(arr)
		expected := 10

		assert(got, expected, t)
	})

}
