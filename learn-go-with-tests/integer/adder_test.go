package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(4, 3)
	expected := 7

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
