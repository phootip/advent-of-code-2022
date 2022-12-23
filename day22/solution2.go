package day22

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

var cubeDimension = map[int]map[int]bool{}

func Sol2() (ans int) {
	fmt.Println("Starting Day22 Solution2...")
	raw := utils.ReadFile("./day22/input.txt")
	cubeSize := 50
	// raw := utils.ReadFile("./day22/example.txt")
	// cubeSize := 4
	raw = raw[:len(raw)-1]
	cube, insts, player := parseRaw(raw)
	player.cubeSize = cubeSize
	player.initCubeDimension()
	debug2(cube)
	ans = resolve2(cube, insts, player)
	fmt.Println("after")
	debug2(cube)
	return ans
}

func (p *Player) initCubeDimension() {
	for i := 0; i < 4; i++ {
		// fmt.Println(i * p.cubeSize)
		cubeDimension[i] = make(map[int]bool)
		for j := 0; j < p.cubeSize; j++ {
			// fmt.Println(j*p.cubeSize)
			if p.cube[i*p.cubeSize][j*p.cubeSize] != 0 {
				cubeDimension[i][j] = true
			}
		}
		fmt.Println()
	}
	fmt.Println(cubeDimension)
}

func resolve2(cube map[int]map[int]int, insts []*Inst, player *Player) (ans int) {
	for i, inst := range insts {
		_ = i
		// move
		for d := inst.distance; d > 0; d-- {
			if !player.move2() {
				break
			}
		}
		// turn
		player.turn(inst.turn)
		fmt.Println(inst)
		fmt.Println(len(insts))
		fmt.Println(i)
		// debug2(player.cube)
	}
	ans += 1000 * (player.y + 1)
	ans += 4 * (player.x + 1)
	ans += player.dirScore()
	return ans
}

func (p *Player) move2() bool {
	nextPosX, nextPosY := p.getNextPos()
	switch p.cube[nextPosY][nextPosX] {
	case 1:
		p.x += p.dirX
		p.y += p.dirY
		return true
	case 2:
		return false
	case 0:
		// x, y := p.getLoopPos2()
		// x, y, dirX, dirY := p.getNextSection()
		x, y, dirX, dirY := p.getNextSection2()
		switch p.cube[y][x] {
		case 1:
			p.x, p.y = x, y
			p.dirX, p.dirY = dirX, dirY
			return true
		case 2:
			return false
		}
	}
	fmt.Println(p.cube[nextPosY][nextPosX])
	panic("move did not return")
}

func (p *Player) getLoopPos2() (int, int) {
	fmt.Println("cubeSize", p.cubeSize)
	fmt.Println("x,y:", p.x, p.y)
	// x, y, newDirX, newDirY := p.getNextSection()
	debug2(p.cube)
	panic("debug")
	if p.dirX == 1 && p.dirY == 0 {
		return p.minXs[p.y], p.y
	} else if p.dirX == 0 && p.dirY == 1 {
		return p.x, p.minYs[p.x]
	} else if p.dirX == -1 && p.dirY == 0 {
		return p.maxXs[p.y], p.y
	} else if p.dirX == 0 && p.dirY == -1 {
		return p.x, p.maxYs[p.x]
	}
	panic("can't get LoopPos")
}

func (p *Player) getNextSection2() (int, int, int, int) {
	secX, secY := p.x/p.cubeSize, p.y/p.cubeSize
	// get relative
	x := p.x - secX*p.cubeSize
	y := p.y - secY*p.cubeSize
	fmt.Println("dir x,y:", p.dirX, p.dirY)
	fmt.Println("section x,y:", secX, secY)
	fmt.Println("relative x,y:", x, y)
	if p.dirX == 1 && p.dirY == 0 {
		if secX == 2 && secY == 0 {
			newSecX, newSecY := 0, 3
			x, y = y, x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 0, 1
			return x, y, dirX, dirY
		}
		if secX == 1 && secY == 1 {
			newSecX, newSecY := 2, 0
			x, y = y, x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 0, -1
			return x, y, dirX, dirY
		}
		if secX == 0 && secY == 3 {
			newSecX, newSecY := 1, 2
			x, y = y, x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 0, -1
			return x, y, dirX, dirY
		}
	} else if p.dirX == 0 && p.dirY == 1 {
		if secX == 0 && secY == 3 {
			newSecX, newSecY := 1, 0
			x, y = y, (p.cubeSize-1)-x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := -1, 0
			return x, y, dirX, dirY
		}
		if secX == 2 && secY == 0 {
			newSecX, newSecY := 1, 1
			x, y = y, x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := -1, 0
			return x, y, dirX, dirY
		}
		if secX == 1 && secY == 2 {
			newSecX, newSecY := 0, 3
			x, y = y, x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := -1, 0
			return x, y, dirX, dirY
		}
	} else if p.dirX == -1 && p.dirY == 0 {
		if secX == 0 && secY == 2 {
			newSecX, newSecY := 1, 0
			x, y = x, (p.cubeSize-1)-y
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 1, 0
			return x, y, dirX, dirY
		}
		if secX == 1 && secY == 1 {
			newSecX, newSecY := 1, 0
			x, y = y, x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 0, 1
			return x, y, dirX, dirY
		}
		if secX == 1 && secY == 0 {
			newSecX, newSecY := 0, 2
			x, y = x, (p.cubeSize-1)-y
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 1, 0
			return x, y, dirX, dirY
		}
	} else if p.dirX == 0 && p.dirY == -1 {
		if secX == 1 && secY == 0 {
			newSecX, newSecY := 0, 3
			x, y = y, (p.cubeSize-1)-x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 1, 0
			return x, y, dirX, dirY
		}
		if secX == 2 && secY == 0 {
			newSecX, newSecY := 0, 3
			x, y = x, (p.cubeSize-1)-y
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 0, -1
			return x, y, dirX, dirY
		}
		if secX == 1 && secY == 1 {
			newSecX, newSecY := 0, 2
			x, y = y, x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 0, 1
			return x, y, dirX, dirY
		}
		if secX == 0 && secY == 2 {
			newSecX, newSecY := 1, 1
			x, y = y, x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 1, 0
			return x, y, dirX, dirY
		}
	}
	// debug2(p.cube)
	panic("no section mapping")
}

func (p *Player) getNextSection() (int, int, int, int) {
	secX, secY := p.x/p.cubeSize, p.y/p.cubeSize
	// get relative
	x := p.x - secX*p.cubeSize
	y := p.y - secY*p.cubeSize
	if p.dirX == 1 && p.dirY == 0 {
		if secX == 2 && secY == 0 {
			newSecX, newSecY := 3, 2
			x, y = x, (p.cubeSize-1)-y
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := -1, 0
			return x, y, dirX, dirY
		}
		if secX == 2 && secY == 1 {
			newSecX, newSecY := 3, 2
			x, y = (p.cubeSize-1)-y, (p.cubeSize-1)-x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 0, 1
			return x, y, dirX, dirY
		}
	} else if p.dirX == 0 && p.dirY == 1 {
		if secX == 2 && secY == 2 {
			newSecX, newSecY := 0, 1
			x, y = (p.cubeSize-1)-x, y
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 0, -1
			return x, y, dirX, dirY
		}
	} else if p.dirX == -1 && p.dirY == 0 {
	} else if p.dirX == 0 && p.dirY == -1 {
		if secX == 1 && secY == 1 {
			newSecX, newSecY := 2, 0
			x, y = y, x
			x += newSecX * p.cubeSize
			y += newSecY * p.cubeSize
			dirX, dirY := 1, 0
			return x, y, dirX, dirY
		}
	}
	// debug2(p.cube)
	fmt.Println("section x,y:", secX, secY)
	fmt.Println("relative x,y:", x, y)
	fmt.Println("dir x,y:", p.dirX, p.dirY)
	panic("no section mapping")
}

func debug2(cube map[int]map[int]int) {
	for i := 0; i < limitY; i++ {
		for j := 0; j < limitX+1; j++ {
			if j%globalPlayer.cubeSize == 0 {
				fmt.Print(" ")
			}
			if globalPlayer.y == i && globalPlayer.x == j {
				fmt.Print("P")
			} else {
				fmt.Print(cube[i][j])
			}
		}
		fmt.Println()
		if i%globalPlayer.cubeSize == globalPlayer.cubeSize-1 {
			fmt.Println()
		}
	}
	fmt.Println()
}
