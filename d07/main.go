package main

import (
	"fmt"
	"strings"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d07(inputs.D07, false))
	fmt.Println(d07(inputs.D07, true))
}

func d07(input string, concat bool) int {
	var testValueStr, numbersStr string
	var ok bool
	var testValue int
	var numbers []int
	var res int

	for _, line := range inputs.Lines(input) {
		testValueStr, numbersStr, ok = strings.Cut(line, ": ")
		if !ok {
			panic("parse error")
		}
		testValue = sugar.MustAtoi(testValueStr)

		numbers = inputs.Numbers(numbersStr, " ")

		if isCalibrated(testValue, numbers[0], numbers[1:], concat) {
			res += testValue
		}
	}

	return res
}

func isCalibrated(testValue, n int, numbers []int, concat bool) bool {
	if len(numbers) == 0 {
		return testValue == n
	}
	if n > testValue {
		return false
	}
	return isCalibrated(testValue, n+numbers[0], numbers[1:], concat) ||
		isCalibrated(testValue, n*numbers[0], numbers[1:], concat) ||
		concat && isCalibrated(testValue, concatNums(n, numbers[0]), numbers[1:], concat)
}

func concatNums(a, b int) int {
	return sugar.MustAtoi(fmt.Sprintf("%d%d", a, b))
}
