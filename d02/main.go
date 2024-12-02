package main

import (
	"fmt"
	"github.com/eunomie/aoc2024/inputs"
)

func main() {
	fmt.Println(d02(inputs.D02, false))
	fmt.Println(d02(inputs.D02, true))
}

func d02(inputStr string, skip bool) int {
	safeReports := 0
	reports := inputs.Lines(inputStr)
	for _, report := range reports {
		levels := inputs.Numbers(report)
		safe := isSafe(levels)
		if safe {
			safeReports++
		} else if skip {
			for i := 0; i < len(levels); i++ {
				safe = isSafe(arrayMinus(levels, i))
				if safe {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}

func arrayMinus(arr []int, idx int) []int {
	var res []int
	for i, a := range arr {
		if i != idx {
			res = append(res, a)
		}
	}
	return res
}

func isSafe(levels []int) bool {
	type Direction int
	const (
		None Direction = iota
		Inc
		Dec
	)
	safe := true
	direction := None
	n := levels[1]
	d := n - levels[0]
	if d > 0 {
		direction = Inc
	} else {
		direction = Dec
	}
	d = abs(d)
	if d < 1 || d > 3 {
		return false
	}
	for i := 2; i < len(levels); i++ {
		v := levels[i]
		d = v - n
		if d > 0 && direction == Dec || d < 0 && direction == Inc {
			safe = false
			break
		}
		d = abs(d)
		if d < 1 || d > 3 {
			safe = false
			break
		}
		n = v
	}
	return safe
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
