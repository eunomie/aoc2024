package main

import "testing"

func TestD19P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d19p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD19P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d19p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
