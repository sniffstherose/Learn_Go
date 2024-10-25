package main

import "testing"

// Fuzz开头
func FuzzDiv(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		Div(a, b)
	})
}
