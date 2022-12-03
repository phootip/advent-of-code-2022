package day1

// package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/phootip/advent-of-code-2022/utils"
)

func getMax(data []int) int {
	result := 0
	for _, v := range data {
		if result < v {
			result = v
		}
	}
	return result
}

func Sol1() int {
	fmt.Println("Starting Day1 Solution1...")
	// raw := utils.ReadFile("./day1/example1.txt")
	raw := utils.ReadFile("./day1/input.txt")
	mem := []int{0}
	// count := 0
	for _, line := range raw {
		if line != "" {
			value, err := strconv.Atoi(line)
			utils.Check(err)
			mem[len(mem)-1] += value
		} else {
			mem = append(mem, 0)
		}
	}
	sort.Ints(mem)
	return getMax(mem)
}

func Sol2() int {
	fmt.Println("Starting Day1 Solution2...")
	// raw := utils.ReadFile("./day1/example1.txt")
	raw := utils.ReadFile("./day1/input.txt")
	mem := []int{0}
	// count := 0
	for _, line := range raw {
		if line != "" {
			value, err := strconv.Atoi(line)
			utils.Check(err)
			mem[len(mem)-1] += value
		} else {
			mem = append(mem, 0)
		}
	}
	sort.Ints(mem)
	return mem[len(mem)-1] + mem[len(mem)-2] + mem[len(mem)-3]
}
