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

func Sol2() (ans int) {
	fmt.Println("Starting Day21 Solution2...")
	// raw := utils.ReadFile("./day21/input.txt")
	raw := utils.ReadFile("./day21/example.txt")
	raw = raw[:len(raw)-1]
	monkeys := parseRaw(raw)
	// monkeys["root"].op = "="
	// you := monkeys["humn"]
	you := "humn"
	// fmt.Println(you)
	// debug(monkeys)
	ans = resolve2(monkeys, you)
	return ans
}
func resolve2(monkeys map[string]*Monkey, you string) (ans int) {
	// ans, path := monkeys["root"].getValue2([]*Monkey{})
	// fmt.Println(path)
	root := monkeys["root"]
	// monkeys["humn"].hasValue = false
	stack := []*Monkey{root}
	// stack := []*Monkey{root.monkey1P, root.monkey2P}
	// visited := []*Monkey{}

	for len(stack) > 0 {
		m := stack[len(stack)-1]

		if m.name == you {
			fmt.Println("found human: ",m)
			for _, m := range stack {
				fmt.Println(m)
			}
			break
		}
		if !m.monkey1P.hasValue {
			stack = append(stack, m.monkey1P)
			continue
		}
		if !m.monkey2P.hasValue {
			stack = append(stack, m.monkey2P)
			continue
		}
		m.value = m.resolveValue()
		m.hasValue = true
		stack = stack[:len(stack)-1]
	}

	return root.value
}

func resolve(monkeys map[string]*Monkey) (ans int) {
	return monkeys["root"].getValue()
}

func (m *Monkey) resolveValue() int {
	a := m.monkey1P.value
	b := m.monkey2P.value
	m.hasValue = true
	switch m.op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	panic("can't resolve value")
}

func (monkey *Monkey) getValue2(path []*Monkey) (int, []*Monkey) {
	// fmt.Println("Getting monkey value:", monkey)
	path = append(path, monkey)
	if monkey.hasValue {
		return monkey.value, path
	}
	monkey1 := monkey.monkey1P
	monkey2 := monkey.monkey2P
	a := monkey1.getValue()
	b := monkey2.getValue()
	switch monkey.op {
	case "+":
		return a + b, path
	case "-":
		return a - b, path
	case "*":
		return a * b, path
	case "/":
		return a / b, path
	case "=":
		fmt.Println("a:", a)
		fmt.Println("b:", b)
		return 0, path
	}
	panic("can't calculate monkey value")
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
		m.monkey1P.parent = m
		m.monkey2P.parent = m
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
	parent   *Monkey
	// path     []string
}
