package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

// 501 too high

func Sol1() int {
	fmt.Println("Starting Day4 Solution1...")
	raw := utils.ReadFile("./day4/input.txt")
	// raw := utils.ReadFile("./day4/example1.txt")
	ans := 0
	for _, line := range raw {
		if line == "" {
			break
		}
		// fmt.Println(line)
		// get sections
		elfs := strings.Split(line, ",")
		elf1 := SToInt(strings.Split(elfs[0], "-"))
		elf2 := SToInt(strings.Split(elfs[1], "-"))
		// fmt.Println(elf1, elf2)
		if inRangeOf(elf1, elf2) || inRangeOf(elf2, elf1) {
			// fmt.Println(elf1, elf2)
			ans += 1
		}
	}
	return ans
}

func Sol2() int {
	fmt.Println("Starting Day4 Solution2...")
	raw := utils.ReadFile("./day4/input.txt")
	// raw := utils.ReadFile("./day4/example1.txt")
	ans := 0
	for _, line := range raw {
		if line == "" {
			break
		}
		// fmt.Println(line)
		// get sections
		elfs := strings.Split(line, ",")
		elf1 := SToInt(strings.Split(elfs[0], "-"))
		elf2 := SToInt(strings.Split(elfs[1], "-"))
		// fmt.Println(elf1, elf2)
		if overlap(elf1, elf2) || overlap(elf2, elf1) {
			fmt.Println(elf1, elf2)
			ans += 1
		}
	}
	return ans
}

func inRangeOf(elf1 []int, elf2 []int) bool {
	return elf1[0] >= elf2[0] && elf1[1] <= elf2[1]
}

func overlap(elf1 []int, elf2 []int) bool {
	return (elf2[0] <= elf1[0] && elf1[0] <= elf2[1]) || (elf2[0] <= elf1[1] && elf1[1] <= elf2[1])
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func SToInt(ss []string) []int {
	result := make([]int, len(ss))
	for i, s := range ss {
		r, err := strconv.Atoi(s)
		utils.Check(err)
		result[i] = r
	}
	return result
}

// func Sol2() int {
// 	fmt.Println("Starting Day3 Solution1...")
// 	// raw := utils.ReadFile("./day4/input.txt")
// 	// raw := utils.ReadFile("./day4/example1.txt")
// 	return 0
// }
