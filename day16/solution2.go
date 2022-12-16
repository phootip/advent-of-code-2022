package day16

import (
	"fmt"
	"strings"

	"github.com/mxschmitt/golang-combinations"
	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol2() (ans int) {
	fmt.Println("Starting Day16 Solution2...")
	mem = make(map[string]map[string]int)
	raw := utils.ReadFile("./day16/input.txt")
	// raw := utils.ReadFile("./day16/example.txt")
	graph := rawToNodes(raw)
	ans = bestPressure2(graph)
	// fmt.Println(mem)
	return ans
}

func bestPressure2(graph map[string]*Node) (ans int) {
	calculated := []string{}
	destination := filterHasFlow(graph)
	fmt.Println("destination: ",len(destination), destination)
	comb := combinations.Combinations(destination, len(destination)/2)
	for _, slice1 := range comb {
		if utils.Contains(calculated, strings.Join(slice1, ",")) {
			continue
		}
		slice2 := utils.Difference(destination, slice1)
		calculated = append(calculated, strings.Join(slice1, ","), strings.Join(slice2, ","))
		// fmt.Println(slice1, slice2)
		result := traverse(graph, slice1, []string{}, "AA", 26, 0) + traverse(graph, slice2, []string{}, "AA", 26, 0)
		ans = utils.Max(result,ans)
	}
	return ans
}
