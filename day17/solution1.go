package day17

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() (ans int) {
	fmt.Println("Starting Day17 Solution1...")
	// raw := utils.ReadFile("./day17/input.txt")
	raw := utils.ReadFile("./day17/example.txt")
	raw = raw[:len(raw)-1]

	fmt.Println(raw)
	simulate(&controller{raw[0], 0})
	return ans
}

func simulate(Control *controller) {
	// stage := map[]
	g := game{}
	g.initGame()
	g.start()
	// spawner := g.spawner
	// fmt.Println(spawner.rocks)

}

func (g *game) start() {
	r := g.spawner.spawnRock()
	fmt.Println(r)
	g.debugStage()
}

func (r *rockSpawner) spawnRock() rock{
	r.idx++
	if r.idx >= 5 {
		r.idx = 0
	}
	return r.rocks[r.idx]
}

func (g *game) initGame() {
	g.width = 7
	g.height = 9999
	g.stage = [10_000][7]int{}
	g.stage[9999] = [7]int{2,2,2,2,2,2,2}
	// for i:=0; i < 9999; i++ {
	// 	g.stage[i] = [7]string{".",".",".",".",".",".",".",}
	// }
	g.spawner = &rockSpawner{}
	g.spawner.initRock()
}

func (g *game) debugStage() {
	for i := g.height-7; i<10_000; i++ {
		fmt.Println(g.stage[i])
	}
}

func (spawner *rockSpawner) initRock() {
	spawner.idx = 4
	spawner.rocks = [5]rock{}
	rock0 := rock{body: [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
	}}
	rock1 := rock{body: [4][4]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
	}}
	rock2 := rock{body: [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{1, 1, 1, 0},
	}}
	rock3 := rock{body: [4][4]int{
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}}
	rock4 := rock{body: [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 0, 0},
	}}
	spawner.rocks[0] = rock0
	spawner.rocks[1] = rock1
	spawner.rocks[2] = rock2
	spawner.rocks[3] = rock3
	spawner.rocks[4] = rock4
	// spawner.rocks = append(spawner.rocks)

}

type game struct {
	stage   [10_000][7]int
	height int
	width   int
	spawner *rockSpawner
}

type rockSpawner struct {
	rocks [5]rock
	idx   int
}

type rock struct {
	body [4][4]int
}

type controller struct {
	command string
	idx     int
}
