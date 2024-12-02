package main

import "testing"

func TestD13P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d13p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD13P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d13p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
