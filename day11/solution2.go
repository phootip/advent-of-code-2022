package day11

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol2() int {
	fmt.Println("Starting Day11 Solution2...")
	raw := utils.ReadFile("./day11/input.txt")
	// raw := utils.ReadFile("./day11/example.txt")
	monkeys := rawToMonkeys(raw)
	lcm := getLCM(monkeys)
	fmt.Println(monkeys)
	for i := 1; i <= 10000; i++ {
		compute2(monkeys, lcm)
		// if i == 1 || i == 20 || i == 1000{
		// 	fmt.Println("loop: ",i)
		// 	debug(monkeys)
		// }
	}
	return getAns(monkeys)
}

func getLCM(monkeys []monkey) int {
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.test
	}
	return lcm
}

func debug(monkeys []monkey) {
	for _, m := range monkeys {
		fmt.Println(m.ans)
	}
}

func compute2(monkeys []monkey, lcm int) {
	for i := range monkeys {
		m := &monkeys[i]
		m.inspect2(monkeys, lcm)
	}
}

func (m *monkey) inspect2(monkeys []monkey, lcm int) {
	for len(m.items) > 0 {
		m.ans++
		item := m.items[0]
		m.items = m.items[1:]
		item = m.op(item) % lcm

		if item % m.test == 0 {
			monkeys[m.ifTrue].items = append(monkeys[m.ifTrue].items, item)
		} else {
			monkeys[m.ifFalse].items = append(monkeys[m.ifFalse].items, item)
		}
	}
}
