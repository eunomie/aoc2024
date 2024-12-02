package main

import "testing"

func TestD21P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d21p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD21P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d21p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
