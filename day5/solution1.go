package day5

import (
	"fmt"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() string {
	fmt.Println("Starting Day5 Solution1...")
	raw := utils.ReadFile("./day5/input.txt")
	// raw := utils.ReadFile("./day5/example1.txt")
	ans := ""
	crates, insts := parseRaw(raw)
	fmt.Println(crates)
	fmt.Println(insts)
	for _, inst := range insts {
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
	for i := 0; i < inst[0]; i++ {
		crates = move(crates, inst[1], inst[2])
	}
	return crates
}

func move(crates [10]string, s int, d int) [10]string {
	c := crates[s][len(crates[s])-1]
	crates[s] = crates[s][:len(crates[s])-1]
	crates[d] = crates[d] + string(c)
	return crates
}

func parseRaw(raw []string) ([10]string, [][]int) {
	crates := [10]string{}
	insts := [][]int{}
	parseCrates := true
	for _, line := range raw {
		fmt.Println(line)
		// parse crates
		if parseCrates {
			for i, r := range line {
				// break if 1 2 3 line
				if r == 49 {
					break
				}
				if i%4 == 1 {
					// fmt.Println(i, i/4, string(r), r)
					if r != 32 {
						crates[(i/4)+1] = string(r) + crates[(i/4)+1]
					}
				}
			}
		} else {
			// parse inst
			fmt.Println("Parsing Inst...")
			if line == "" {
				break
			}
			insts = append(insts, parseInst(line))
		}
		if line == "" {
			parseCrates = false
		}
	}
	return crates, insts
}

func parseInst(line string) []int {
	splited := strings.Split(line, " ")
	filtered := []string{}
	for i, s := range splited {
		if i%2 == 1 {
			filtered = append(filtered, s)
		}
	}
	return utils.SToInt(filtered)
}
