package day24

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() (ans int) {
	fmt.Println("Starting Day23 Solution1...")
	// raw := utils.ReadFile("./day24/input.txt")
	raw := utils.ReadFile("./day24/example.txt")
	raw = raw[:len(raw)-1]
	game := parseRaw(raw)
	// game.debug()
	fmt.Println("generating rounds...")
	game.saveRound()
	game.shortestPath()
	_ = game
	return ans
}

func (g *Game) shortestPath() (ans int) {
	root := State{x: 1, y: 0, steps: 0}
	root.Fscore(g)
	heap := []State{root}
	for len(heap) > 0 {
		// get lowest score
		node := heap[0]
		heap = heap[1:]
		node.steps++
		graph := g.rounds[node.steps%len(g.rounds)]
		fmt.Println(graph)

		break
	}
	return ans
}

type State struct {
	x, y  int
	steps int
	score int
}

func (g *Game) saveRound() {
	i := 0
	for {
		points := []Point{}
		for _, obj := range g.objects {
			if obj.class == "blizzard" || obj.class == "wall" {
				point := Point{obj.x, obj.y}
				points = append(points, point)
			}
		}
		g.rounds = append(g.rounds, points)
		if i > 0 && g.roundEqual(0, i) {
			fmt.Println("round equal:", i)
			g.rounds = g.rounds[:len(g.rounds)-1]
			return
		}
		g.updateBliz()
		i++
	}
}

func (g *Game) roundEqual(a int, b int) bool {
	roundA := g.rounds[a]
	roundB := g.rounds[b]
	for i := range roundA {
		if roundA[i] != roundB[i] {
			return false
		}
	}
	return true
}

type Point struct {
	x, y int
}

func (s State) Fscore(g *Game) int {
	s.score = s.steps + (g.limitX - 1 - s.x) + (g.limitY - s.y)
	return s.score
}

func (g *Game) updateBliz() {
	for _, obj := range g.objects {
		if obj.class != "blizzard" {
			continue
		}
		g.move(obj)
	}
}

func (g *Game) move(obj *Object) {
	obj.x += obj.dirX
	obj.y += obj.dirY
	if obj.x == g.limitX {
		obj.x = 1
	} else if obj.x == 0 {
		obj.x = g.limitX - 1
	} else if obj.y == g.limitY {
		obj.y = 1
	} else if obj.y == 0 {
		obj.y = g.limitY - 1
	}
}

func parseRaw(raw []string) *Game {
	g := &Game{}
	g.limitX = len(raw[0]) - 1
	g.limitY = len(raw) - 1
	g.graph = [][]int{}
	g.rounds = [][]Point{}
	for i, line := range raw {
		for j, r := range line {
			switch string(r) {
			case "#":
				wall := &Object{x: j, y: i, class: "wall", icon: "#"}
				g.objects = append(g.objects, wall)
			case ">":
				bliz := &Object{x: j, y: i, class: "blizzard", dirX: 1, dirY: 0, icon: ">"}
				g.objects = append(g.objects, bliz)
			case "<":
				bliz := &Object{x: j, y: i, class: "blizzard", dirX: -1, dirY: 0, icon: "<"}
				g.objects = append(g.objects, bliz)
			case "^":
				bliz := &Object{x: j, y: i, class: "blizzard", dirX: 0, dirY: -1, icon: "^"}
				g.objects = append(g.objects, bliz)
			case "v":
				bliz := &Object{x: j, y: i, class: "blizzard", dirX: 0, dirY: 1, icon: "v"}
				g.objects = append(g.objects, bliz)
			}
		}
	}
	g.objects = append(g.objects, &Object{x: 1, y: -1, class: "wall"})
	g.objects = append(g.objects, &Object{x: g.limitX - 1, y: g.limitY + 1, class: "wall"})
	g.objects = append(g.objects, &Object{x: 1, y: 0, class: "player", icon: "P"})
	g.objects = append(g.objects, &Object{x: g.limitX - 1, y: g.limitY, class: "goal", icon: "G"})
	return g

}

func (g *Game) getObjAt(x int, y int) (result []*Object) {
	for _, obj := range g.objects {
		if obj.x == x && obj.y == y {
			result = append(result, obj)
		}
	}
	return result
}

func (g *Game) debugG(graph []Point) {
	for i := 0; i <= g.limitY; i++ {
		for j := 0; j <= g.limitX; j++ {
			objs := g.getObjAt(j, i)
			if len(objs) == 0 {
				fmt.Print(".")
			} else if len(objs) > 1 {
				fmt.Print(len(objs))
			} else {
				fmt.Print(objs[0].icon)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Game) debug() {
	for i := 0; i <= g.limitY; i++ {
		for j := 0; j <= g.limitX; j++ {
			objs := g.getObjAt(j, i)
			if len(objs) == 0 {
				fmt.Print(".")
			} else if len(objs) > 1 {
				fmt.Print(len(objs))
			} else {
				fmt.Print(objs[0].icon)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type Game struct {
	graph          [][]int
	limitX, limitY int
	objects        []*Object
	rounds         [][]Point
}

type Object struct {
	x, y       int
	dirX, dirY int
	class      string
	icon       string
}
