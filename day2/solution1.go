package day2

import (
	"fmt"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day2 Solution1...")
		shapeScore  := map[string]int{"X": 1, "Y": 2, "Z": 3}
	resultScore := map[string]map[string]int{"A": {"X": 3, "Y": 6, "Z": 0},"B": {"X": 0, "Y": 3, "Z": 6},"C": {"X": 6, "Y": 0, "Z": 3}}

	raw := utils.ReadFile("./day2/input.txt")
	// raw := utils.ReadFile("./day2/example1.txt")
	result := 0
	for _, line := range raw {
		if line == "" {
			break
		}
		data := strings.Split(line, " ")
		// fmt.Println("shape score", shapeScore[data[1]])
		// fmt.Println("result score", resultScore[data[0]][data[1]])
		result += shapeScore[data[1]] + resultScore[data[0]][data[1]]
	}
	return result
}

func Sol2() int {
	fmt.Println("Starting Day2 Solution2...")
	resultScore  := map[string]int{"X": 0, "Y": 3, "Z": 6}
	shapeScore := map[string]map[string]int{"A": {"X": 3, "Y": 1, "Z": 2},"B": {"X": 1, "Y": 2, "Z": 3},"C": {"X": 2, "Y": 3, "Z": 1}}

	raw := utils.ReadFile("./day2/input.txt")
	// raw := utils.ReadFile("./day2/example1.txt")
	result := 0
	for _, line := range raw {
		if line == "" {
			break
		}
		// fmt.Println(line)
		data := strings.Split(line, " ")
		// fmt.Println(data)
		// fmt.Println("shape score", shapeScore[data[0]][data[1]])
		// fmt.Println("result score", resultScore[data[1]])
		result += resultScore[data[1]] + shapeScore[data[0]][data[1]]
	}
	return result
}
