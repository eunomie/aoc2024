package main

import (
	"fmt"
	"github.com/eunomie/aoc2024/inputs"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(d01p1(inputs.D01))
}

func d01p1(inputStr string) int {
	var l1, l2 []int
	for _, line := range strings.Split(inputStr, "\n") {
		if line == "" {
			continue
		}
		v1, v2, ok := strings.Cut(line, "   ")
		if !ok {
			continue
		}
		i1, err := strconv.Atoi(v1)
		if err != nil {
			panic(err)
		}
		i2, err := strconv.Atoi(v2)
		if err != nil {
			panic(err)
		}
		l1 = append(l1, i1)
		l2 = append(l2, i2)
	}

	slices.Sort(l1)
	slices.Sort(l2)

	l := len(l1)

	var res int

	for i := 0; i < l; i++ {
		res += abs(l2[i] - l1[i])
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
