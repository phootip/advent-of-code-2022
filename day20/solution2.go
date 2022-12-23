package day20

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol2() (ans int) {
	fmt.Println("Starting Day20 Solution2...")
	raw := utils.ReadFile("./day20/input.txt")
	// raw := utils.ReadFile("./day20/example.txt")
	// raw := utils.ReadFile("./day20/example2.txt")
	raw = raw[:len(raw)-1]
	root := parseRaw2(raw, 811589153)
	// root := parseRaw2(raw, 1)
	debug()
	for i := 0; i < 10; i++ {
		swapNodes2(root)
	}
	debug()
	ans = answer1()
	return ans
}

func swapNodes2(root *Node) {
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
		for num >= len(mem) || num <= -len(mem){
			num = num%len(mem) + num/len(mem)
		}
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

func parseRaw2(raw []string, key int) *Node {
	// phantom := Node{}
	for i, line := range raw {
		mem = append(mem, &Node{nil, nil, utils.StringToInt(line)*key})
		if i > 0 {
			mem[i-1].next = mem[i]
			mem[i].prev = mem[i-1]
		}
	}
	mem[len(mem)-1].next = mem[0]
	mem[0].prev = mem[len(mem)-1]
	return mem[0]
}
