package day20

import (
	"fmt"

	"github.com/phootip/advent-of-code-2022/utils"
)

// 1306 too low
func Sol1() (ans int) {
	fmt.Println("Starting Day20 Solution1...")
	raw := utils.ReadFile("./day20/input.txt")
	// raw := utils.ReadFile("./day20/example.txt")
	raw = raw[:len(raw)-1]
	data := parseRaw(raw)
	ans = answer1(data)
	return ans
}

func answer1(data []int) (ans int) {
	newData := make([]int, len(data))
	copy(newData, data)
	fmt.Println("newData: ", newData)
	// swap
	for _, num := range data {
		idx := utils.SliceIndex(newData, num)
		// fmt.Println("idx: ", idx)
		// fmt.Println("num: ", num)
		newData = swap(newData, idx, idx+num)
		fmt.Println()
	}
	idx0 := utils.SliceIndex(newData, 0)
	idx1 := (idx0+1000)%len(newData)
	idx2 := (idx0+2000)%len(newData)
	idx3 := (idx0+3000)%len(newData)
	fmt.Println("idx1:",idx1)
	fmt.Println("idx2:",idx2)
	fmt.Println("idx3:",idx3)
	ans1 := newData[idx1]
	ans2 := newData[idx2]
	ans3 := newData[idx3]
	return ans1+ans2+ans3
}

func swap(slice []int, idx1 int, idx2 int) []int {
	fmt.Println("idx1&2: ",idx1, idx2)
	idx1 = reduceIndex(len(slice), idx1)
	idx2 = reduceIndex(len(slice), idx2)
	fmt.Println("idx1&2 after process: ",idx1, idx2)
	num := slice[idx1]
	fmt.Println("Moving num:", num)
	result := append(slice[:idx1], slice[idx1+1:]...)
	result = append(result[:idx2+1], result[idx2:]...)
	result[idx2] = num
	return result
}

func reduceIndex(length int, idx int)  int {
	for idx >= length || idx <= -length{
		idx = idx%length+idx/length
	}
	if idx < 0 {
		idx = length-1 + idx
	}
	return idx
}

func parseRaw(raw []string) (data []int) {
	for _, line := range raw {
		num := utils.StringToInt(line)
		data = append(data, num)
	}
	return data
}
