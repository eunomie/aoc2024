package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d14p1(inputs.D14, 100, 101, 103))
	fmt.Println(d14p2(inputs.D14, 101, 103))
}

type Position struct {
	X, Y int
}

type Velocity struct {
	X, Y int
}

type Robot struct {
	Current  Position
	Velocity Velocity
}

func d14p1(input string, iterations, width, height int) int {
	lineReg := regexp.MustCompile(`^p=(\d+),(\d+) v=(-?\d+),(-?\d+)$`)
	var rx, ry /*, vx, vy*/ int
	var quadtl, quadtr, quadbl, quadbr int
	var matches []string
	for _, line := range inputs.Lines(input) {
		matches = lineReg.FindStringSubmatch(line)
		rx = (sugar.MustAtoi(matches[1]) + iterations*sugar.MustAtoi(matches[3])) % width
		if rx < 0 {
			rx += width
		}
		ry = (sugar.MustAtoi(matches[2]) + iterations*sugar.MustAtoi(matches[4])) % height
		if ry < 0 {
			ry += height
		}
		if rx < width/2 && ry < height/2 {
			quadtl++
		} else if rx > width/2 && ry > height/2 {
			quadbr++
		} else if rx < width/2 && ry > height/2 {
			quadbl++
		} else if rx > width/2 && ry < height/2 {
			quadtr++
		}
	}

	return quadtl * quadtr * quadbl * quadbr
}

func d14p2(input string, width, height int) int {
	runes := []rune{}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			runes = append(runes, ' ')
		}
		runes = append(runes, '\n')
	}
	b := string(runes)
	lineReg := regexp.MustCompile(`^p=(\d+),(\d+) v=(-?\d+),(-?\d+)$`)
	robots := []Robot{}
	var robot Robot
	var rx, ry int
	var matches []string
	for _, line := range inputs.Lines(input) {
		matches = lineReg.FindStringSubmatch(line)
		robot = Robot{
			Current: Position{
				X: sugar.MustAtoi(matches[1]),
				Y: sugar.MustAtoi(matches[2]),
			},
			Velocity: Velocity{
				X: sugar.MustAtoi(matches[3]),
				Y: sugar.MustAtoi(matches[4]),
			},
		}
		robots = append(robots, robot)
	}

	iterations := width * height
	for i := 0; i < iterations; i++ {
		board := sugar.NewBoard(b)

		for j := 0; j < len(robots); j++ {
			robot = robots[j]
			rx = (robot.Current.X + /*iterations**/ robot.Velocity.X) % width
			if rx < 0 {
				rx += width
			}
			ry = (robot.Current.Y + /*iterations**/ robot.Velocity.Y) % height
			if ry < 0 {
				ry += height
			}
			robot.Current.X = rx
			robot.Current.Y = ry
			robots[j] = robot

			board.GetRune(robot.Current.X, robot.Current.Y).Write('*')
		}
		fmt.Println(board)
		fmt.Println(i, "/", iterations)
		fmt.Println()
		time.Sleep(20 * time.Millisecond)
	}

	return 0
}
