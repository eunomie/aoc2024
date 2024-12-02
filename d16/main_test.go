package main

import "testing"

func TestD16P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d16p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD16P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d16p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
