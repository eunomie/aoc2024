package main

import "testing"

func TestD07P1(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	expected := 3749

	if v := d07(input, false); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD07P2(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	expected := 11387

	if v := d07(input, true); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
