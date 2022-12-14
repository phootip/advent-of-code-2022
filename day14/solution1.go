package day14

import (
	"fmt"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day14 Solution1...")
	raw := utils.ReadFile("./day14/input.txt")
	// raw := utils.ReadFile("./day14/example.txt")
	rockMap := rawToRockMap(raw)
	// debug(rockMap)
	ans := simulate(rockMap)
	// fmt.Println()
	// debug(rockMap)
	return ans
}

func rawToRockMap(raw []string) map[int]map[int]int {
	rockMap := make(map[int]map[int]int)
	for _, line := range raw {
		if line == "" {
			break
		}
		pointsString := strings.Split(line, " -> ")
		points := pointsStringToInt(pointsString)
		for i := 0; i < len(points)-1; i++ {
			rockMap = plot(rockMap, points[i], points[i+1])
		}
	}
	return rockMap
}

func debug(rockMap map[int]map[int]int) {
	for y := 0; y < 200; y++ {
		for x := 450; x < 600; x++ {
			fmt.Print(rockMap[y][x])
		}
		fmt.Println(" : ", y)
	}
}

func pointsStringToInt(pointsString []string) [][]int {
	points := make([][]int, len(pointsString))
	for i, pointString := range pointsString {
		points[i] = utils.SToInt(strings.Split(pointString, ","))
	}
	return points
}

func plot(rockMap map[int]map[int]int, source []int, des []int) map[int]map[int]int {
	x1, y1 := source[0], source[1]
	x2, y2 := des[0], des[1]
	if x1 != x2 {
		dir := (x2 - x1) / abs(x2-x1)
		for x := x1; x != x2 + dir; x += dir {
			if rockMap[y1] == nil {
				rockMap[y1] = make(map[int]int)
			}
			rockMap[y1][x] = 1
		}
	}
	if y1 != y2 {
		dir := (y2 - y1) / abs(y2-y1)
		for y := y1; y != y2 + dir; y += dir {
			if rockMap[y] == nil {
				rockMap[y] = make(map[int]int)
			}
			rockMap[y][x1] = 1
		}
	}
	return rockMap
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func simulate(rockMap map[int]map[int]int) (ans int) {
	x, y := 500, 0
	for {
		if rockMap[y+1][x] == 0 {
			y++
		} else if rockMap[y+1][x-1] == 0 {
			x--
			y++
		} else if rockMap[y+1][x+1] == 0 {
			x++
			y++
		} else {
			if rockMap[y] == nil {
				rockMap[y] = make(map[int]int)
			}
			rockMap[y][x] = 2
			x, y = 500, 0
			ans += 1
		}
		// debug(rockMap)
		// fmt.Println(x,y)
		// fmt.Println()
		if y > 1000 {
			return ans
		}
		// if ans == 21 {
		// 	return ans
		// }
	}
}
