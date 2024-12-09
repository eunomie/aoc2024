package main

import "testing"

func TestD09P1(t *testing.T) {
	input := `2333133121414131402`

	expected := 1928

	if v := d09p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD09P2(t *testing.T) {
	input := `2333133121414131402`

	expected := 2858

	if v := d09p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
