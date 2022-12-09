package day8

import (
	"fmt"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

// 2225 too high
func Sol1() int {
	fmt.Println("Starting Day8 Solution1...")
	raw := utils.ReadFile("./day8/input.txt")
	// raw := utils.ReadFile("./day8/example.txt")
	treeMap := genMap(raw)
	// fmt.Println(treeMap)
	return countVisible(treeMap)
}

func genMap(raw []string) (treeMap [][]int) {
	treeMap = [][]int{}
	for _, line := range raw {
		if line == "" {
			return treeMap
		}
		row := utils.SToInt(strings.Split(line, ""))
		treeMap = append(treeMap, row)
	}
	return treeMap
}

func countVisible(treeMap [][]int) int {
	ans := len(treeMap)*2 + len(treeMap[0])*2 - 4
	fmt.Println(ans)
	for y := 1; y < len(treeMap)-1; y++ {
		for x := 1; x < len(treeMap[0])-1; x++ {
			ans += isVisible(&Payload{treeMap, x, y})
		}
	}
	return ans
}

func isVisible(payload *Payload) int {
	if isVisibleUp(payload) || isVisibleDown(payload) || isVisibleLeft(payload) || isVisibleRight(payload) {
		return 1
	}
	return 0
}

type Payload struct {
	treeMap [][]int
	x       int
	y       int
}

func isVisibleUp(payload *Payload) bool {
	treeMap := payload.treeMap
	x := payload.x
	y := payload.y
	treeHeight := treeMap[y][x]
	for k := y - 1; k >= 0; k-- {
		tempHeight := treeMap[k][x]
		if tempHeight >= treeHeight {
			return false
		}
	}
	return true
}

func isVisibleDown(payload *Payload) bool {
	treeMap := payload.treeMap
	x := payload.x
	y := payload.y
	treeHeight := treeMap[y][x]
	for k := y + 1; k < len(treeMap); k++ {
		tempHeight := treeMap[k][x]
		if tempHeight >= treeHeight {
			return false
		}
	}
	return true
}

func isVisibleLeft(payload *Payload) bool {
	treeMap := payload.treeMap
	x := payload.x
	y := payload.y
	treeHeight := treeMap[y][x]
	for k := x - 1; k >= 0; k-- {
		tempHeight := treeMap[y][k]
		if tempHeight >= treeHeight {
			return false
		}
	}
	return true
}

func isVisibleRight(payload *Payload) bool {
	treeMap := payload.treeMap
	x := payload.x
	y := payload.y
	treeHeight := treeMap[y][x]
	for k := x + 1; k < len(treeMap[0]); k++ {
		tempHeight := treeMap[y][k]
		if tempHeight >= treeHeight {
			return false
		}
	}
	return true
}
