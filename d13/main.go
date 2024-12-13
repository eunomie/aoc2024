package main

import (
	"fmt"
	"regexp"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	//fmt.Println(d13p1(inputs.D13))
	fmt.Println(d13p2(inputs.D13))
}

func d13p1(input string) int {
	lines := inputs.Lines(input)
	nbGames := len(lines) / 4

	var lineA, lineB, linePrize string
	buttonRegExp := regexp.MustCompile("^Button [AB]: X\\+(\\d+), Y\\+(\\d+)$")
	prizeRegExp := regexp.MustCompile("^Prize: X=(\\d+), Y=(\\d+)$")
	var ax, ay, bx, by, px, py int
	var matches []string
	var cost int
	var hasSolution bool
	var limit int
	var start int

	const (
		costA = 3
		costB = 1
		maxA  = 100
		maxB  = 100
	)

	var nbA, nbB int
	var totalCost int

	for i := 0; i < nbGames; i++ {
		lineA = lines[i*4]
		matches = buttonRegExp.FindStringSubmatch(lineA)
		ax, ay = sugar.MustAtoi(matches[1]), sugar.MustAtoi(matches[2])
		lineB = lines[i*4+1]
		matches = buttonRegExp.FindStringSubmatch(lineB)
		bx, by = sugar.MustAtoi(matches[1]), sugar.MustAtoi(matches[2])
		linePrize = lines[i*4+2]
		matches = prizeRegExp.FindStringSubmatch(linePrize)
		px, py = sugar.MustAtoi(matches[1]), sugar.MustAtoi(matches[2])

		hasSolution = false
		limit = min(max(px/ax, py/ay, px/bx, py/by), 100)
		start = min(px/ax, py/ay, px/bx, py/by)
		for nbA = start; nbA < limit; nbA++ {
			if (px-nbA*ax)%bx != 0 {
				continue
			}
			nbB = (px - nbA*ax) / bx
			if nbA > maxA || nbB > maxB {
				continue
			}
			if nbA*ay+nbB*by == py {
				if !hasSolution {
					cost = nbA*costA + nbB*costB
					hasSolution = true
				} else if nbA*costA+nbB*costB < cost {
					cost = nbA*costA + nbB*costB
				}
			}
		}
		if hasSolution {
			totalCost += cost
		}
	}

	return totalCost
}

func d13p2(input string) int {
	lines := inputs.Lines(input)
	nbGames := len(lines) / 4

	var lineA, lineB, linePrize string
	buttonRegExp := regexp.MustCompile("^Button [AB]: X\\+(\\d+), Y\\+(\\d+)$")
	prizeRegExp := regexp.MustCompile("^Prize: X=(\\d+), Y=(\\d+)$")
	var ax, ay, bx, by, px, py int
	var matches []string

	const (
		costA = 3
		costB = 1
	)

	var totalCost int

	var nbA, nbB int
	var da, dax, day int
	for i := 0; i < nbGames; i++ {
		lineA = lines[i*4]
		matches = buttonRegExp.FindStringSubmatch(lineA)
		ax, ay = sugar.MustAtoi(matches[1]), sugar.MustAtoi(matches[2])
		lineB = lines[i*4+1]
		matches = buttonRegExp.FindStringSubmatch(lineB)
		bx, by = sugar.MustAtoi(matches[1]), sugar.MustAtoi(matches[2])
		linePrize = lines[i*4+2]
		matches = prizeRegExp.FindStringSubmatch(linePrize)
		px, py = sugar.MustAtoi(matches[1])+10000000000000, sugar.MustAtoi(matches[2])+10000000000000

		da = ax*by - bx*ay
		if da <= 0 {
			continue
		}
		dax = px*by - bx*py
		day = py*ax - ay*px
		if dax%da != 0 || day%da != 0 {
			continue
		}
		nbA = dax / da
		nbB = day / da
		totalCost += nbA*costA + nbB*costB
	}

	return totalCost
}
