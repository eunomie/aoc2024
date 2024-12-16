package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d16p1(inputs.D16))
	fmt.Println(d16p2(inputs.D16))
}

type direction = int

const (
	wall  = '#'
	start = 'S'
	end   = 'E'

	east direction = iota
	west
	north
	south
)

func d16p1(input string) int {
	board := sugar.NewBoard(input)
	reindeer := board.Find(start)

	return m(Pos{reindeer, east})
}

func moveForward(r sugar.Rune, dir direction) sugar.Rune {
	switch dir {
	case north:
		return r.Top()
	case south:
		return r.Bottom()
	case east:
		return r.Right()
	case west:
		return r.Left()
	default:
		panic("unreachable")
	}
}

func turnCW(dir direction) direction {
	switch dir {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	default:
		panic("unreachable")
	}
}

func turnCCW(dir direction) direction {
	switch dir {
	case north:
		return west
	case west:
		return south
	case south:
		return east
	case east:
		return north
	default:
		panic("unreachable")
	}
}

type Pos struct {
	r   sugar.Rune
	dir direction
}

func (p Pos) IsWall() bool {
	return p.r.Is(wall)
}

func (p Pos) IsEnd() bool {
	return p.r.Is(end)
}

func (p Pos) MoveForward() Pos {
	return Pos{
		moveForward(p.r, p.dir),
		p.dir,
	}
}

func (p Pos) TurnCW() Pos {
	dir := turnCW(p.dir)
	return Pos{
		p.r,
		dir,
	}
}

func (p Pos) MoveCW() Pos {
	return p.TurnCW().MoveForward()
}

func (p Pos) TurnCCW() Pos {
	dir := turnCCW(p.dir)
	return Pos{
		p.r,
		dir,
	}
}

func (p Pos) MoveCCW() Pos {
	return p.TurnCCW().MoveForward()
}

func (p Pos) Index() int {
	return p.r.Index()
}

func (p Pos) Key() string {
	return fmt.Sprintf("%v;%v", p.dir, p.Index())
}

type Item struct {
	pos     Pos
	visited []int
	score   int
}

func m(pos Pos) int {
	ms := math.MaxInt32
	stack := []Item{{
		pos,
		[]int{},
		0,
	}}

	visited := map[string]int{}

	for {
		if len(stack) == 0 {
			break
		}

		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if i.score > ms {
			continue
		}
		if i.pos.IsWall() {
			continue
		}

		if slices.Contains(i.visited, i.pos.Index()) {
			continue
		}

		if i.pos.IsEnd() {
			ms = min(ms, i.score)
			continue
		}

		if s, ok := visited[i.pos.Key()]; ok && s < i.score {
			continue
		}
		visited[i.pos.Key()] = i.score

		stack = append(stack, Item{
			i.pos.MoveForward(),
			append(i.visited, i.pos.Index()),
			i.score + 1,
		})
		stack = append(stack, Item{
			i.pos.MoveCW(),
			append(i.visited, i.pos.Index()),
			i.score + 1001,
		})
		stack = append(stack, Item{
			i.pos.MoveCCW(),
			append(i.visited, i.pos.Index()),
			i.score + 1001,
		})
	}

	return ms
}

func d16p2(input string) int {
	board := sugar.NewBoard(input)
	reindeer := board.Find(start)

	return m2(Pos{reindeer, east})
}

func m2(pos Pos) int {
	ms := math.MaxInt32
	tiles := map[int]bool{}
	stack := []Item{{
		pos,
		[]int{},
		0,
	}}

	visited := map[string]int{}

	for {
		if len(stack) == 0 {
			break
		}

		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if i.score > ms {
			continue
		}
		if i.pos.IsWall() {
			continue
		}

		if slices.Contains(i.visited, i.pos.Index()) {
			continue
		}

		if i.pos.IsEnd() {
			if i.score < ms {
				tiles = map[int]bool{}
			}
			if i.score <= ms {
				for _, v := range i.visited {
					tiles[v] = true
				}
				tiles[i.pos.Index()] = true
				ms = i.score
			}
			continue
		}

		if s, ok := visited[i.pos.Key()]; ok && s < i.score {
			continue
		}
		visited[i.pos.Key()] = i.score

		stack = append(stack, Item{
			i.pos.MoveForward(),
			append(i.visited, i.pos.Index()),
			i.score + 1,
		})
		stack = append(stack, Item{
			i.pos.MoveCW(),
			append(i.visited, i.pos.Index()),
			i.score + 1001,
		})
		stack = append(stack, Item{
			i.pos.MoveCCW(),
			append(i.visited, i.pos.Index()),
			i.score + 1001,
		})
	}

	return len(tiles)
}
