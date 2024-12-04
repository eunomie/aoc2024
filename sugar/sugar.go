package sugar

import (
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"

	"github.com/eunomie/aoc2024/inputs"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Each[S interface{ ~[]E }, E any](arr S, fn func(E)) {
	for _, v := range arr {
		fn(v)
	}
}

func Map[S interface{ ~[]E }, E any, R any](arr S, fn func(E) R) []R {
	res := make([]R, 0, len(arr))
	for _, v := range arr {
		res = append(res, fn(v))
	}
	return res
}

func FilterIndex[S interface{ ~[]E }, E any](arr S, idx int) []E {
	if idx < 0 {
		return arr
	}
	res := make([]E, 0, len(arr)-1)
	for i, v := range arr {
		if i != idx {
			res = append(res, v)
		}
	}
	return res
}

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func IfFunc[T any](cond bool, a, b func() T) T {
	if cond {
		return a()
	}
	return b()
}

func MustAtoi(a string) int {
	v, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return v
}

func Reverse[S ~[]E, E any](s S) S {
	c := slices.Clone(s)
	slices.Reverse(c)
	return c
}

func StringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func StringEqualOrReverse(s, eq string) bool {
	if s == eq {
		return true
	}
	if StringReverse(s) == eq {
		return true
	}
	return false
}

func PivotString(s string) string {
	lines := strings.Split(s, "\n")
	nbLines := len(lines)
	lineLen := len(lines[0])
	runes := make([]rune, len(s))
	var idx int
	for i := lineLen - 1; i >= 0; i-- {
		for j := 0; j < nbLines; j++ {
			runes[idx] = rune(lines[j][i])
			idx++
		}
		if idx < len(s) {
			runes[idx] = '\n'
			idx++
		}
	}
	return string(runes)
}

func Diagonals(input string) string {
	lines := inputs.Lines(input)
	lineLen := len(lines[0])
	nbLines := len(lines)
	runes := make([]rune, len(input)+nbLines-1)
	firstTriangle := true
	var i, row, col, ir, ic int
	for {
		runes[i] = rune(lines[row][col])
		i++
		row++
		col++
		if row == nbLines || col == lineLen {
			if firstTriangle {
				ic++
				if ic == lineLen {
					firstTriangle = false
					ir++
					if ir == lineLen {
						break
					}
					ic = 0
				}
			} else {
				ir++
				if ir == lineLen {
					break
				}
			}
			row = ir
			col = ic
			runes[i] = '\n'
			i++
			continue
		}
	}

	return string(runes)
}
