package main

import "testing"

func TestD03P1(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	expected := 161

	if v := d03p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD03P2(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	expected := 48

	if v := d03p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
