package day16

import (
	"fmt"
	"strings"

	"github.com/phootip/advent-of-code-2022/utils"
)

var mem map[string]map[string]int

func Sol1() (ans int) {
	fmt.Println("Starting Day16 Solution1...")
	mem = make(map[string]map[string]int)
	raw := utils.ReadFile("./day16/input.txt")
	// raw := utils.ReadFile("./day16/example.txt")
	graph := rawToNodes(raw)
	// graph = nodesToGraph(graph)
	// graph = filterEmptyEdge(graph)
	// printG(graph)

	ans = bestPressure(graph)
	// fmt.Println(mem)
	return ans
}

func bestPressure(graph map[string]*Node) (ans int) {
	destination := filterHasFlow(graph)
	fmt.Println("destination: ", destination)
	ans = traverse(graph, destination, []string{}, "AA", 30, 0)
	return ans
}

func traverse(graph map[string]*Node, destination []string, path []string, nodeName string, time int, pressure int) int {
	// fmt.Println("Traversing: ", path)
	// fmt.Println("Destination: ", destination)
	// fmt.Println("current Node: ", nodeName)
	// fmt.Println("time left: ", time)
	// fmt.Println("Pressure: ", pressure)
	// fmt.Println()
	if time <= 0{
		// fmt.Println("terminate: ",pressure, path)
		return pressure
	}
	node := graph[nodeName]
	if node.flowRate != 0 && utils.Contains(destination, nodeName) {
		time -= 1
		pressure += node.flowRate * time
		destination = utils.Filter(destination, nodeName)
	}
	if len(destination) == 0 {
		return pressure
	}
	maxResult := 0
	for _, desName := range destination {
		desNode := graph[desName]
		cost := shortestCost(graph, node, desNode)
		// fmt.Println(desName, cost)
		result := traverse(graph, destination, append(path, nodeName), desName, time-cost, pressure)
		maxResult = utils.Max(maxResult, result)
	}
	return maxResult
}

func shortestCost(graph map[string]*Node, source *Node, des *Node) (cost int) {
	if mem[source.name][des.name] != 0 {
		// fmt.Println("already calculated")
		return mem[source.name][des.name]
	}
	queue := []*Node{source}
	queueCost := []int{0}
	visited := []*Node{}
	var node *Node

	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]
		cost, queueCost = queueCost[0], queueCost[1:]
		if node == des {
			if mem[source.name] == nil {
				mem[source.name] = make(map[string]int)
			}
			mem[source.name][des.name] = cost
			return cost
		}

		for _, n := range node.adj {
			if !contains(visited, n) {
				queue = append(queue, n)
				queueCost = append(queueCost, cost+1)
			}
		}
	}
	panic("can't find route")
	return -1
}

func filterHasFlow(graph map[string]*Node) (destination []string) {
	temp := []string{}
	for _, v := range graph {
		if v.flowRate > 0 {
			temp = append(temp, v.name)
		}
	}
	return temp
}

func filterEmptyEdge(graph map[string]*Node) map[string]*Node {
	temp := map[string]*Node{}
	for k, v := range graph {
		if len(v.edges) > 0 {
			temp[k] = v
		}
	}
	return temp
}

func nodesToGraph(nodes map[string]*Node) map[string]*Node {
	node := nodes["AA"]
	var cost int
	var parent *Node

	visited := []*Node{}
	queue := []*Node{node}
	queueCost := []int{0}
	queueParent := []*Node{node}
	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]
		cost, queueCost = queueCost[0], queueCost[1:]
		parent, queueParent = queueParent[0], queueParent[1:]
		visited = append(visited, node)
		if node.flowRate != 0 {
			parent.edges = append(parent.edges, Edge{node.name, cost})
			node.edges = append(node.edges, Edge{parent.name, cost})
		}

		for _, n := range node.adj {
			if !contains(visited, n) {
				queue = append(queue, n)
				queueCost = append(queueCost, cost+1)
				if node.flowRate != 0 {
					queueParent = append(queueParent, node)
				} else {
					queueParent = append(queueParent, parent)
				}
			}
		}
	}

	return nodes

}

func contains[T int | rune | string | *Node](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func rawToNodes(raw []string) map[string]*Node {
	raw = raw[0 : len(raw)-1]
	nodes := map[string]*Node{}
	for _, line := range raw {
		splited := strings.Split(line, "; ")
		name := splited[0][6:8]
		if nodes[name] == nil {
			nodes[name] = &Node{}
		}
		node := nodes[name]
		node.name = name
		node.flowRate = utils.StringToInt(splited[0][23:])
		nodes[name] = node
		adj := strings.Split(splited[1], ", ")
		adj[0] = adj[0][len(adj[0])-2:]
		for _, n := range adj {
			if nodes[n] == nil {
				nodes[n] = &Node{}
			}
			node.adj = append(node.adj, nodes[n])
		}
	}
	return nodes
}

func printG(graph map[string]*Node) {
	for _, n := range graph {
		fmt.Printf("%+v\n", n)
	}
}

type Node struct {
	name     string
	flowRate int
	adj      []*Node
	edges    []Edge
}

type Edge struct {
	des  string
	cost int
}
