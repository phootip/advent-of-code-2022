package day11

import (
	"fmt"
	"sort"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day11 Solution1...")
	raw := utils.ReadFile("./day11/input.txt")
	// raw := utils.ReadFile("./day11/example.txt")
	monkeys := rawToMonkeys(raw)
	fmt.Println(monkeys)
	for i := 0; i < 20; i++ {
		compute(monkeys)
	}
	return getAns(monkeys)
}

func getAns(monkeys []monkey) int {
	result := []int{}
	for _, m := range monkeys {
		result = append(result, m.ans)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result[0] * result[1]
}

func compute(monkeys []monkey) {
	for i := range monkeys {
		m := &monkeys[i]
		m.inspect(monkeys)
	}
}

func (m *monkey) inspect(monkeys []monkey) {
	var item int
	for len(m.items) > 0 {
		m.ans++
		item, m.items = m.items[0], m.items[1:len(m.items)]
		item = m.op(item) / 3

		if item%m.test == 0 {
			monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, item)
		} else {
			monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, item)
		}
	}
}

func rawToMonkeys(raw []string) (monkeys []monkey) {
	var lastMonkey *monkey
	state := "newMonkey"
	for _, line := range raw {
		if line == "" {
			continue
		}
		if len(monkeys) > 0 {
			lastMonkey = &monkeys[len(monkeys)-1]
		}
		if state == "newMonkey" {
			newMonkey := monkey{items: []int{}}
			monkeys = append(monkeys, newMonkey)
			state = "items"
		} else if state == "items" {
			line = line[18:]
			splited := strings.Split(line, ", ")
			items := utils.SToInt(splited)
			lastMonkey.items = items
			state = "operation"
		} else if state == "operation" {
			line = line[19:]
			splited := strings.Split(line, " ")
			op := operation(splited)
			lastMonkey.op = op
			state = "test"
		} else if state == "test" {
			line = line[21:]
			lastMonkey.test = utils.StringToInt(line)
			state = "ifTrue"
		} else if state == "ifTrue" {
			line = line[29:]
			lastMonkey.ifTrue = utils.StringToInt(line)
			state = "ifFalse"
		} else if state == "ifFalse" {
			line = line[30:]
			lastMonkey.ifFalse = utils.StringToInt(line)
			state = "newMonkey"
		}
	}
	return monkeys
}

func operation(command []string) func(int) int {
	a := command[0]
	op := command[1]
	b := command[2]

	return func(old int) int {
		if op == "+" {
			return stringToVar(a, old) + stringToVar(b, old)
		} else {
			return stringToVar(a, old) * stringToVar(b, old)
		}
	}
}

func stringToVar(a string, old int) int {
	if a == "old" {
		return old
	} else {
		return utils.StringToInt(a)
	}
}

type monkey struct {
	items   []int
	op      func(int) int
	test    int
	ifTrue  int
	ifFalse int
	ans     int
}
