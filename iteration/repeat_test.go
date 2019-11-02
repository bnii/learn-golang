package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assert := func(expected, actual string, t *testing.T) {
		t.Helper()
		if expected != actual {
			t.Errorf("expected %q got %q", expected, actual)
		}
	}

	t.Run("five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assert(expected, repeated, t)

	})

	t.Run("0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assert(expected, repeated, t)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}
