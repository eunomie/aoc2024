package main

import "testing"

func TestD24P1(t *testing.T) {
	input := ``

	expected := 0

	if v := d24p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD24P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d24p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
