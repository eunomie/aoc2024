package main

import "testing"

func TestD10P1(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	expected := 36

	if v := d10p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD10P2(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	expected := 81

	if v := d10p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
