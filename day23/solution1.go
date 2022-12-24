package day23

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

// 3963 too high
func Sol1() (ans int) {
	fmt.Println("Starting Day23 Solution1...")
	raw := utils.ReadFile("./day23/input.txt")
	// raw := utils.ReadFile("./day23/example.txt")
	// raw := utils.ReadFile("./day23/example2.txt")
	raw = raw[:len(raw)-1]
	game := parseRaw(raw)
	_ = game
	for i := 1; i < 11; i++ {
		// fmt.Println("round:", i)
		game.simulation(i)
		game.updateDirOrder()
		// game.debug()
	}
	game.updateMinMaxXY()
	// game.debug()
	ans = game.emptyGround()
	return ans
}

func (g *Game) emptyGround() (ans int) {
	for i := g.minY; i <= g.maxY; i++ {
		for j := g.minX; j <= g.maxX; j++ {
			switch g.land[i][j] {
			case 0:
				ans += 1
			case 1:
				ans += 1
			}
		}
	}
	return ans
}

func (g *Game) updateMinMaxXY() {
	minY, maxY := 99999, 0
	minX, maxX := 99999, 0
	for k, v := range g.land {
		for k2, v2 := range v {
			if v2 == 2 {
				minY = utils.Min(minY, k)
				maxY = utils.Max(maxY, k)
				minX = utils.Min(minX, k2)
				maxX = utils.Max(maxX, k2)
			}

		}
	}
	g.minX, g.maxX = minX, maxX
	g.minY, g.maxY = minY, maxY
}

func (g *Game) simulation(round int) {
	g.reservedLand = make(map[int]map[int]int)
	// ReserveLand
	for i := g.minY; i <= g.maxY; i++ {
		for j := g.minX; j <= g.maxX; j++ {
			if g.land[i][j] == 2 {
				x, y := g.nextLand(round, j, i)
				g.reserve(x, y)
			}
		}
	}
	newLand := g.copyLand()
	// Move
	for i := g.minY; i <= g.maxY; i++ {
		for j := g.minX; j <= g.maxX; j++ {
			if g.land[i][j] == 2 {
				x, y := g.nextLand(round, j, i)
				if g.reservedLand[y][x] == 3 {
					newLand[i][j] = 1
					if newLand[y] == nil {
						newLand[y] = make(map[int]int)
					}
					newLand[y][x] = 2
					g.minX = utils.Min(g.minX, x)
					g.maxX = utils.Max(g.maxX, x)
					g.minY = utils.Min(g.minY, y)
					g.maxY = utils.Max(g.maxY, y)
				}
			}
		}
	}
	g.land = newLand
	// g.debugR()
	// g.debug()
}

func (g *Game) copyLand() map[int]map[int]int {
	newLand := make(map[int]map[int]int)
	for k, v := range g.land {
		newLand[k] = make(map[int]int)
		for k2, v2 := range v {
			newLand[k][k2] = v2
		}
	}
	return newLand

}

func (g *Game) reserve(x int, y int) {
	if g.reservedLand[y] == nil {
		g.reservedLand[y] = make(map[int]int)
	}
	if g.reservedLand[y][x] != 3 {
		g.reservedLand[y][x] = 3
	} else {
		g.reservedLand[y][x] = 4
	}
}

func (g *Game) updateDirOrder() {
	g.directionOrder = append(g.directionOrder[1:], g.directionOrder[0])
}

func (g *Game) nextLand(round int, x int, y int) (int, int) {
	// check no elf
	if g.noElf(x, y) {
		return x, y
	}
	for _, direction := range g.directionOrder {
		switch direction {
		case "N":
			if g.land[y-1][x-1] != 2 && g.land[y-1][x] != 2 && g.land[y-1][x+1] != 2 {
				return x, y - 1
			}
		case "S":
			if g.land[y+1][x-1] != 2 && g.land[y+1][x] != 2 && g.land[y+1][x+1] != 2 {
				return x, y + 1
			}
		case "W":
			if g.land[y-1][x-1] != 2 && g.land[y][x-1] != 2 && g.land[y+1][x-1] != 2 {
				return x - 1, y
			}
		case "E":
			if g.land[y-1][x+1] != 2 && g.land[y][x+1] != 2 && g.land[y+1][x+1] != 2 {
				return x + 1, y
			}
		}
	}
	return x, y
}

func (g *Game) noElf(x int, y int) bool {
	for _, i := range [3]int{-1, 0, 1} {
		if g.land[y-1][x+i] == 2 || g.land[y+1][x+i] == 2 {
			return false
		}
	}
	if g.land[y][x+1] == 2 || g.land[y][x-1] == 2 {
		return false
	}
	return true
}

func parseRaw(raw []string) *Game {
	g := &Game{}
	g.land = make(map[int]map[int]int)
	g.reservedLand = make(map[int]map[int]int)
	g.directionOrder = []string{"N", "S", "W", "E"}
	g.minX, g.minY = 0, 0
	g.maxX, g.maxY = len(raw[0]), len(raw)
	for i, line := range raw {
		g.reservedLand[i] = make(map[int]int)
		g.land[i] = make(map[int]int)
		for j, r := range line {
			switch string(r) {
			case ".":
				g.land[i][j] = 1
			case "#":
				g.land[i][j] = 2
			}
		}
	}
	return g
}

func (g *Game) debug() {
	for i := g.minY; i <= g.maxY; i++ {
		for j := g.minX; j <= g.maxX; j++ {
			switch g.land[i][j] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print(".")
			case 2:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func (g *Game) debugR() {
	for i := g.minY; i < g.maxY; i++ {
		for j := g.minX; j < g.maxX; j++ {
			switch g.reservedLand[i][j] {
			case 0:
				fmt.Print(".")
			case 2:
				fmt.Print("#")
			case 3:
				fmt.Print("R")
			case 4:
				fmt.Print("C")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type Game struct {
	land           map[int]map[int]int
	reservedLand   map[int]map[int]int
	directionOrder []string
	minX, minY     int
	maxX, maxY     int
}
