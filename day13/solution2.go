package day13

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol2() int {
	fmt.Println("Starting Day13 Solution2...")
	raw := utils.ReadFile("./day13/input.txt")
	// raw := utils.ReadFile("./day13/example.txt")
	// raw := utils.ReadFile("./day13/example1.txt")
	raw = append(raw, "[[2]]", "[[6]]")
	ans := 1
	packets := rawToPackets(raw)
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == 1
	})
	for i, packet := range packets {
		if checkDivider(packet) {
			ans *= i+1
		}
	}
	return ans
}

func checkDivider(packet any) bool{
	slice := packet.([]any)
	if len(slice) == 1 {
		if reflect.TypeOf(slice[0]).Kind() == reflect.Slice && len(slice[0].([]any)) == 1 {
			theInt, ok := slice[0].([]any)[0].(int)
			if ok && (theInt == 2 || theInt == 6) {
				return true
			}
		}
	}
	return false
}

func rawToPackets(raw []string) []any {
	result := []any{}
	for _, line := range raw {
		if line == "" {
			continue
		}
		signal, _ := lineToSignal(line[1:])
		result = append(result, signal)
	}
	return result
}
