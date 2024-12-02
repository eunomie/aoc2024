package main

import "testing"

func TestD23P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d23p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD23P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d23p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
