package day18

import (
	"fmt"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() (ans int) {
	fmt.Println("Starting Day18 Solution1...")
	raw := utils.ReadFile("./day18/input.txt")
	// raw := utils.ReadFile("./day18/example.txt")
	raw = raw[:len(raw)-1]
	coords := rawToCoords(raw)
	for _, coord := range coords {
		fmt.Println(coord)
	}
	ans = countSides(coords)
	return ans
}

// func rawToCoords(raw []string) (coords []Coordinate) {
func rawToCoords(raw []string) (coords [][3]int) {
	for _, line := range raw {
		coor := utils.SToInt(strings.Split(line, ","))
		// coords = append(coords, Coordinate{coor[0], coor[1], coor[2], 6})
		arr := [3]int{}
		copy(arr[:], coor)
		coords = append(coords, arr)
	}
	return coords
}

func countSides(coords [][3]int) (ans int) {
	// mem := [][3]int{}
	for i, coor := range coords {
		ans += 6
		for _, coor2 := range coords[:i] {
			if adjacent(coor, coor2) {
				ans -= 2
			}
		}
	}
	return ans
}

func adjacent(coor1 [3]int, coor2 [3]int) bool {
	// fmt.Println("Check adj: ", coor1, coor2)
	// at least2 side equal
	equalSide := 0
	nonEqualIdx := -1
	for i := range coor1 {
		if coor1[i] == coor2[i] {
			equalSide++
		} else {
			nonEqualIdx = i
		}
	}
	if equalSide == 2 && utils.Abs(coor1[nonEqualIdx] - coor2[nonEqualIdx]) == 1 {
		return true
	}
	return false
}

type Coordinate struct {
	x int
	y int
	z int
	surface int
}
