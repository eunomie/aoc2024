package main

import "testing"

func TestD15P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d15p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD15P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d15p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
