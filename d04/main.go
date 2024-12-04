package main

import (
	"fmt"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d04p1(inputs.D04))
	fmt.Println(d04p2(inputs.D04))
}

func d04p1(input string) int {
	var res int
	const xmas = "XMAS"
	const samx = "SAMX"

	board := sugar.NewBoard(input)

	res += board.CountNonOverlapping(xmas)
	res += board.CountNonOverlapping(samx)
	res += board.CountNonOverlappingInDiagonal(xmas)
	res += board.CountNonOverlappingInDiagonal(samx)

	board = board.Pivot()
	res += board.CountNonOverlapping(xmas)
	res += board.CountNonOverlapping(samx)
	res += board.CountNonOverlappingInDiagonal(xmas)
	res += board.CountNonOverlappingInDiagonal(samx)

	return res
}

func d04p2(input string) int {
	nbCross := 0

	board := sugar.NewBoard(input)
	board.ForEachRunes(func(r sugar.Rune) {
		c1 := fmt.Sprintf("%s%s%s", r.TopLeft(), r, r.BottomRight())
		c2 := fmt.Sprintf("%s%s%s", r.TopRight(), r, r.BottomLeft())
		if sugar.StringEqualOrReverse(c1, "MAS") && sugar.StringEqualOrReverse(c2, "MAS") {
			nbCross++
		}
	}, 'A')

	return nbCross
}
