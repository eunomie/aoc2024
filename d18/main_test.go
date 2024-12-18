package main

import "testing"

func TestD18P1(t *testing.T) {
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

	expected := 22

	if v := d18p1(input, 7, 12); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}

func TestD18P2(t *testing.T) {
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

	expected := "6,1"

	if v := d18p2(input, 7, 12); v != expected {
		t.Errorf("expcted %v, got %v", expected, v)
	}
}
