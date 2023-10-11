package integers

import "testing"

func TestAdd(t *testing.T) {
	expected := 4
	sum := Add(2, 2)

	if sum != expected {
		t.Errorf("expected: %d result: %d", expected, sum)
	}
}
