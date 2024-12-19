package main

import (
	"fmt"
	"strings"

	"github.com/eunomie/aoc2024/inputs"
)

func main() {
	fmt.Println(d19p1(inputs.D19))
	fmt.Println(d19p2(inputs.D19))
}

func d19p1(input string) int {
	return len(possibleDesigns(input))
}

func possibleDesigns(input string) []int {
	lines := inputs.Lines(input)
	patterns := strings.Split(lines[0], ", ")

	var possibles []int
	designs := lines[2:]

	for i, design := range designs {
		stack := []string{design}

		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, pattern := range patterns {
				if strings.HasPrefix(current, pattern) {
					if len(current) == len(pattern) {
						// we are at the end
						possibles = append(possibles, i)
						goto NEXT
					}
					stack = append(stack, current[len(pattern):])
				}
			}
		}
	NEXT:
	}

	return possibles
}

func d19p2(input string) int {
	lines := inputs.Lines(input)
	patterns := strings.Split(lines[0], ", ")

	designs := lines[2:]
	all := 0

	for _, design := range designs {
		pd := findAll(design, patterns)

		all += pd
	}

	return all
}

var found = map[string]int{}

func findAll(current string, patterns []string) int {
	if r, ok := found[current]; ok {
		return r
	}
	res := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(current, pattern) {
			if len(current) == len(pattern) {
				res++
			} else {
				res += findAll(current[len(pattern):], patterns)
			}
		}
	}
	found[current] = res
	return res
}
