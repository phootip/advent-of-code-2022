package day20

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

var mem = []*Node{}

// 1306 too low
// -3277 value duped brah
func Sol1() (ans int) {
	fmt.Println("Starting Day20 Solution1...")
	raw := utils.ReadFile("./day20/input.txt")
	// raw := utils.ReadFile("./day20/example.txt")
	raw = raw[:len(raw)-1]
	root := parseRaw(raw)
	swapNodes(root)
	ans = answer1()
	return ans
}

func answer1() (ans int) {
	root := findNodeWithVal(0)
	// debug()
	for i := 0; i <= 3000; i++ {
		if i == 1000 || i == 2000 || i == 3000 {
			fmt.Println(root.value)
			ans += root.value
		}
		root = root.next
	}
	return ans
}

func swapNodes(root *Node) {
	// debug()
	// swap node
	for _, n := range mem {
		node := n
		num := n.value
		// fmt.Println("Moving num:", num)
		movingNode := node
		node.prev.next = node.next
		node.next.prev = node.prev
		node = node.next
		// fmt.Println("Node popped")
		// debug()
		for num > 0 {
			node = node.next
			num--
		}
		for num < 0 {
			num++
			node = node.prev
		}
		movingNode.next = node
		movingNode.prev = node.prev
		movingNode.prev.next = movingNode
		movingNode.next.prev = movingNode
		// fmt.Println("Node added")
		// debug()
		// fmt.Println()
	}
}

func parseRaw(raw []string) *Node {
	// phantom := Node{}
	for i, line := range raw {
		mem = append(mem, &Node{nil, nil, utils.StringToInt(line)})
		if i > 0 {
			mem[i-1].next = mem[i]
			mem[i].prev = mem[i-1]
		}
	}
	mem[len(mem)-1].next = mem[0]
	mem[0].prev = mem[len(mem)-1]
	return mem[0]
}

func debug() {
	root := findNodeWithVal(0)
	for i := range mem {
		_ = i
		fmt.Print(root.value, ",")
		root = root.next
	}
	fmt.Println()
	// for i := range mem {
	// 	_ = i
	// 	// fmt.Print(root.value,",")
	// 	fmt.Println(root.prev.value, root.value, root.next.value)
	// 	root = root.next
	// }
	fmt.Println()
}

// func findNode(node *Node) *Node {

// }

func findNodeWithVal(value int) *Node {
	for _, node := range mem {
		if node.value == value {
			return node
		}
	}
	return nil
}

type Node struct {
	next  *Node
	prev  *Node
	value int
}
