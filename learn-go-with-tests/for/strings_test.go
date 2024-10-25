package iteration

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	type args struct {
		src  string
		find string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Find character in string", args{"hello", "e"}, 1},
		{"Find non-existing character", args{"hello", "z"}, -1},
		{"Empty string", args{"", "a"}, -1},
		{"Find string in string", args{"hello world", "world"}, 6},
		{"Find empty string", args{"hello", ""}, 0},
		{"Find multiple occurrences", args{"banana", "a"}, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := strings.Index(test.args.src, test.args.find)
			if got != test.want {
				t.Errorf("want '%q' got '%q'", test.want, got)
			}
		})
	}
}
