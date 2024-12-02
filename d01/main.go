package main

import (
	"fmt"
	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
	"slices"
)

func main() {
	fmt.Println(d01p1(inputs.D01))
	fmt.Println(d01p2(inputs.D01))
}

func d01p1(inputStr string) int {
	var l1, l2 []int
	lines := inputs.Lines(inputStr)
	for _, line := range lines {
		values := inputs.Numbers(line)
		i1, i2 := values[0], values[1]
		l1 = append(l1, i1)
		l2 = append(l2, i2)
	}

	slices.Sort(l1)
	slices.Sort(l2)

	l := len(l1)

	var res int

	for i := 0; i < l; i++ {
		res += sugar.Abs(l2[i] - l1[i])
	}

	return res
}

func d01p2(inputStr string) int {
	var l1 []int
	l2 := map[int]int{}
	lines := inputs.Lines(inputStr)
	for _, line := range lines {
		values := inputs.Numbers(line)
		i1, i2 := values[0], values[1]
		l1 = append(l1, i1)
		l2[i2]++
	}

	var res int

	for _, v := range l1 {
		if nb, ok := l2[v]; ok {
			res += v * nb
		}
	}

	return res
}
