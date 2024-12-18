package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	//fmt.Println(d18p1(inputs.D18, 71, 1024))
	fmt.Println(d18p2(inputs.D18, 71, 1024))
}

func createBoardInput(size int) string {
	var res []string
	for i := 0; i < size; i++ {
		res = append(res, strings.Repeat(".", size))
	}
	return strings.Join(res, "\n")
}

func d18p1(input string, size, nbBytes int) int {
	inputBoard := createBoardInput(size)
	board := sugar.NewBoard(inputBoard)
	target := board.GetRune(size-1, size-1)

	coordsReg := regexp.MustCompile(`^(\d+),(\d+)$`)
	var matches []string
	var x, y int
	for i, line := range inputs.Lines(input) {
		if i == nbBytes {
			break
		}
		matches = coordsReg.FindStringSubmatch(line)
		x = sugar.MustAtoi(matches[1])
		y = sugar.MustAtoi(matches[2])
		board.GetRune(x, y).Write(wall)
	}

	return m(board, 0, target.Index())
}

func d18p2(input string, size, nbBytes int) string {
	inputBoard := createBoardInput(size)
	board := sugar.NewBoard(inputBoard)
	target := board.GetRune(size-1, size-1)

	coordsReg := regexp.MustCompile(`^(\d+),(\d+)$`)
	var matches []string
	var x, y int
	bytes := inputs.Lines(input)
	for i, line := range bytes {
		if i == nbBytes {
			break
		}
		matches = coordsReg.FindStringSubmatch(line)
		x = sugar.MustAtoi(matches[1])
		y = sugar.MustAtoi(matches[2])
		board.GetRune(x, y).Write(wall)
	}

	for i := nbBytes; i < len(bytes); i++ {
		matches = coordsReg.FindStringSubmatch(bytes[i])
		x = sugar.MustAtoi(matches[1])
		y = sugar.MustAtoi(matches[2])
		board.GetRune(x, y).Write(wall)
		if m(board, 0, target.Index()) == math.MaxInt32 {
			return fmt.Sprintf("%v,%v", x, y)
		}
	}

	return ""
}

type Item struct {
	r       sugar.Rune
	visited []int
}

const (
	wall = '#'
)

func m(board *sugar.Board, start, target int) int {
	ms := math.MaxInt32
	stack := []Item{{
		board.At(start),
		[]int{},
	}}

	visited := map[int]int{}

	var v []int
	var it int
	var nbValidPaths int

	for {
		it++
		if len(stack) == 0 {
			break
		}

		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if i.r.Index() == target {
			nbValidPaths++
			ms = min(ms, len(i.visited))
			continue
		}

		if i.r.IsEmpty() {
			continue
		}

		if i.r.Is(wall) {
			continue
		}

		if s, ok := visited[i.r.Index()]; ok && s <= len(i.visited) {
			continue
		}
		visited[i.r.Index()] = len(i.visited)

		v = append(i.visited, i.r.Index())

		stack = append(stack, Item{
			i.r.Top(),
			v,
		})
		stack = append(stack, Item{
			i.r.Right(),
			v,
		})
		stack = append(stack, Item{
			i.r.Bottom(),
			v,
		})
		stack = append(stack, Item{
			i.r.Left(),
			v,
		})
	}

	if ms != math.MaxInt32 {
		fmt.Println(ms, "-", it, "iterations", nbValidPaths, "valid paths")
	} else {
		fmt.Println("no valid path")
	}

	return ms
}
