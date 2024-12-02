package main

import "testing"

func TestD02P1(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	expected := 2

	if v := d02(input, false); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD02P2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	expected := 4

	if v := d02(input, true); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
