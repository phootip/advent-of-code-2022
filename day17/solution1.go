package day17

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/phootip/advent-of-code-2022/utils"
)

// var dc
var dc *gg.Context

func Sol1() (ans int) {
	dc = gg.NewContext(100, 100_000)
	// dc = gg.NewContext(10_000,10_000)
	// dc = gg.NewContext(1000, 1000)
	fmt.Println("Starting Day17 Solution1...")
	raw := utils.ReadFile("./day17/input.txt")
	// raw := utils.ReadFile("./day17/example.txt")
	raw = raw[:len(raw)-1]

	ans = simulate(&controller{raw[0], 0})
	return ans
}

func simulate(control *controller) int {
	// stage := map[]
	g := game{}
	g.initGame(control)
	ans := g.start()
	return ans
	// spawner := g.spawner
	// fmt.Println(spawner.rocks)

}

func (g *game) start() int {
	for true {
		r := g.spawner.spawnRock()
		g.addRock(r)
		// g.debugStage()

		for g.hasRock {
			g.countLoop++
			g.controlRock()
			// g.debugStage()
			if g.canMoveDown() {
				g.moveDown()
			} else {
				g.terminateRock()
			}
			// g.debugStage()
			// fmt.Println(g.countLoop)
		}
		if g.countRock == 2022 {
			return 9999-g.height
		}
	}
	panic("out of game loop")
}

func (g *game) controlRock() {
	command := g.control.getCommand()
	// fmt.Println(command)
	if command == ">" && g.canMoveRight() {
		g.rock1.position[0]++
	} else if command == "<" && g.canMoveLeft() {
		g.rock1.position[0]--
	}
}

func (g *game) canMoveRight() bool {
	r := g.rock1
	x,y := r.position[0], r.position[1]
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			if r.body[j][i] == 1 && (i+x+1 >= 7 || g.stage[j+y][i+x+1] != 0) {
				return false
			}
		}
	}
	return true
}
func (g *game) canMoveLeft() bool {
	r := g.rock1
	x,y := r.position[0], r.position[1]
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			if r.body[j][i] == 1 && (i+x-1 < 0 || g.stage[j+y][i+x-1] != 0) {
				return false
			}
		}
	}
	return true
}
func (g *game) canMoveDown() bool {
	r := g.rock1
	x, y := r.position[0], r.position[1]
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			if r.body[j][i] == 1 && g.stage[j+y+1][i+x] != 0 {
				return false
			}
		}
	}
	return true
}
func (control *controller) getCommand() string {
	command := string(control.command[control.idx])
	control.idx++
	if control.idx >= len(control.command) {
		control.idx = 0
	}
	return command
}

func (g *game) terminateRock() {
	r := g.rock1
	x, y := r.position[0], r.position[1]
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			if r.body[j][i] == 1 {
				g.stage[j+y][i+x] = 3
			}
		}
	}
	g.hasRock = false
	g.countRock++
	g.height = utils.Min(g.height,r.position[1]+(4-r.height))
}

func (g *game) moveDown() {
	g.rock1.position[1]++
}

// func (g *game) addRock(r rock) {
// 	for j := 3; j >= 0; j-- {
// 		for i := 0; i < 4; i++ {
// 			g.stage[g.height+j-7][i+2] = r.body[j][i]
// 		}
// 	}
// }

func (g *game) addRock(r rock) {
	r.position = [2]int{2, g.height - 7}
	g.rock1 = &r
	g.hasRock = true
}

func (r *rockSpawner) spawnRock() rock {
	r.idx++
	if r.idx >= 5 {
		r.idx = 0
	}
	return r.rocks[r.idx]
}

func (g *game) initGame(control *controller) {
	g.width = 7
	g.height = 9999
	g.stage = [10_000][7]int{}
	g.stage[9999] = [7]int{2, 2, 2, 2, 2, 2, 2}
	g.spawner = &rockSpawner{}
	g.hasRock = false
	g.spawner.initRock()
	g.control = control
}

func (spawner *rockSpawner) initRock() {
	spawner.idx = 4
	spawner.rocks = [5]rock{}
	rock0 := rock{body: [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
	}, height: 1}
	rock1 := rock{body: [4][4]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
	}, height: 3}
	rock2 := rock{body: [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{1, 1, 1, 0},
	}, height: 3}
	rock3 := rock{body: [4][4]int{
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}, height: 4}
	rock4 := rock{body: [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 0, 0},
	}, height: 2}
	spawner.rocks[0] = rock0
	spawner.rocks[1] = rock1
	spawner.rocks[2] = rock2
	spawner.rocks[3] = rock3
	spawner.rocks[4] = rock4
	// spawner.rocks = append(spawner.rocks)

}

type game struct {
	stage     [10_000][7]int
	height    int
	width     int
	spawner   *rockSpawner
	rock1     *rock
	hasRock   bool
	countRock int
	control   *controller
	countLoop int
}

type rockSpawner struct {
	rocks [5]rock
	idx   int
}

type rock struct {
	body     [4][4]int
	height   int
	position [2]int
}

type controller struct {
	command string
	idx     int
}
