package day18

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

// 4266 too high
func Sol2() (ans int) {
	fmt.Println("Starting Day18 Solution2...")
	raw := utils.ReadFile("./day18/input.txt")
	// raw := utils.ReadFile("./day18/example.txt")
	raw = raw[:len(raw)-1]
	coords := rawToCoords(raw)
	ans = countSides2(coords)
	return ans
}

func countSides2(coords [][3]int) (ans int) {
	// mem := [][][3]int{}
	for i, coor := range coords {
		ans += 6
		for _, coor2 := range coords[:i] {
			if adjacent(coor, coor2) {
				ans -= 2
			}
		}
		trappedSides := checkTrapped(coor, coords[:i], coords)
		ans -= trappedSides*6

	}
	return ans
}

func checkTrapped(coord1 [3]int, coords [][3]int, universeCoords [][3]int) int {
	// fmt.Println("checkTrapped for: ", coord1)
	// fmt.Println("List to check: ", coords)

	// check for one side trapped
	// checking Z
	// check no adjacent
	result := 0
	for _, axis := range [3]int{0, 1, 2} {
		for _, dir := range [2]int{-2, 2} {
			if checkTrappedAxis(coord1, coords, universeCoords, axis, dir) {
				result += 1
			}
			// fmt.Println("result: ", result)
		}
	}
	return result
}

func checkTrappedAxis(coord1 [3]int, coords [][3]int, universeCoords [][3]int, axis int, dir int) bool {
	// check adj
	adj := coord1
	adj[axis] += dir / utils.Abs(dir)
	if coordExist(adj, universeCoords) {
		return false
	}
	// fmt.Println("checkTrapped for: ", coord1)
	temp := coord1
	temp[axis] += dir
	if !coordExist(temp, coords) {
		return false
	}
	for _, axis2 := range [3]int{0, 1, 2} {
		if axis == axis2 {
			continue
		}
		for _, offset := range [2]int{-1, 1} {
			temp2 := adj
			temp2[axis2] += offset
			// if coord1 == [3]int{2, 2, 3} {
			// 	fmt.Println("temp: ", temp)
			// 	fmt.Println("debug temp2: ", temp2)
			// }
			if !coordExist(temp2, coords) {
				return false
			}
		}
	}
	return true
}

func coordExist(coord1 [3]int, coords [][3]int) bool {
	for _, coord2 := range coords {
		if coord1 == coord2 {
			return true
		}
	}
	return false
}
