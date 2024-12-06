package main

import (
	"fmt"
	"sync/atomic"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println(d06p1(inputs.D06))
	fmt.Println(d06p2(inputs.D06))
}

const (
	guardUp     = '^'
	guardRight  = '>'
	guardDown   = 'v'
	guardLeft   = '<'
	obstacle    = '#'
	empty       = '.'
	obstruction = 'O'
)

type Direction int

const (
	DirUp Direction = iota
	DirRight
	DirDown
	DirLeft
)

func d06p1(input string) int {
	return len(positions(input))
}

func positions(input string) map[int]bool {
	visitedPositions := map[int]bool{}

	board := sugar.NewBoard(input)

	var gPos sugar.Rune
	var dir Direction
	if gPos = board.Find(guardUp); !gPos.IsEmpty() {
		dir = DirUp
	} else if gPos = board.Find(guardRight); !gPos.IsEmpty() {
		dir = DirRight
	} else if gPos = board.Find(guardDown); !gPos.IsEmpty() {
		dir = DirDown
	} else if gPos = board.Find(guardLeft); !gPos.IsEmpty() {
		dir = DirLeft
	} else {
		panic("guard not found")
	}
	visitedPositions[gPos.Index()] = true

	for {
		switch dir {
		case DirUp:
			nextPos := gPos.Top()
			if nextPos.IsEmpty() {
				goto END
			}
			if nextPos.Is(obstacle) {
				dir = DirRight
				continue
			}
			visitedPositions[nextPos.Index()] = true
			gPos = nextPos
		case DirRight:
			nextPos := gPos.Right()
			if nextPos.IsEmpty() {
				goto END
			}
			if nextPos.Is(obstacle) {
				dir = DirDown
				continue
			}
			visitedPositions[nextPos.Index()] = true
			gPos = nextPos
		case DirDown:
			nextPos := gPos.Bottom()
			if nextPos.IsEmpty() {
				goto END
			}
			if nextPos.Is(obstacle) {
				dir = DirLeft
				continue
			}
			visitedPositions[nextPos.Index()] = true
			gPos = nextPos
		case DirLeft:
			nextPos := gPos.Left()
			if nextPos.IsEmpty() {
				goto END
			}
			if nextPos.Is(obstacle) {
				dir = DirUp
				continue
			}
			visitedPositions[nextPos.Index()] = true
			gPos = nextPos
		}
	}
END:

	return visitedPositions
}

func d06p2(input string) int {
	allPos := positions(input)

	initialBoard := sugar.NewBoard(input)
	var initialRune sugar.Rune
	var initialDir Direction
	if initialRune = initialBoard.Find(guardUp); !initialRune.IsEmpty() {
		initialDir = DirUp
	} else if initialRune = initialBoard.Find(guardRight); !initialRune.IsEmpty() {
		initialDir = DirRight
	} else if initialRune = initialBoard.Find(guardDown); !initialRune.IsEmpty() {
		initialDir = DirDown
	} else if initialRune = initialBoard.Find(guardLeft); !initialRune.IsEmpty() {
		initialDir = DirLeft
	} else {
		panic("guard not found")
	}
	initialPosition := initialRune.Index()

	var res atomic.Int32

	eg := errgroup.Group{}
	eg.SetLimit(10)

	for p, _ := range allPos {
		if p == initialPosition {
			continue
		}
		eg.Go(func() error {
			board := sugar.NewBoard(input)
			board.Write(p, obstruction)
			gPos := board.GetRuneAt(initialPosition)
			dir := initialDir
			visitedPositions := map[string]bool{
				fmt.Sprintf("%v%v", initialPosition, initialDir): true,
			}

			for {
				switch dir {
				case DirUp:
					nextPos := gPos.Top()
					if nextPos.IsEmpty() {
						goto END
					}
					if nextPos.Is(obstacle) || nextPos.Index() == p {
						dir = DirRight
					} else {
						gPos = nextPos
					}
				case DirRight:
					nextPos := gPos.Right()
					if nextPos.IsEmpty() {
						goto END
					}
					if nextPos.Is(obstacle) || nextPos.Index() == p {
						dir = DirDown
					} else {
						gPos = nextPos
					}
				case DirDown:
					nextPos := gPos.Bottom()
					if nextPos.IsEmpty() {
						goto END
					}
					if nextPos.Is(obstacle) || nextPos.Index() == p {
						dir = DirLeft
					} else {
						gPos = nextPos
					}
				case DirLeft:
					nextPos := gPos.Left()
					if nextPos.IsEmpty() {
						goto END
					}
					if nextPos.Is(obstacle) || nextPos.Index() == p {
						dir = DirUp
					} else {
						gPos = nextPos
					}
				}
				curPos := fmt.Sprintf("%v%v", gPos.Index(), dir)
				if _, ok := visitedPositions[curPos]; ok {
					res.Add(1)
					break
				}
				visitedPositions[curPos] = true
			}
		END:
			return nil
		})
	}

	_ = eg.Wait()

	return int(res.Load())
}
