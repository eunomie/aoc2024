package main

import "testing"

func TestD04P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d04p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD04P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d04p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
