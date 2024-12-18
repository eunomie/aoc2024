package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/eunomie/aoc2024/inputs"
	"github.com/eunomie/aoc2024/sugar"
)

func main() {
	fmt.Println(d17p1(inputs.D17))
	fmt.Println(d17p2(inputs.D17))
}

type Computer struct {
	registerA int
	registerB int
	registerC int
	program   []int
}

func (c *Computer) combo(op int) int {
	if op <= 3 {
		return op
	}
	if op == 4 {
		return c.registerA
	}
	if op == 5 {
		return c.registerB
	}
	if op == 6 {
		return c.registerC
	}
	panic("invalid operation")
}

func (c *Computer) runAll() []int {
	pointer := 0
	l := len(c.program)
	var inst Instruction
	var op int
	var outs []int
	for {
		if pointer >= l {
			break
		}
		inst = c.program[pointer]
		op = c.program[pointer+1]

		switch inst {
		case adv:
			c.registerA = div2(c.registerA, c.combo(op))
			pointer += 2
		case bxl:
			c.registerB = c.registerB ^ op
			pointer += 2
		case bst:
			comb := c.combo(op)
			comb = comb % 8
			c.registerB = comb
			pointer += 2
		case jnz:
			if c.registerA > 0 {
				pointer = op
			} else {
				pointer += 2
			}
		case bxc:
			c.registerB = c.registerB ^ c.registerC
			pointer += 2
		case out:
			outs = append(outs, c.combo(op)%8)
			pointer += 2
		case bdv:
			c.registerB = div2(c.registerA, c.combo(op))
			pointer += 2
		case cdv:
			c.registerC = div2(c.registerA, c.combo(op))
			pointer += 2
		}
	}

	return outs
}

type Instruction = int

const (
	adv Instruction = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func d17p1(input string) string {
	lines := inputs.Lines(input)
	ra := sugar.MustAtoi(strings.TrimPrefix(lines[0], "Register A: "))
	rb := sugar.MustAtoi(strings.TrimPrefix(lines[1], "Register B: "))
	rc := sugar.MustAtoi(strings.TrimPrefix(lines[2], "Register C: "))
	program := inputs.Numbers(strings.TrimPrefix(lines[4], "Program: "), ",")

	computer := &Computer{
		registerA: ra,
		registerB: rb,
		registerC: rc,
		program:   program,
	}

	outs := computer.runAll()

	return strings.Join(sugar.Map(outs, strconv.Itoa), ",")
}

func div2(a, nb int) int {
	for i := 0; i < nb; i++ {
		a = a / 2
	}
	return a
}

func d17p2(input string) int {
	lines := inputs.Lines(input)
	ra := sugar.MustAtoi(strings.TrimPrefix(lines[0], "Register A: "))
	rb := sugar.MustAtoi(strings.TrimPrefix(lines[1], "Register B: "))
	rc := sugar.MustAtoi(strings.TrimPrefix(lines[2], "Register C: "))
	ip := strings.TrimPrefix(lines[4], "Program: ")
	program := inputs.Numbers(ip, ",")
	lp := len(program)
	fmt.Println("len:", lp)

	// program will loop until A is set to 0
	// A is always set to A / 8
	// so the last step means A was between 1 and 7 (so it can be divided by 8 and result is 0)
	// previous was 8 times this value, so between 8 and 56 and so on
	// In the end, the interval is way too big, so try from the end

	var i int
	var a int
	for i = lp - 1; i >= 0; i-- {
		a *= 8
		for {
			if a == ra {
				continue
			}
			computer := &Computer{
				registerA: a,
				registerB: rb,
				registerC: rc,
				program:   program,
			}
			res := computer.runAll()
			if len(res) == lp && slices.Equal(res, program) {
				return a
			}
			if slices.Equal(res, program[i:]) {
				break
			}

			a++
		}
	}

	return -1
}
