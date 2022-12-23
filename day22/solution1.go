package day22

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

var limitY = 0
var limitX = 0
var minXs = []int{}
var maxXs = []int{}
var minYs = []int{}
var maxYs = []int{}

var globalPlayer = &Player{}

// 61338 too low
func Sol1() (ans int) {
	fmt.Println("Starting Day22 Solution1...")
	raw := utils.ReadFile("./day22/input.txt")
	// raw := utils.ReadFile("./day22/example.txt")
	raw = raw[:len(raw)-1]
	cube, insts, player := parseRaw(raw)
	ans = resolve(cube, insts, player)
	return ans
}

func resolve(cube map[int]map[int]int, insts []*Inst, player *Player) (ans int) {
	for i, inst := range insts {
		_ = i
		// move
		for d := inst.distance; d > 0; d-- {
			if !player.move() {
				break
			}
		}
		// turn
		player.turn(inst.turn)
		// fmt.Println(inst)
		// debug(cube)
		// if i == 3 {
		// 	break
		// }

	}
	ans += 1000 * (player.y + 1)
	ans += 4 * (player.x + 1)
	ans += player.dirScore()
	return ans
}

func (p *Player) dirScore() int {
	if p.dirX == 1 && p.dirY == 0 {
		return 0
	} else if p.dirX == 0 && p.dirY == 1 {
		return 1
	} else if p.dirX == -1 && p.dirY == 0 {
		return 2
	} else if p.dirX == 0 && p.dirY == -1 {
		return 3
	}
	panic("wrong direction")
}

func (p *Player) turn(command string) {
	switch command {
	case "R":
		if p.dirX == 1 && p.dirY == 0 {
			p.dirX = 0
			p.dirY = 1
		} else if p.dirX == 0 && p.dirY == 1 {
			p.dirX = -1
			p.dirY = 0
		} else if p.dirX == -1 && p.dirY == 0 {
			p.dirX = 0
			p.dirY = -1
		} else if p.dirX == 0 && p.dirY == -1 {
			p.dirX = 1
			p.dirY = 0
		}
	case "L":
		if p.dirX == 1 && p.dirY == 0 {
			p.dirX = 0
			p.dirY = -1
		} else if p.dirX == 0 && p.dirY == -1 {
			p.dirX = -1
			p.dirY = 0
		} else if p.dirX == -1 && p.dirY == 0 {
			p.dirX = 0
			p.dirY = 1
		} else if p.dirX == 0 && p.dirY == 1 {
			p.dirX = 1
			p.dirY = 0
		}
	case "":
		return
	}
}

func (p *Player) move() bool {
	nextPosX, nextPosY := p.getNextPos()
	switch p.cube[nextPosY][nextPosX] {
	case 1:
		p.x += p.dirX
		p.y += p.dirY
		return true
	case 2:
		return false
	case 0:
		x, y := p.getLoopPos()
		switch p.cube[y][x] {
		case 1:
			p.x, p.y = x, y
			return true
		case 2:
			return false
		}
		return true
	}
	return false
}

func (p *Player) getLoopPos() (int, int) {
	if p.dirX == 1 && p.dirY == 0 {
		return minXs[p.y], p.y
	} else if p.dirX == 0 && p.dirY == 1 {
		return p.x, minYs[p.x]
	} else if p.dirX == -1 && p.dirY == 0 {
		return maxXs[p.y], p.y
	} else if p.dirX == 0 && p.dirY == -1 {
		return p.x, maxYs[p.x]
	}
	panic("can't get LoopPos")
}

func (p *Player) getNextPos() (int, int) {
	return p.x + p.dirX, p.y + p.dirY
}

func parseRaw(raw []string) (map[int]map[int]int, []*Inst, *Player) {
	player := &Player{dirX: 1, x: 99999}
	rawCube := raw[:len(raw)-2]
	limitY = len(rawCube)
	cube := parseCube(rawCube, player)
	rawInst := raw[len(raw)-1]
	insts := parseInst(rawInst)
	for i := range cube[0] {
		player.x = utils.Min(player.x, i)
	}
	globalPlayer = player
	player.cube = cube
	// debug(cube)
	// debugInst(insts)
	// fmt.Println("minXs: ",len(minXs), minXs)
	// fmt.Println("maxXs: ",len(maxXs), maxXs)
	// fmt.Println("minYs: ",len(minYs), minYs)
	// fmt.Println("maxYs: ",len(maxYs), maxYs)
	return cube, insts, player
}

func parseInst(rawInst string) []*Inst {
	insts := []*Inst{}
	temp := ""
	for _, r := range rawInst {
		if string(r) != "L" && string(r) != "R" {
			temp += string(r)
		} else {
			insts = append(insts, &Inst{utils.StringToInt(temp), string(r)})
			temp = ""
		}
	}
	insts = append(insts, &Inst{utils.StringToInt(temp), ""})
	return insts
}

func parseCube(rawCube []string, player *Player) map[int]map[int]int {
	cube := make(map[int]map[int]int)
	for i, line := range rawCube {
		minX, maxX := 99999, 0
		cube[i] = make(map[int]int)
		for j, r := range line {
			limitX = utils.Max(limitX, j)
			if string(r) != " " {
				minX = utils.Min(minX, j)
				maxX = utils.Max(maxX, j)
			}
			switch string(r) {
			case ".":
				cube[i][j] = 1
			case "#":
				cube[i][j] = 2
			}
		}
		minXs = append(minXs, minX)
		maxXs = append(maxXs, maxX)
	}
	player.minXs = minXs
	player.maxXs = maxXs

	for j := 0; j < limitX; j++ {
		minY, maxY := 99999, 0
		for i := 0; i < limitY; i++ {
			if cube[i][j] != 0 {
				minY = utils.Min(minY, i)
				maxY = utils.Max(maxY, i)
			}
		}
		minYs = append(minYs, minY)
		maxYs = append(maxYs, maxY)
	}
	player.minYs = minYs
	player.maxYs = maxYs
	return cube
}

type Player struct {
	x     int
	y     int
	dirX  int
	dirY  int
	cube  map[int]map[int]int
	minXs []int
	maxXs []int
	minYs []int
	maxYs []int
}

type Inst struct {
	distance int
	turn     string
}

func debugInst(insts []*Inst) {
	for _, inst := range insts {
		fmt.Println(inst)
	}
}

func debug(cube map[int]map[int]int) {
	for i := 0; i < limitY; i++ {
		for j := 0; j < limitX; j++ {
			if globalPlayer.y == i && globalPlayer.x == j {
				fmt.Print("P")
			} else {
				fmt.Print(cube[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
