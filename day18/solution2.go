package day18

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

// 4266 too high
// 2024 too low
func Sol2() (ans int) {
	fmt.Println("Starting Day18 Solution2...")
	raw := utils.ReadFile("./day18/input.txt")
	// raw := utils.ReadFile("./day18/example.txt")
	raw = raw[:len(raw)-1]
	coords := rawToCoords(raw)
	ans = countSides2(coords)
	return ans
}

func countSides3(coords [][3]int) (ans int) {
	minCoor, maxCoor := getMinMaxCoor(coords)
	fmt.Println(minCoor, maxCoor)

	ans += steamSurface(coords, minCoor, maxCoor)

	return ans
}

func steamSurface(coords [][3]int, minCoor [3]int, maxCoor [3]int) (ans int) {
	for _, axises := range [][3]int{{0, 1, 2}, {0, 2, 1}, {1, 2, 0}} {
		i, j, k := axises[0], axises[1], axises[2]
		for x := minCoor[i]; x <= maxCoor[i]; x++ {
			for y := minCoor[j]; y <= maxCoor[j]; y++ {
				for z := minCoor[k]; z <= maxCoor[k]; z++ {
					temp := [3]int{}
					temp[i] = x
					temp[j] = y
					temp[k] = z
					if coordExist(temp, coords) {
						fmt.Println(temp)
						ans += 2
						break
					}
				}
			}
		}
	}

	return ans
}

func getMinMaxCoor(coords [][3]int) ([3]int, [3]int) {
	minCoor := [3]int{1000, 1000, 1000}
	maxCoor := [3]int{0, 0, 0}
	for _, coord := range coords {
		for i := range [3]int{} {
			minCoor[i] = utils.Min(minCoor[i], coord[i])
			maxCoor[i] = utils.Max(maxCoor[i], coord[i])
		}
	}
	return minCoor, maxCoor
}

func countSides2(coords [][3]int) (ans int) {
	minZ, maxZ := getMinMaxZ(coords)
	grouped := groupCoordsByZ(coords, minZ, maxZ)
	minLimit, maxLimit := getMinMaxCoor(coords)
	for i := range [3]int{} {
		minLimit[i]--
		maxLimit[i]++
	}
	// fmt.Println(coords)
	// fmt.Println(grouped)
	start := grouped[minZ][0]
	start[2]--
	ans = bfs(start, coords, minLimit, maxLimit)
	return ans
}

func bfs(root [3]int, coords [][3]int, minLimit [3]int, maxLimit [3]int) int {
	// fmt.Println("root:", root)
	// fmt.Println("limit:", minLimit,maxLimit)
	var node [3]int
	queue := [][3]int{root}
	visited := map[[3]int]bool{}
	visited[root] = true
	ans := 0

	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]
		adjNodes := getAdjCoords(node, minLimit, maxLimit)
		// fmt.Println("node: ", node)
		// fmt.Println(adjNodes)

		// check surface
		for _, adjNode := range adjNodes {
			if coordExist(adjNode, coords){
				ans += 1
				continue
			}
			if visited[adjNode] {
				continue
			}
			// fmt.Println("visiting node: ", adjNode)
			queue = append(queue, adjNode)
			visited[adjNode] = true
		}
	}
	return ans

	// visited := map[int]bool{}

}

func getAdjCoords(node [3]int, minLimit [3]int, maxLimit [3]int) (adjCoords [][3]int) {
	for _, axis := range [3]int{0, 1, 2} {
		for _, dir := range [2]int{-1, 1} {
			temp := node
			temp[axis] += dir
			if temp[axis] < minLimit[axis] || temp[axis] > maxLimit[axis]{
				continue
			}
			adjCoords = append(adjCoords, temp)
		}
	}
	return adjCoords
}

func groupCoordsByZ(coords [][3]int, minZ int, maxZ int) map[int][][3]int {
	result := map[int][][3]int{}
	for z := minZ; z <= maxZ; z++ {
		result[z] = filterCoordWithZ(coords, z)
	}
	return result
}

func getMinMaxZ(coords [][3]int) (int, int) {
	minZ := 1000
	maxZ := 0
	for _, coord := range coords {
		minZ = utils.Min(minZ, coord[2])
		maxZ = utils.Max(maxZ, coord[2])
	}
	return minZ, maxZ
}

func filterCoordWithZ(coords [][3]int, z int) (result [][3]int) {
	for _, coord := range coords {
		if coord[2] == z {
			result = append(result, coord)
		}
	}
	return result
}

func checkTrapped(coord1 [3]int, coords [][3]int, universeCoords [][3]int) int {
	// fmt.Println("checkTrapped for: ", coord1)
	// fmt.Println("List to check: ", coords)
	result := 0
	for _, axis := range [3]int{0, 1, 2} {
		for _, dir := range [2]int{-2, 2} {
			if checkTrappedAxis(coord1, coords, universeCoords, axis, dir) {
				result += 1
			}
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
