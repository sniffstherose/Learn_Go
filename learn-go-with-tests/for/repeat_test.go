package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	getExpected := func(t *testing.T, character string, cnt int) string {
		t.Helper()
		var res string
		for i := 0; i < cnt; i++ {
			res += character
		}
		return res
	}

	var cnt int = 6
	var character = "a"
	repeat := Repeat(character, cnt)
	expected := getExpected(t, character, cnt)

	if repeat != expected {
		t.Errorf("expected '%q' repeat '%q'", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	ans := Repeat("a", 5)
	fmt.Println(ans)
	// Output: aaaaa
}
