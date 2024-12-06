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
			nextPos.Write(guardUp)
			gPos.Write(empty)
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
			nextPos.Write(guardRight)
			gPos.Write(empty)
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
			nextPos.Write(guardLeft)
			gPos.Write(empty)
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
			nextPos.Write(guardUp)
			gPos.Write(empty)
			gPos = nextPos
		}
	}
END:

	return visitedPositions
}

func d06p2(input string) int {
	allPos := positions(input)
	l := len(allPos)

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
	var i atomic.Int32

	eg := errgroup.Group{}
	eg.SetLimit(100)

	for p, _ := range allPos {
		if p == initialPosition {
			i.Add(1)
			fmt.Println(i.Load(), "/", l, "(skipped)")
			continue
		}
		eg.Go(func() error {
			i.Add(1)
			fmt.Println(i.Load(), "/", l)
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
					if nextPos.Is(obstacle, obstruction) {
						dir = DirRight
						gPos.Write(guardRight)
					} else {
						nextPos.Write(guardUp)
						gPos.Write(empty)
						gPos = nextPos
					}
				case DirRight:
					nextPos := gPos.Right()
					if nextPos.IsEmpty() {
						goto END
					}
					if nextPos.Is(obstacle, obstruction) {
						dir = DirDown
						gPos.Write(guardDown)
					} else {
						nextPos.Write(guardRight)
						gPos.Write(empty)
						gPos = nextPos
					}
				case DirDown:
					nextPos := gPos.Bottom()
					if nextPos.IsEmpty() {
						goto END
					}
					if nextPos.Is(obstacle, obstruction) {
						dir = DirLeft
						gPos.Write(guardLeft)
					} else {
						nextPos.Write(guardDown)
						gPos.Write(empty)
						gPos = nextPos
					}
				case DirLeft:
					nextPos := gPos.Left()
					if nextPos.IsEmpty() {
						goto END
					}
					if nextPos.Is(obstacle, obstruction) {
						dir = DirUp
						gPos.Write(guardUp)
					} else {
						nextPos.Write(guardLeft)
						gPos.Write(empty)
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
