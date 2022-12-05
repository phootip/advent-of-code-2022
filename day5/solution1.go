package day5

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() string {
	fmt.Println("Starting Day5 Solution1...")
	// raw := utils.ReadFile("./day5/input.txt")
	raw := utils.ReadFile("./day5/example1.txt")
	ans := ""
	crates, inst := parseRaw(raw)
	fmt.Println(crates)
	fmt.Println(inst)
	return ans
}

func parseRaw(raw []string) ([10]string, []int) {
	crates := [10]string{}
	inst := []int{}
	for _, line := range raw {
		if line == "" {
			break
		}
		fmt.Println(line)
		// parse crates
		for i, r := range line {
			// break if 1 2 3 line
			if r == 49 {
				break
			}
			if i%4 == 1 {
				// fmt.Println(i, i/4, string(r), r)
				if r != 32 {
					crates[(i/4)+1] += string(r)
				}
			}
		}
	}
	fmt.Println(crates[1])
	fmt.Println(crates[2])
	fmt.Println(crates[3])
	return crates, inst
}
