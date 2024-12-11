package main

import "testing"

func TestD11P1(t *testing.T) {
	input := `125 17`

	expected := 55312

	if v := d11p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD11P2(t *testing.T) {
	input := ``

	expected := 0

	if v := d11p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
