package day12

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day12 Solution1...")
	raw := utils.ReadFile("./day12/input.txt")
	// raw := utils.ReadFile("./day12/example.txt")
	pointMap, start, end := rawToMap(raw)
	ans := bfs(pointMap, start, end)
	return ans
}

func bfs(pointMap [][]rune, start path, end path) int {
	queue := []path{start}
	visited := map[path]bool{start: true}
	// fmt.Println(pointMap)
	fmt.Println("start, end: ",start, end)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x, y := node.x, node.y
		c := pointMap[y][x]
		if x == end.x && y == end.y {
			return node.depth
		}

		if y-1 >= 0 && !visited[path{x: x, y: y - 1}] && c-pointMap[y-1][x] >= -1 {
			visited[path{x: x, y: y - 1}] = true
			queue = append(queue, path{x, y - 1, node.depth + 1})
		}
		if y+1 < len(pointMap) && !visited[path{x: x, y: y + 1}] && c-pointMap[y+1][x] >= -1 {
			visited[path{x: x, y: y + 1}] = true
			queue = append(queue, path{x, y + 1, node.depth + 1})
		}
		if x-1 >= 0 && !visited[path{x: x - 1, y: y}] && c-pointMap[y][x-1] >= -1 {
			visited[path{x: x - 1, y: y}] = true
			queue = append(queue, path{x - 1, y, node.depth + 1})
		}
		if x+1 < len(pointMap[0]) && !visited[path{x: x + 1, y: y}] && c-pointMap[y][x+1] >= -1 {
			visited[path{x: x + 1, y: y}] = true
			queue = append(queue, path{x + 1, y, node.depth + 1})
		}

		// fmt.Println("queue: ", queue)
		// if loop == 2 {
		// 	break
		// }
	}

	return 0
}

type path struct {
	x, y  int
	depth int
}

type point struct {
	x, y int
}

func rawToMap(raw []string) (pointMap [][]rune, start path, end path) {
	pointMap = make([][]rune, len(raw)-1)
	for i, line := range raw {
		if line == "" {
			break
		}
		temp := make([]rune, len(line))
		for j, r := range line {
			c := string(r)
			if c == "S" {
				temp[j] = []rune("a")[0]
				start.x = j
				start.y = i
			} else if c == "E" {
				temp[j] = []rune("z")[0]
				end.x = j
				end.y = i
			} else {
				temp[j] = r
			}
		}
		pointMap[i] = temp
	}
	return pointMap, start, end
}
