package inputs

import (
	"strconv"
	"strings"
)

func Lines(inputStr string) []string {
	return strings.Split(inputStr, "\n")
}

func Numbers(inputStr, sep string) []int {
	var res []int
	for _, nStr := range strings.Split(inputStr, sep) {
		if nStr == "" {
			continue
		}
		n, _ := strconv.Atoi(nStr)
		res = append(res, n)
	}
	return res
}

func Numbers64(inputStr, sep string) []int64 {
	var res []int64
	for _, nStr := range strings.Split(inputStr, sep) {
		if nStr == "" {
			continue
		}
		n, _ := strconv.Atoi(nStr)
		res = append(res, int64(n))
	}
	return res
}
