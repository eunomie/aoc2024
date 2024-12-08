package main

import (
	"fmt"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d08p1(inputs.D08))
	fmt.Println(d08p2(inputs.D08))
}

func d08p1(input string) int {
	board := sugar.NewBoard(input)

	usedPositions := map[rune][]sugar.Rune{}
	var positions []sugar.Rune
	var ok bool

	board.ForEachRunesNot(func(r sugar.Rune) {
		positions, ok = usedPositions[r.Rune()]
		if !ok {
			positions = []sugar.Rune{r}
		} else {
			positions = append(positions, r)
		}
		usedPositions[r.Rune()] = positions
	}, '.')

	var antennaPositions = map[int]bool{}
	var l, i, j, dc, dl int
	var a sugar.Rune

	for _, positions = range usedPositions {
		l = len(positions)
		for i = 0; i < l-1; i++ {
			for j = i + 1; j < l; j++ {
				dc, dl = positions[i].Diff(positions[j])
				a = positions[i].Move(-dc, -dl)
				if !a.IsEmpty() {
					antennaPositions[a.Index()] = true
				}
				a = positions[j].Move(dc, dl)
				if !a.IsEmpty() {
					antennaPositions[a.Index()] = true
				}
			}
		}
	}

	return len(antennaPositions)
}

func d08p2(input string) int {
	board := sugar.NewBoard(input)

	usedPositions := map[rune][]sugar.Rune{}
	var positions []sugar.Rune
	var ok bool
	var antennaPositions = map[int]bool{}

	board.ForEachRunesNot(func(r sugar.Rune) {
		positions, ok = usedPositions[r.Rune()]
		if !ok {
			positions = []sugar.Rune{r}
		} else {
			positions = append(positions, r)
		}
		usedPositions[r.Rune()] = positions
		antennaPositions[r.Index()] = true
	}, '.')

	var l, i, j, dc, dl int
	var a sugar.Rune

	for _, positions = range usedPositions {
		l = len(positions)
		for i = 0; i < l-1; i++ {
			for j = i + 1; j < l; j++ {
				dc, dl = positions[i].Diff(positions[j])
				a = positions[i]
				for {
					a = a.Move(-dc, -dl)
					if a.IsEmpty() {
						break
					}
					antennaPositions[a.Index()] = true
				}
				a = positions[j]
				for {
					a = a.Move(dc, dl)
					if a.IsEmpty() {
						break
					}
					antennaPositions[a.Index()] = true
				}
			}
		}
	}

	return len(antennaPositions)
}
