package inputs

import (
	"strconv"
	"strings"
)

func Lines(inputStr string) []string {
	return strings.Split(inputStr, "\n")
}

func Numbers(inputStr string) []int {
	var res []int
	for _, nStr := range strings.Split(inputStr, " ") {
		if nStr == "" {
			continue
		}
		n, _ := strconv.Atoi(nStr)
		res = append(res, n)
	}
	return res
}
