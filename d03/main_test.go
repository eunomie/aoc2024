package main

import "testing"

func TestD03P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d03p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD03P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d03p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
