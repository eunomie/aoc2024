package main

import (
	"fmt"
	"github.com/eunomie/aoc2024/inputs"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(d01p2(inputs.D01))
}

func d01p2(inputStr string) int {
	var l1 []int
	l2 := map[int]int{}
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
