package integer

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(4, 3)
	expected := 7

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
