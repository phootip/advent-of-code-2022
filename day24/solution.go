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
	ans = game.shortestPath()
	return ans
}

func (g *Game) shortestPath() (ans int) {
	root := State{Point: Point{x: 1, y: 0}, steps: 0}
	goal := g.goal
	root.Fscore(g)
	heap := []State{root}
	for len(heap) > 0 {
		// get lowest score
		node := heap[0]
		heap = heap[1:]
		fmt.Println("node:",node)
		if node.x == goal.x && node.y == goal.y {
			return node.steps
		}
		node.steps++
		adjPoints := g.getAdjPoint(node)
		for _, point := range adjPoints {
			newState := State{Point: point, steps: node.steps}
			newState.Fscore(g)
			heap = insertHeap(heap, newState)
		}
		// fmt.Println(heap)
	}
	panic("can't reach goal")
}

// can be optimise with binary insert
func insertHeap(heap []State, state State) []State{
	for i := range heap {
		if heap[i].score >= state.score {
			newHeap := append(heap[:i], state)
			newHeap = append(newHeap, heap[i:]...)
			return newHeap
		}
	}
	return append(heap, state)
}

func (g *Game) getAdjPoint(state State) (result []Point) {
	graph := g.rounds[state.steps%len(g.rounds)]
	for _, dir := range []Point{{1, 0}, {0, 1}, {0, 0}, {-1, 0}, {0, -1}} {
		adjPoint := Point{dir.x + state.x, dir.y + state.y}
		if graph[adjPoint] {
			continue
		}
		result = append(result, adjPoint)
	}
	return result
}

type State struct {
	Point
	steps int
	score int
}

func (g *Game) saveRound() {
	i := 0
	for {
		points := map[Point]bool{}
		for _, obj := range g.objects {
			if obj.class == "blizzard" || obj.class == "wall" {
				point := Point{obj.x, obj.y}
				points[point] = true
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

func (s *State) Fscore(g *Game) int {
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
	g.goal = *g.objects[len(g.objects)-1]
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

func (g *Game) debugG(graph map[Point]bool) {
	for i := 0; i <= g.limitY; i++ {
		for j := 0; j <= g.limitX; j++ {
			if graph[Point{j, i}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
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
	rounds         []map[Point]bool
	goal           Object
}

type Object struct {
	x, y       int
	dirX, dirY int
	class      string
	icon       string
}
