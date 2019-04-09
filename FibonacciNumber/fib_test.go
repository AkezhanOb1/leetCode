package main

import "testing"

func TestFib(t *testing.T) {
	total := fib(1)
	if total != 1 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 1)
	}
}
