package day3

import (
	"fmt"
	"strings"

	"github.com/juliangruber/go-intersect"
	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day3 Solution1...")
	raw := utils.ReadFile("./day3/input.txt")
	// raw := utils.ReadFile("./day3/example1.txt")
	result := 0
	for _, line := range raw {
		if line == "" {
			break
		}
		fmt.Println(line)
		v := getDup(line)
		result += v
	}
	return result
}

func getDup(line string) int {
	mem := []rune{}
	for i, r := range line {
		if i < len(line)/2 {
			// fmt.Print("first compartment")
			mem = append(mem, r)
		} else {
			if utils.Contains(mem, r) {
				return runeToAns(r)
			}
		}
		// fmt.Printf("i%d r %c %v \n", i, r, r)
	}
	return 0
}

func runeToAns(r rune) int {
	ans := int(r) - 96
	if ans < 0 {
		ans += 58
	}
	return ans
}

func Sol2() int {
	fmt.Println("Starting Day3 Solution2...")
	raw := utils.ReadFile("./day3/input.txt")
	// raw := utils.ReadFile("./day3/example1.txt")
	result := 0
	mem := [][]string{}
	for i, line := range raw {
		if line == "" {
			break
		}
		fmt.Println(i, line)
		mem = append(mem, strings.Split(line,""))
		if i%3 == 2 {
			fmt.Println("calculating...")
			r := getBadges(mem)
			fmt.Println(runeToAns(r))
			result += runeToAns(r)
			mem = [][]string{}
		}
	}
	return result
}

func getBadges(sacks [][]string) rune {
	fmt.Println(sacks)
	a := intersect.Hash(intersect.Hash(sacks[0], sacks[1]), sacks[2])
	for _, r := range a[0].(string) {
		return r
	}
	return 0
}
