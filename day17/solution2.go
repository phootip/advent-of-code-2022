package day17

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/phootip/advent-of-code-2022/utils"
)

// var mem map[int]map[int]int
var mem [][2]int
var mem2 []int
// 1_000_000_000_000
// 999,999,999,950 + 29 = 999,999,999,979
// 28,571,428,570 steps
// rock 35
func Sol2() (ans int) {
	dc = gg.NewContext(100, 100_000)
	fmt.Println("Starting Day17 Solution2...")
	raw := utils.ReadFile("./day17/input.txt")
	// raw := utils.ReadFile("./day17/example.txt")
	raw = raw[:len(raw)-1]

	ans = simulate2(&controller{raw[0], 0})
	// checkDupe(mem, 35)
	// maxDupe()
	// fmt.Println(mem2)
	// checkAns(35)
	// checkAns(1760)
	// (1_000_000_000_000-21) rocks / 35 rocks = 28_571_428_570 steps
	// 35 * 28_571_428_570
	// ans2 := ans2(30,51,35,53)
	// fmt.Println("value off by: ",1514285714288 - ans2)
	ans2(1480,2325,1760,2737)
	return ans
}

// 1555113636384 too low
// 1555113636386 too high
func ans2(rockStart int, heightStart int, rockStep int, heightStep int) int {
	step := (1_000_000_000_000 - rockStart) / rockStep
	step_left := (1_000_000_000_000 - rockStart) % rockStep
	// step_left := 1_000_000_000_000 - (step*rockStep+rockStart)
	// fmt.Println(len(mem2)-rockStep)
	// fmt.Println(len(mem2)-1)
	// fmt.Println(mem2[len(mem2)-rockStep:])
	value_left := mem2[len(mem2)-(rockStep-step_left)-1] - mem2[len(mem2)-rockStep-1]
	// value_left := mem2[len(mem2)-(rockStep-step_left)-1] - mem2[len(mem2)-rockStep-1]
	fmt.Println("step: ", step)
	fmt.Println("step_left: ", step_left)
	fmt.Println("value_left: ", value_left)
	ans := heightStart + step*heightStep + value_left
	// ans := step*heightStep + value_left
	fmt.Println(ans)
	return ans
}

func checkAns(step int) {

	for i := len(mem2) - 1; i > 0; i -= step {
		fmt.Println("rock: ",i+1)
		fmt.Println("height: ", mem2[i])
		if i > step {
			fmt.Println("rock step: ", step)
			fmt.Println("height step: ", mem2[i] - mem2[i-step])
		}
		fmt.Println()
	}
}

func maxDupe() int {
	for i := len(mem)/2; i > 0; i-- {
		if checkDupe(mem,i) {
			// return i
			fmt.Println(i)
			// fmt.Println(mem2[i])
		}
	}
	return -1
}

func simulate2(control *controller) int {
	// stage := map[]
	g := game{}
	g.initGame(control)
	ans := g.start2()
	return ans
	// spawner := g.spawner
	// fmt.Println(spawner.rocks)

}

func (g *game) start2() int {
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
				g.terminateRock2()
			}
			// g.debugStage()
			// fmt.Println(g.countLoop)
		}
		// if g.countRock == 6595 {
		if g.countRock == 5000 {
			return 9999-g.height
		}
	}
	panic("out of game loop")
}

func (g *game) terminateRock2() {
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
	temp := g.height
	g.height = utils.Min(g.height, r.position[1]+(4-r.height))
	// fmt.Println("height diff: ", g.height-temp)
	// fmt.Println("rockId: ", g.rock1.id)
	// fmt.Println("countRock: ", g.countRock)
	mem = append(mem,	[2]int{g.rock1.id, temp - g.height})
	mem2 = append(mem2, 9999-g.height)
}

func checkDupe(mem[][2]int, length int) bool {
	if len(mem) < length * 2 {
		return false
	}
	divider := len(mem) - length
	return utils.SliceEqual(mem[divider:],mem[divider-length:divider]) 
	// fmt.Println(mem[divider:])
	// fmt.Println(mem[divider-length:divider])
}
