package day13

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() int {
	fmt.Println("Starting Day13 Solution1...")
	raw := utils.ReadFile("./day13/input.txt")
	// raw := utils.ReadFile("./day13/example.txt")
	// raw := utils.ReadFile("./day13/example1.txt")
	pairs := rawToParis(raw)
	ans := 0
	for i, aPair := range pairs {
		if compare(aPair.first, aPair.second) == 1 {
			ans += i + 1
			// fmt.Println("ans: ", ans)
		}
		// fmt.Println()
	}
	return ans
}

func compare(a, b any) int {
	a = ifIntToSlice(a)
	b = ifIntToSlice(b)
	// fmt.Println(a,"---", b)
	for i := range a.([]any) {
		// fmt.Println("i: ", i)
		// fmt.Println(a.([]any)[i])
		// fmt.Println(b.([]any)[i])
		if i > len(b.([]any))-1 {
			return -1
		}
		cInt, cIsInt := a.([]any)[i].(int)
		dInt, dIsInt := b.([]any)[i].(int)
		// fmt.Println(cInt, dInt)
		if !(cIsInt && dIsInt) {
			subResult := compare(a.([]any)[i], b.([]any)[i])
			// fmt.Println("subResult: ", subResult)
			if subResult != 0 {
				return subResult
			}
		} else {
			if cInt < dInt {
				return 1
			}
			if cInt > dInt {
				return -1
			}
		}
	}
	if len(a.([]any)) < len(b.([]any)) {
		return 1
	}
	return 0
}

func ifIntToSlice(a any) any {
	var c any
	switch b := a.(type) {
	case int:
		c = []any{b}
	default:
		c = b
	}
	return c
}

func rawToParis(raw []string) []pair {
	state := "first"
	result := []pair{}
	aPair := pair{}
	for _, line := range raw {
		if line == "" {
			continue
		}
		switch state {
		case "first":
			aPair.first, _ = lineToSignal(line[1:])
			state = "second"
		case "second":
			aPair.second, _ = lineToSignal(line[1:])
			result = append(result, aPair)
			state = "first"
		}
	}
	return result
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
			return result, i + 1
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

type pair struct {
	first  []any
	second []any
}
