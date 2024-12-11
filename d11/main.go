package main

import (
	"fmt"
	"time"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d11p1(inputs.D11))
	fmt.Println(d11p2(inputs.D11))
}

func d11p1(input string) int {
	stones := inputs.Numbers(input, " ")

	t := time.Now()
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}
	fmt.Println(time.Since(t))

	return len(stones)
}

func blink(stones []int) []int {
	var newStones []int
	for _, stone := range stones {
		newStones = blinkStone(stone, newStones)
	}
	return newStones
}

type Op struct {
	stone      int
	iterations int
}

var cachedOperations = map[Op]int{}

func d11p2(input string) int {
	stones := inputs.Numbers(input, " ")
	var nbStones int

	t := time.Now()
	for _, stone := range stones {
		nbStones += blinkN(stone, 75)
	}
	fmt.Println(time.Since(t))

	return nbStones
}

func blinkN(stone int, left int) int {
	if left == 0 {
		return 1
	}
	op := Op{stone, left}
	if v, ok := cachedOperations[op]; ok {
		return v
	}
	if stone == 0 {
		v := blinkN(1, left-1)
		cachedOperations[op] = v
		return v
	}
	stoneStr := fmt.Sprintf("%d", stone)
	if len(stoneStr)%2 == 0 {
		runes := []rune(stoneStr)
		v := blinkN(sugar.MustAtoi(string(runes[:len(runes)/2])), left-1) +
			blinkN(sugar.MustAtoi(string(runes[len(runes)/2:])), left-1)
		cachedOperations[op] = v
		return v
	}
	v := blinkN(stone*2024, left-1)
	cachedOperations[op] = v
	return v
}

func blinkStone(stone int, newStones []int) []int {
	if stone == 0 {
		newStones = append(newStones, 1)
		return newStones
	}
	stoneStr := fmt.Sprintf("%d", stone)
	if len(stoneStr)%2 == 0 {
		runes := []rune(stoneStr)
		newStones = append(newStones,
			sugar.MustAtoi(string(runes[:len(runes)/2])),
			sugar.MustAtoi(string(runes[len(runes)/2:])),
		)
		return newStones
	}
	newStones = append(newStones, stone*2024)
	return newStones
}
