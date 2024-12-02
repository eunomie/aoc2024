package main

import "testing"

func TestD10P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d10p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD10P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d10p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
