package main

import "testing"

func TestD19P1(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	expected := 6

	if v := d19p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD19P2(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	expected := 16

	if v := d19p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
