package main

import (
	"fmt"

	"github.com/eunomie/aoc2024/inputs"
)

func main() {
	fmt.Println(d09p1(inputs.D09))
	fmt.Println(d09p2(inputs.D09))
}

func d09p1(input string) int {
	var disk = diskMap(input)

	var i int
	var j = len(disk)
	for i = 0; i < len(disk); i++ {
		if disk[i] == -1 {
			if i == j {
				break
			}
			for {
				j--
				if disk[j] != -1 {
					disk[i] = disk[j]
					disk[j] = -1
					break
				}
			}
		}
	}

	return checksum(disk)
}

func diskMap(input string) []int {
	var disk []int

	var id int
	var file = true
	var v int32
	var nbBlocks int
	var i int
	for _, v = range input {
		nbBlocks = btoi(v)
		if file {
			for i = 0; i < nbBlocks; i++ {
				disk = append(disk, id)
			}
			id++
		} else {
			for i = 0; i < nbBlocks; i++ {
				disk = append(disk, -1)
			}
		}
		file = !file
	}

	return disk
}

func checksum(disk []int) int {
	var res, i int
	for i = 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}
		res += disk[i] * i
	}

	return res
}

func btoi(v int32) int {
	return int(v - '0')
}

func d09p2(input string) int {
	var disk = diskMap(input)
	var movedBlocks = make([]int, len(disk))

	var i int
	var j int
	var end = -1
	var l int
	var lfree int
	var k int
	var v = -2

	end = len(disk) - 1
	v = disk[end]

	for j = len(disk) - 2; j >= 0; j-- {
		if disk[j] != v {
			l = end - j

			if v != -1 && movedBlocks[j+1] != -1 {
				lfree = 0
				for i = 0; i < j+1; i++ {
					if disk[i] != -1 {
						lfree = 0
					} else {
						lfree++
						if lfree == l {
							for k = 0; k < l; k++ {
								disk[i-k] = v
								disk[j+k+1] = -1
								movedBlocks[i-k] = -1
							}
							lfree = 0
							break
						}
					}
				}
			}

			end = j
			v = disk[end]
		}
	}

	return checksum(disk)
}
