package main

import "testing"

func TestD17P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d17p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD17P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d17p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
