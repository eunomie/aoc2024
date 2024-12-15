package main

import (
	"fmt"
	"strings"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d15p1(inputs.D15))
	fmt.Println(d15p2(inputs.D15))
}

type direction = int32

const (
	up    direction = '^'
	right direction = '>'
	down  direction = 'v'
	left  direction = '<'

	robot = '@'
	wall  = '#'
	box   = 'O'
	empty = '.'
	lbox  = '['
	rbox  = ']'
)

func d15p1(input string) int {
	mapw, instructionLines, ok := strings.Cut(input, "\n\n")
	if !ok {
		panic("unreadable input")
	}

	board := sugar.NewBoard(mapw)
	instructions := strings.ReplaceAll(instructionLines, "\n", "")
	r := board.Find(robot)

	var next sugar.Rune

	for _, inst := range instructions {
		switch inst {
		case up:
			next = r.Top()
		case right:
			next = r.Right()
		case down:
			next = r.Bottom()
		case left:
			next = r.Left()
		}
		if next.Is(wall) {
			// ignore it, we can't move walls
		} else if next.Is(box) {
			if tryMove(next, inst) {
				next.Write(robot)
				r.Write(empty)
				r = next
			}
		} else {
			// free to move
			next.Write(robot)
			r.Write(empty)
			r = next
		}
	}

	var coordinates int
	board.ForEachRunes(func(r sugar.Rune) {
		coordinates += r.Col() + 100*r.Line()
	}, box)

	return coordinates
}

func tryMove(next sugar.Rune, dir direction) bool {
	var i int
	var b sugar.Rune
	b = next
	for i = 1; ; i++ {
		switch dir {
		case up:
			b = b.Top()
		case right:
			b = b.Right()
		case down:
			b = b.Bottom()
		case left:
			b = b.Left()
		}
		if b.Is(empty) {
			// it's only necessary to write this one as we are pushing a full set of boxes.
			// OO. become OOO then robot will be written at "next" position @OO
			b.Write(box)
			return true
		}
		if b.Is(wall) {
			return false
		}
	}
}

func d15p2(input string) int {
	mapw, instructionLines, ok := strings.Cut(input, "\n\n")
	if !ok {
		panic("unreadable input")
	}

	nmap := []rune{}
	for _, e := range mapw {
		if e == wall {
			nmap = append(nmap, wall, wall)
		} else if e == box {
			nmap = append(nmap, lbox, rbox)
		} else if e == empty {
			nmap = append(nmap, empty, empty)
		} else if e == robot {
			nmap = append(nmap, robot, empty)
		} else if e == '\n' {
			nmap = append(nmap, '\n')
		}
	}

	board := sugar.NewBoard(string(nmap))

	instructions := strings.ReplaceAll(instructionLines, "\n", "")
	r := board.Find(robot)

	var next sugar.Rune

	for _, inst := range instructions {
		switch inst {
		case up:
			next = r.Top()
		case right:
			next = r.Right()
		case down:
			next = r.Bottom()
		case left:
			next = r.Left()
		}
		if next.Is(wall) {
			// ignore it, we can't move walls
		} else if next.Is(lbox, rbox) {
			if tryMove2(board, next, inst) {
				next.Write(robot)
				r.Write(empty)
				r = next
			}
		} else {
			// free to move
			next.Write(robot)
			r.Write(empty)
			r = next
		}
	}

	var coordinates int
	board.ForEachRunes(func(r sugar.Rune) {
		coordinates += r.Col() + 100*r.Line()
	}, lbox)

	return coordinates
}

func tryMove2(board *sugar.Board, next sugar.Rune, dir direction) bool {
	var i int
	var b sugar.Rune
	b = next
	if dir == right || dir == left {
		for i = 1; ; i++ {
			switch dir {
			case right:
				b = b.Right()
			case left:
				b = b.Left()
			}
			if b.Is(wall) {
				return false
			}
			if b.Is(empty) {
				if dir == right {
					b = next
					for j := 1; j <= i; j += 2 {
						b = b.Right()
						b.Write(lbox)
						b = b.Right()
						b.Write(rbox)
					}
				} else if dir == left {
					b = next
					for j := 1; j <= i; j += 2 {
						b = b.Left()
						b.Write(rbox)
						b = b.Left()
						b.Write(lbox)
					}
				}
				return true
			}
		}
	} else if dir == up || dir == down {
		positions := [][]int{}
		if b.Is(rbox) {
			positions = append(positions, []int{b.Left().Index(), b.Index()})
		} else if b.Is(lbox) {
			positions = append(positions, []int{b.Index(), b.Right().Index()})
		}
		for i = 0; ; i++ {
			pos := []int{}
			for _, p := range positions[i] {
				b = board.At(p)
				if dir == up {
					b = b.Top()
				} else {
					b = b.Bottom()
				}
				if b.Is(wall) {
					return false
				}
				if b.Is(rbox) {
					pos = append(pos, b.Left().Index(), b.Index())
				} else if b.Is(lbox) {
					pos = append(pos, b.Index(), b.Right().Index())
				}
			}
			if len(pos) == 0 {
				break
			}
			positions = append(positions, pos)
		}
		done := map[int]bool{}
		for i = len(positions) - 1; i >= 0; i-- {
			for _, p := range positions[i] {
				if _, ok := done[p]; ok {
					continue
				}
				b = board.At(p)
				if dir == up {
					b.Top().Write(b.Rune())
					b.Write(empty)
				} else {
					b.Bottom().Write(b.Rune())
					b.Write(empty)
				}
				done[p] = true
			}
		}
		return true
	}
	return false
}
