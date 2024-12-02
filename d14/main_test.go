package main

import "testing"

func TestD14P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d14p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD14P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d14p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
