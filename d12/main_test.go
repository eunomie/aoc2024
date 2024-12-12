package main

import (
	"fmt"
	"testing"
)

func TestD12P1(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	expected := 1930

	if v := d12p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD12P2(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			1206,
		},
		{
			`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			236,
		},
		{
			`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			368,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if v := d12p2(tc.input); v != tc.expected {
				t.Errorf("expcted %v, got %v", tc.expected, v)
			}
		})
	}
}
