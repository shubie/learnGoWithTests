package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 4)
	want := 6

	if sum != want {
		t.Errorf("got %d want %d", sum, want)
	}
}
