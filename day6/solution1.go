package day6

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day6 Solution1...")
	raw := utils.ReadFile("./day6/input.txt")
	// raw := utils.ReadFile("./day6/example1.txt")
	data := raw[0]
	// for _, line := range raw {
	// 	fmt.Println(findPacket(line))
	// }
	return findPacket(data, 4)
}
func Sol2() int {
	fmt.Println("Starting Day6 Solution2...")
	raw := utils.ReadFile("./day6/input.txt")
	// raw := utils.ReadFile("./day6/example1.txt")
	data := raw[0]
	// for _, line := range raw {
	// 	fmt.Println(findPacket(line, 14))
	// }
	return findPacket(data, 14)
}

func findPacket(stream string, l int) int {
	var removed rune
	container := []rune{}
	dict := make(map[rune]int)
	for i, r := range stream {
		if len(container) < l {
			container = append(container, r)
			dict[r] = dict[r] + 1
		} else if isUnique(dict) {
			return i
		} else {
			removed, container = container[0], container[1:]
			dict[removed] -= 1
			container = append(container, r)
			dict[r] = dict[r] + 1
		}
	}
	return -1
}

func isUnique(dict map[rune]int) bool {
	for _, v := range dict {
		if v > 1 {
			return false
		}
	}
	return true
}
