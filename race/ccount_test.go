package main

import "testing"

func TestRace(t *testing.T) {
	res := run()
	if res != 10 {
		t.Errorf("concurrent count wrong: expected 10, got %v", res)
	}
}
