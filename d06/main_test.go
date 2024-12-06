package main

import "testing"

func TestD06P1(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expected := 41

	if v := d06p1(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD06P2(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expected := 6

	if v := d06p2(input); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
