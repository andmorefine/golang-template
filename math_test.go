package main

import (
	"testing"
)

func TestAverage(t *testing.T) {
	got := Average(5, 7)
	wont := 6

	if got != wont {
		t.Errorf("got %q wont %q", got, wont)
	}
}
