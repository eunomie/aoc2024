package main

import (
	"fmt"

	"github.com/eunomie/aoc2024/inputs"
)

func main() {
	fmt.Println(d05p1(inputs.D05))
	fmt.Println(d05p2(inputs.D05))
}

type (
	Constraint struct {
		before, after int
	}

	Constraints struct {
		constraints []Constraint
	}
)

func (c *Constraints) AddConstraint(before, after int) {
	c.constraints = append(c.constraints, Constraint{before, after})
}

func (c *Constraints) IsValid(before, after int) bool {
	for _, constraint := range c.constraints {
		if constraint.before == before && constraint.after == after {
			return true
		}
		if constraint.before == after && constraint.after == before {
			return false
		}
	}
	panic("not found")
}

func d05p1(input string) int {
	var res int

	c := &Constraints{
		constraints: []Constraint{},
	}
	lines := inputs.Lines(input)
	readingConstraints := true
	for _, line := range lines {
		if line == "" {
			readingConstraints = false
			continue
		}
		if readingConstraints {
			n := inputs.Numbers(line, "|")
			c.AddConstraint(n[0], n[1])
		} else {
			pages := inputs.Numbers(line, ",")
			nbPages := len(pages)
			for i := 0; i < nbPages; i++ {
				for j := 0; j < i; j++ {
					if !c.IsValid(pages[j], pages[i]) {
						goto NEXT
					}
				}
				if i < nbPages-1 {
					for j := i + 1; j < nbPages; j++ {
						if !c.IsValid(pages[i], pages[j]) {
							goto NEXT
						}
					}
				}
			}
			res += pages[nbPages/2]
		NEXT:
		}
	}

	return res
}

func d05p2(input string) int {
	var res int

	c := &Constraints{
		constraints: []Constraint{},
	}
	lines := inputs.Lines(input)
	readingConstraints := true
	for _, line := range lines {
		if line == "" {
			readingConstraints = false
			continue
		}
		if readingConstraints {
			n := inputs.Numbers(line, "|")
			c.AddConstraint(n[0], n[1])
		} else {
			pages := inputs.Numbers(line, ",")
			nbPages := len(pages)
			fixed := false
		REDO:
			for i := 0; i < nbPages; i++ {
				for j := 0; j < i; j++ {
					if !c.IsValid(pages[j], pages[i]) {
						pages[i], pages[j] = pages[j], pages[i]
						fixed = true
						goto REDO
					}
				}
				if i < nbPages-1 {
					for j := i + 1; j < nbPages; j++ {
						if !c.IsValid(pages[i], pages[j]) {
							pages[i], pages[j] = pages[j], pages[i]
							fixed = true
							goto REDO
						}
					}
				}
			}
			if fixed {
				res += pages[nbPages/2]
			}
		}
	}

	return res
}
