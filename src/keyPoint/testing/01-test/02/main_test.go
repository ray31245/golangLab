package main

import "testing"

func TestAverage(t *testing.T) {
	a := average(1, 2, 3)
	if a != 2 {
		t.Error("the average should be ", 2, "Got", a)
	}
}
