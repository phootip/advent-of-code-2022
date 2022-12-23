package day21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() (ans int) {
	fmt.Println("Starting Day21 Solution1...")
	raw := utils.ReadFile("./day21/input.txt")
	// raw := utils.ReadFile("./day21/example.txt")
	raw = raw[:len(raw)-1]
	monkeys := parseRaw(raw)
	// debug(monkeys)
	ans = resolve(monkeys)
	return ans
}

// func resolve(monkeys map[string]*Monkey) (ans int) {
// 	root := monkeys["root"]
// 	stack := []*Monkey{root}
// 	visited := []*Monkey{}

// 	for len(stack) > 0 {
// 		monkey := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		monkey1 := monkeys[monkey.monkey1]
// 		monkey2 := monkeys[monkey.monkey2]
// 		if monkey.hasValue {
// 			continue
// 		}
// 		if !monkey1.hasValue {
// 			stack = append(stack, monkey1)
// 			visited = append(visited, monkey1)
// 		}
// 		if !monkey2.hasValue {
// 			stack = append(stack, monkey2)
// 			visited = append(visited, monkey2)
// 		}
// 	}
// 	return monkeys["root"].value
// }

func resolve(monkeys map[string]*Monkey) (ans int) {
	// root := monkeys["root"]
	return monkeys["root"].getValue()
}

func (monkey *Monkey) getValue() int {
	// fmt.Println("Getting monkey value:", monkey)
	if monkey.hasValue {
		return monkey.value
	}
	monkey1 := monkey.monkey1P
	monkey2 := monkey.monkey2P
	a := monkey1.getValue()
	b := monkey2.getValue()
	switch monkey.op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	panic("can't calculate monkey value")
}

func parseRaw(raw []string) map[string]*Monkey {
	monkeys := make(map[string]*Monkey)
	for _, line := range raw {
		line2 := strings.Split(line, ": ")
		monkey := &Monkey{}
		name := line2[0]
		monkey.name = name
		value, err := strconv.Atoi(line2[1])
		if err == nil {
			monkey.value = value
			monkey.hasValue = true
		} else {
			line3 := strings.Split(line2[1], " ")
			monkey.monkey1 = line3[0]
			monkey.op = line3[1]
			monkey.monkey2 = line3[2]
		}
		monkeys[name] = monkey
	}
	for _, m := range monkeys {
		if m.hasValue {
			continue
		}
		m.monkey1P = monkeys[m.monkey1]
		m.monkey2P = monkeys[m.monkey2]
	}
	return monkeys
}

func debug(data map[string]*Monkey) {
	for _, m := range data {
		fmt.Println(m)
	}
}

type Monkey struct {
	name     string
	hasValue bool
	value    int
	monkey1  string
	monkey1P *Monkey
	op       string
	monkey2  string
	monkey2P *Monkey
}
