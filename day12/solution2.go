package day12

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol2() int {
	fmt.Println("Starting Day12 Solution2...")
	raw := utils.ReadFile("./day12/input.txt")
	// raw := utils.ReadFile("./day12/example.txt")
	pointMap, end, start := rawToMap(raw)
	ans := bfs2(pointMap, start, end)
	return ans
}

func bfs2(pointMap [][]rune, start path, end path) int {
	queue := []path{start}
	visited := map[path]bool{start: true}
	_, _ = queue, visited
	// fmt.Println(pointMap)
	fmt.Println("start, end: ",start, end)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x, y := node.x, node.y
		c := pointMap[y][x]
		if c == []rune("a")[0] {
			return node.depth
		}

		if y-1 >= 0 && !visited[path{x: x, y: y - 1}] && c-pointMap[y-1][x] <= 1 {
			visited[path{x: x, y: y - 1}] = true
			queue = append(queue, path{x, y - 1, node.depth + 1})
		}
		if y+1 < len(pointMap) && !visited[path{x: x, y: y + 1}] && c-pointMap[y+1][x] <= 1 {
			visited[path{x: x, y: y + 1}] = true
			queue = append(queue, path{x, y + 1, node.depth + 1})
		}
		if x-1 >= 0 && !visited[path{x: x - 1, y: y}] && c-pointMap[y][x-1] <= 1 {
			visited[path{x: x - 1, y: y}] = true
			queue = append(queue, path{x - 1, y, node.depth + 1})
		}
		if x+1 < len(pointMap[0]) && !visited[path{x: x + 1, y: y}] && c-pointMap[y][x+1] <= 1 {
			visited[path{x: x + 1, y: y}] = true
			queue = append(queue, path{x + 1, y, node.depth + 1})
		}
	}

	return 0
}
