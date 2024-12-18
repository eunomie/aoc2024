package main

import (
	"strconv"
	"testing"
)

func TestD17P1(t *testing.T) {
	testCase := []struct {
		input    string
		expected string
	}{
		{
			`Register A: 10
Register B: 0
Register C: 0

Program: 5,0,5,1,5,4`,
			"0,1,2",
		},
		{
			`Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
			"4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
			"4,6,3,5,6,3,5,2,1,0",
		},
	}

	for i, tc := range testCase {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if v := d17p1(tc.input); v != tc.expected {
				t.Errorf("expcted %v, got %v", tc.expected, v)
			}
		})
	}
}

func TestD17P2(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

	expected := 117440

	if v := d17p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestA(t *testing.T) {
	a := 7 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8
	b := 1 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8 * 8
	//a := div2(7, 3)
	if a != 0 {
		t.Errorf("expcted %v, got %v", b, a)
	}
}
