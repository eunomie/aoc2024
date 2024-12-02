package main

import "testing"

func TestD08P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d08p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD08P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d08p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
