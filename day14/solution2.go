package day14

import (
	"fmt"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol2() int {
	fmt.Println("Starting Day14 Solution2...")
	raw := utils.ReadFile("./day14/input.txt")
	// raw := utils.ReadFile("./day14/example.txt")
	rockMap := rawToRockMap2(raw)
	// debug(rockMap)
	ans := simulate2(rockMap)
	// fmt.Println()
	// debug(rockMap)
	return ans
}

func simulate2(rockMap map[int]map[int]int) (ans int) {
	x, y := 500, 0
	for {
		if rockMap[0][500] != 0 {
			return ans
		}
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
		if y > 300 {
			panic("Y go through floor")
		}
	}
}

func rawToRockMap2(raw []string) map[int]map[int]int {
	rockMap := make(map[int]map[int]int)
	maxY := 0
	for _, line := range raw {
		if line == "" {
			break
		}
		pointsString := strings.Split(line, " -> ")
		points := pointsStringToInt(pointsString)
		for i := 0; i < len(points)-1; i++ {
			rockMap = plot(rockMap, points[i], points[i+1])
		}
		for _, point := range points {
			maxY = utils.Max(maxY, point[1])
		}
	}
	maxY += 2
	rockMap[maxY] = make(map[int]int)
	for x := 0; x < 1000; x++ {
		rockMap[maxY][x] = 1
	}
	return rockMap
}
