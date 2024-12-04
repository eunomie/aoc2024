package main

import (
	"fmt"
	"regexp"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d04p1(inputs.D04))
	fmt.Println(d04p2(inputs.D04))
}

func d04p1(input string) int {
	var res int
	const xmas = "XMAS"
	xmasReg := regexp.MustCompile(xmas)
	strs := xmasReg.FindAllString(input, -1)
	res += len(strs)

	d := sugar.Diagonals(input)
	strs = xmasReg.FindAllString(d, -1)
	res += len(strs)

	p := sugar.PivotString(input)
	strs = xmasReg.FindAllString(p, -1)
	res += len(strs)

	d = sugar.Diagonals(p)
	strs = xmasReg.FindAllString(d, -1)
	res += len(strs)

	p = sugar.PivotString(p)
	strs = xmasReg.FindAllString(p, -1)
	res += len(strs)

	d = sugar.Diagonals(p)
	strs = xmasReg.FindAllString(d, -1)
	res += len(strs)

	p = sugar.PivotString(p)
	strs = xmasReg.FindAllString(p, -1)
	res += len(strs)

	d = sugar.Diagonals(p)
	strs = xmasReg.FindAllString(d, -1)
	res += len(strs)

	return res
}

func d04p2(input string) int {
	nbCross := 0
	lines := inputs.Lines(input)
	nbLines := len(lines)
	lineLen := len(lines[0])
	for row := 1; row < nbLines-1; row++ {
		for col := 1; col < lineLen-1; col++ {
			if lines[row][col] == 'A' {
				tl := lines[row-1][col-1]
				tr := lines[row-1][col+1]
				bl := lines[row+1][col-1]
				br := lines[row+1][col+1]
				cross := fmt.Sprintf("%s%s%s%s", string(tl), string(tr), string(br), string(bl))
				switch cross {
				case "MMSS", "MSSM", "SSMM", "SMMS":
					nbCross++
				}
			}
		}
	}
	return nbCross
}
