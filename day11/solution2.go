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
	fmt.Println(monkeys)
	for i := 0; i < 20; i++ {
		compute(monkeys)
	}
	return getAns(monkeys)
}
