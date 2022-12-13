package day13

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day13 Solution1...")
	// raw := utils.ReadFile("./day13/input.txt")
	raw := utils.ReadFile("./day13/example.txt")
	// raw := utils.ReadFile("./day13/example1.txt")
	rawToParis(raw)
	return 0
}

func rawToParis(raw []string) {
	state := "first"
	data := pair[int]{}
	_ = data
	for _, line := range raw {
		if line == "" {
			continue
		}
		switch state {
		case "first":
			a, _ := lineToSignal(line[1:])
			fmt.Println(a)
			state = "second"
		case "second":
			a, _ := lineToSignal(line[1:])
			fmt.Println(a)
			fmt.Println()
			state = "first"
		}
	}
}

func lineToSignal(line string) (result []any, skipTo int) {
	temp := ""
	for i, r := range line {
		c := string(r)
		// fmt.Println("temp: ", temp)
		if skipTo != 0 {
			skipTo--
			// fmt.Println("skipping: ", i, string(r))
			continue
		}
		// fmt.Println("parsing: ", i, string(r))
		switch c {
		case "[":
			var subResult any
			subResult, skipTo = lineToSignal(line[i+1:])
			result = append(result, subResult)
		case "]":
			if temp != "" {
				result = append(result, utils.StringToInt(temp))
			}
			return result, i+1
		case ",":
			if temp == "" {
				continue
			}
			result = append(result, utils.StringToInt(temp))
			temp = ""
		default:
			temp += string(r)
		}
	}
	panic("why reach end of function?")
	return result, -1
}


type pair[T int | []T] struct {
	first  []T
	second []T
}
