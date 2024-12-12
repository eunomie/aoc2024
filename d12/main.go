package main

import (
	"fmt"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d12p1(inputs.D12))
	fmt.Println(d12p2(inputs.D12))
}

func d12p1(input string) int {
	price := 0
	visitedPositions := map[int]struct{}{}
	board := sugar.NewBoard(input)
	board.ForEachRunes(func(r sugar.Rune) {
		if _, visited := visitedPositions[r.Index()]; visited {
			return
		}
		positions, area, perimeter := findRegion(r)
		for _, p := range positions {
			visitedPositions[p] = struct{}{}
		}
		price += area * perimeter
	})

	return price
}

func findRegion(r sugar.Rune) (positions []int, area int, perimeter int) {
	plant := r.Rune()
	var plot sugar.Rune
	plots := []sugar.Rune{r}
	var p sugar.Rune
	visitedPositions := map[int]struct{}{}
	for {
		if len(plots) == 0 {
			break
		}
		plot = plots[0]
		plots = plots[1:]
		if plot.IsEmpty() {
			perimeter++
			continue
		}
		if !plot.Is(plant) {
			perimeter++
			continue
		}
		if _, visited := visitedPositions[plot.Index()]; visited {
			continue
		}
		visitedPositions[plot.Index()] = struct{}{}

		plot.Write('.')
		positions = append(positions, plot.Index())
		area++

		for _, p = range []sugar.Rune{plot.Top(), plot.Bottom(), plot.Left(), plot.Right()} {
			if p.IsEmpty() {
				plots = append(plots, p)
				continue
			}
			if _, visited := visitedPositions[p.Index()]; !visited {
				plots = append(plots, p)
			}
		}
	}

	return
}

const plant = 'p'

func d12p2(input string) int {
	refBoard := sugar.NewBoard(input)
	price := 0
	visitedPositions := map[int]struct{}{}
	board := sugar.NewBoard(input)
	board.ForEachRunes(func(r sugar.Rune) {
		if _, visited := visitedPositions[r.Index()]; visited {
			return
		}
		positions, area, _ := findRegion(r)
		for _, p := range positions {
			visitedPositions[p] = struct{}{}
		}
		isolated := refBoard.IsolateWithRune(positions, plant)
		corners := findCorners(isolated)
		price += area * corners
	})

	return price
}

func findCorners(isolated *sugar.Board) int {
	corners := 0
	var l, t, r, b, tl, tr, br, bl bool
	isolated.ForEachRunes(func(ru sugar.Rune) {
		l, t, r, b = ru.Left().Is(plant), ru.Top().Is(plant), ru.Right().Is(plant), ru.Bottom().Is(plant)
		tl, tr, br, bl = ru.TopLeft().Is(plant), ru.TopRight().Is(plant), ru.BottomRight().Is(plant), ru.BottomLeft().Is(plant)
		if !l && !t && !tl {
			corners++
		}
		if !r && !t && !tr {
			corners++
		}
		if !r && !b && !br {
			corners++
		}
		if !l && !b && !bl {
			corners++
		}
		if l && t && !tl {
			corners++
		}
		if r && t && !tr {
			corners++
		}
		if r && b && !br {
			corners++
		}
		if l && b && !bl {
			corners++
		}
		if !l && !t && tl {
			corners++
		}
		if !r && !t && tr {
			corners++
		}
		if !r && !b && br {
			corners++
		}
		if !l && !b && bl {
			corners++
		}
	}, plant)
	return corners
}
