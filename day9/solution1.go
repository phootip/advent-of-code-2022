package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

// 2225 too high
func Sol1() int {
	fmt.Println("Starting Day9 Solution1...")
	raw := utils.ReadFile("./day9/input.txt")
	// raw := utils.ReadFile("./day9/example.txt")
	data := genData(raw)
	return simulate(data)
}

func simulate(data []pair) int {
	head := point{0, 0}
	tail := point{0, 0}
	visited := map[point]bool{}
	for _, inst := range data {
		// fmt.Println(inst)
		for i := inst.distance; i > 0; i-- {
			move(&head, inst.direction)
			updateTail(&head, &tail)
			// fmt.Println("head at: ", head)
			// fmt.Println("tail at: ", tail)
			visited[tail] = true
		}
	}
	return len(visited)
}

type pair struct {
	direction string
	distance  int
}
type point struct {
	x int
	y int
}

func move(p *point, direction string) {
	switch direction {
	case "U":
		p.y -= 1
	case "D":
		p.y += 1
	case "L":
		p.x -= 1
	case "R":
		p.x += 1
	}
}

func updateTail(head *point, tail *point) {
	if abs(head.x-tail.x) <= 1 && abs(head.y-tail.y) <= 1 {
		return
	}
	if head.y != tail.y {
		tail.y += (head.y - tail.y)/abs(head.y - tail.y)
	}
	if head.x != tail.x {
		tail.x += (head.x - tail.x)/abs(head.x - tail.x)
	}
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func genData(raw []string) []pair {
	data := []pair{}
	for _, line := range raw {
		if line == "" {
			return data
		}
		temp := strings.Split(line, " ")
		num, err := strconv.Atoi(temp[1])
		utils.Check(err)
		data = append(data, pair{temp[0], num})
	}
	return data
}
