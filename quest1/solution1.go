// package quest1
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) []string {
	b, err := os.ReadFile(filename)
	s := string(b)
	check(err)
	data := strings.Split(s, "\n")
	return data

}

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
	fmt.Println("Starting sol1...")
	// raw := readFile("./example1.txt")
	raw := readFile("./input.txt")
	mem := []int{0}
	// count := 0
	for _, line := range raw {
		if line != "" {
			value, err := strconv.Atoi(line)
			check(err)
			mem[len(mem)-1] += value
		} else {
			mem = append(mem, 0)
		}
	}
	sort.Ints(mem)
	return getMax(mem)
}

func Sol2() int {
	fmt.Println("Starting sol2...")
	// raw := readFile("./example1.txt")
	raw := readFile("./input.txt")
	mem := []int{0}
	// count := 0
	for _, line := range raw {
		if line != "" {
			value, err := strconv.Atoi(line)
			check(err)
			mem[len(mem)-1] += value
		} else {
			mem = append(mem, 0)
		}
	}
	sort.Ints(mem)
	return mem[len(mem)-1] + mem[len(mem)-2] + mem[len(mem)-3]
}
