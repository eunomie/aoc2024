package sugar

import (
	"golang.org/x/exp/constraints"
	"strconv"
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
