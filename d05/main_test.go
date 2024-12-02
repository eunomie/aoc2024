package main

import "testing"

func TestD05P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d05p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD05P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d05p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
