package main

import "testing"

func TestD22P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d22p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD22P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d22p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
