package day8

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol2() int {
	fmt.Println("Starting Day8 Solution2...")
	raw := utils.ReadFile("./day8/input.txt")
	// raw := utils.ReadFile("./day8/example.txt")
	treeMap := genMap(raw)
	// fmt.Println(treeMap)
	return bestScore(treeMap)
}

func bestScore(treeMap [][]int) int {
	best := 0
	for y := 0; y < len(treeMap); y++ {
		for x := 0; x < len(treeMap[0]); x++ {
			best = max(best,calScore(treeMap, x, y))
		}
	}

	return best
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func calScore(treeMap [][]int, x int, y int) int {
	ans := countUp(treeMap, x, y)
	ans *= countDown(treeMap, x, y)
	ans *= countLeft(treeMap, x, y)
	ans *= countRight(treeMap, x, y)
	return ans
}

func countUp(treeMap [][]int, x int, y int) int {
	result := 0
	treeHeight := treeMap[y][x]
	for k := y - 1; k >= 0; k-- {
		tempHeight := treeMap[k][x]
		if tempHeight >= treeHeight {
			return result + 1
		} else {
			result += 1
		}
	}
	return result
}

func countDown(treeMap [][]int, x int, y int) int {
	result := 0
	treeHeight := treeMap[y][x]
	for k := y + 1; k < len(treeMap); k++ {
		tempHeight := treeMap[k][x]
		if tempHeight >= treeHeight {
			return result + 1
		} else {
			result += 1
		}
	}
	return result
}

func countLeft(treeMap [][]int, x int, y int) int {
	result := 0
	treeHeight := treeMap[y][x]
	for k := x - 1; k >= 0; k-- {
		tempHeight := treeMap[y][k]
		if tempHeight >= treeHeight {
			return result + 1
		} else {
			result += 1
		}
	}
	return result
}

func countRight(treeMap [][]int, x int, y int) int {
	result := 0
	treeHeight := treeMap[y][x]
	for k := x + 1; k < len(treeMap[0]); k++ {
		tempHeight := treeMap[y][k]
		if tempHeight >= treeHeight {
			return result + 1
		} else {
			result += 1
		}
	}
	return result
}
