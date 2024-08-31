package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat character 5 times", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"
		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})

	t.Run("Repeat character 2 times", func(t *testing.T) {
		actual := Repeat("a", 2)
		expected := "aa"
		if actual != expected {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
