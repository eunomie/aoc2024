package main

import (
	"fmt"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d10p1(inputs.D10))
	fmt.Println(d10p2(inputs.D10))
}

func d10p1(input string) int {
	board := sugar.NewBoard(input)

	var scores int

	board.ForEachRunes(func(r sugar.Rune) {
		m := map[int]struct{}{}
		ts := find(board, r, '1')
		for _, s := range ts {
			m[s] = struct{}{}
		}
		scores += len(m)
	}, '0')

	return scores
}

func find(board *sugar.Board, r sugar.Rune, next rune) (scores []int) {
	var n sugar.Rune
	n = r.Top()
	if !n.IsEmpty() && n.Rune() == next {
		if next == '9' {
			scores = append(scores, n.Index())
		} else {
			s := find(board, n, next+1)
			scores = append(scores, s...)
		}
	}
	n = r.Right()
	if !n.IsEmpty() && n.Rune() == next {
		if next == '9' {
			scores = append(scores, n.Index())
		} else {
			s := find(board, n, next+1)
			scores = append(scores, s...)
		}
	}
	n = r.Bottom()
	if !n.IsEmpty() && n.Rune() == next {
		if next == '9' {
			scores = append(scores, n.Index())
		} else {
			s := find(board, n, next+1)
			scores = append(scores, s...)
		}
	}
	n = r.Left()
	if !n.IsEmpty() && n.Rune() == next {
		if next == '9' {
			scores = append(scores, n.Index())
		} else {
			s := find(board, n, next+1)
			scores = append(scores, s...)
		}
	}
	return scores
}

func d10p2(input string) int {
	board := sugar.NewBoard(input)

	var scores int

	board.ForEachRunes(func(r sugar.Rune) {
		ts := find(board, r, '1')
		scores += len(ts)
	}, '0')

	return scores
}
