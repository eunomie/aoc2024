package main

import "testing"

func TestD01P2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected := 31

	if v := d01p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
