package main

import "testing"

func TestD12P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d12p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD12P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d12p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
