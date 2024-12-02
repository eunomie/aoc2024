package main

import "testing"

func TestD20P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d20p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD20P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d20p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
