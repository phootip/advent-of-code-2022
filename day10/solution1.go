package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day10 Solution1...")
	raw := utils.ReadFile("./day10/input.txt")
	// raw := utils.ReadFile("./day10/example.txt")
	insts := rawToInst(raw)
	com := cpu{x: 1, cycle: 0, counter: 0, insts: insts, terminated: false, ans: 0}
	for !com.terminated {
		com.compute()
	}
	return com.ans
}

func (com *cpu) compute() {
	if com.counter >= len(com.insts) {
		com.terminated = true
		return
	}
	com.cycle++
	com.updateAns()
	// fmt.Println("cycles: ", com.cycle)
	if com.loaded {
		// com.updateAns()
		com.x += com.insts[com.counter].arg
		com.loaded = false
		com.counter++
		return
	}
	switch com.insts[com.counter].op {
	case "addx":
		com.loaded = true
		return
	}
	// com.updateAns()
	com.counter++
}

func (com *cpu) updateAns() {
	if com.cycle%40-20 == 0 && com.cycle <= 220 {
		fmt.Println("signal: ", com.cycle, com.x)
		com.ans += com.cycle * com.x
	}
}

func rawToInst(raw []string) []instruction {
	insts := make([]instruction, 0)
	for _, line := range raw {
		if line == "" {
			return insts
		}
		temp := strings.Split(line, " ")
		op := temp[0]
		if op == "noop" {
			insts = append(insts, instruction{op, 0})
		} else {
			arg, err := strconv.Atoi(temp[1])
			utils.Check(err)
			insts = append(insts, instruction{op, arg})
		}
	}
	return insts

}

type instruction struct {
	op  string
	arg int
}

type cpu struct {
	x          int
	cycle      int
	counter    int
	insts      []instruction
	loaded     bool
	terminated bool
	ans        int
}
