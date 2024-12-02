package main

import "testing"

func TestD18P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d18p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD18P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d18p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
