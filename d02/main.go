package main

import (
	"fmt"
	"slices"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d02(inputs.D02, false))
	fmt.Println(d02(inputs.D02, true))
}

func d02(inputStr string, skip bool) int {
	safeReports := 0
	reports := inputs.Lines(inputStr)
	for _, report := range reports {
		levels := inputs.Numbers(report, " ")
		safe := isSafe(levels)
		if safe {
			safeReports++
		} else if skip {
			for i := 0; i < len(levels); i++ {
				safe = isSafe(sugar.FilterIndex(levels, i))
				if safe {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}

func isSafe(levels []int) bool {
	if !slices.IsSorted(levels) && !slices.IsSorted(sugar.Reverse(levels)) {
		return false
	}
	n := levels[0]
	for i := 1; i < len(levels); i++ {
		v := levels[i]
		d := sugar.Abs(v - n)
		if d < 1 || d > 3 {
			return false
		}
		n = v
	}
	return true
}
