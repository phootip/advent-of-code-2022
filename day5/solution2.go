package day5

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol2() string {
	fmt.Println("Starting Day5 Solution2...")
	raw := utils.ReadFile("./day5/input.txt")
	// raw := utils.ReadFile("./day5/example1.txt")
	ans := ""
	crates, insts := parseRaw(raw)
	fmt.Println(crates)
	fmt.Println(insts)
	for _, inst := range insts {
		fmt.Println("Moving crates")
		crates = moveN(crates, inst)
	fmt.Println(inst)
	fmt.Println(crates)
	}
	for _, s := range crates {
		if len(s) > 0 {
			ans += string(s[len(s)-1])
		}
	}
	return ans
}

func moveN(crates [10]string, inst []int) [10]string {
	n := inst[0]
	s := inst[1]
	d := inst[2]
	c := crates[s][len(crates[s])-n:len(crates[s])]
	crates[s] = crates[s][:len(crates[s])-n]
	crates[d] = crates[d] + c
	return crates
}
