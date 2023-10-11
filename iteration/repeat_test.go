package iteration

import "testing"

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	result := Repeat("a", 5)

	if result != expected {
		t.Errorf("expected: %q result %q", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
