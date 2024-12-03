package main

import (
	"fmt"
	"regexp"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d03p1(inputs.D03))
	fmt.Println(d03p2(inputs.D03))
}

func d03p1(input string) int {
	res := 0
	mulRegExp := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	matches := mulRegExp.FindAllStringSubmatch(input, -1)
	for _, v := range matches {
		res += sugar.MustAtoi(v[1]) * sugar.MustAtoi(v[2])
	}
	return res
}

func d03p2(input string) int {
	res := 0
	mulRegExp := regexp.MustCompile("(mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\))")
	matches := mulRegExp.FindAllStringSubmatch(input, -1)
	enabled := true
	for _, v := range matches {
		if v[1] == "do()" {
			enabled = true
		} else if v[1] == "don't()" {
			enabled = false
		} else if enabled {
			res += sugar.MustAtoi(v[2]) * sugar.MustAtoi(v[3])
		}
	}
	return res
}
