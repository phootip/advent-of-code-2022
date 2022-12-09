package day9

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

// 2225 too high
func Sol2() int {
	fmt.Println("Starting Day9 Solution2...")
	raw := utils.ReadFile("./day9/input.txt")
	// raw := utils.ReadFile("./day9/example.txt")
	// raw := utils.ReadFile("./day9/example2.txt")
	data := genData(raw)
	return simulate2(data)
}

func simulate2(data []pair) int {
	knot := [10]point{}
	head := &knot[0]
	tail := &knot[9]
	visited := map[point]bool{}
	for _, inst := range data {
		// fmt.Println(inst)
		for i := inst.distance; i > 0; i-- {
			move(head, inst.direction)
			for k := 0; k < 9; k++ {
				updateTail(&knot[k],&knot[k+1])
			}
			// fmt.Println("head at: ", head)
			// fmt.Println("tail at: ", tail)
			visited[*tail] = true
		}
	}
	return len(visited)
}
