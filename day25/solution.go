package day25

import (
	"fmt"
	"math"

	"github.com/phootip/advent-of-code-2022/utils"
)

func Sol1() (ans string) {
	fmt.Println("Starting Day25 Solution1...")
	raw := utils.ReadFile("./day25/input.txt")
	// raw := utils.ReadFile("./day25/example.txt")
	raw = raw[:len(raw)-1]
	ans = snafuToDeci(raw)

	return ans
}

func snafuToDeci(raw []string) string {
	result := []int{}
	// ans := 0
	sum := 0
	for i, num := range raw {
		// fmt.Println("processing:", num)
		result = append(result, snafuDecode(num))
		// fmt.Println("decoded:", result[i])
		sum += result[i]
	}
	fmt.Println(sum)
	return snafuEncode(sum)
}

func snafuEncode(num int) (result string) {
	// 1 digit, -2,2
	// 2 digit, -22,22
	// 3 digit, -67,67
	digitCount := 0
	limit := 0
	intToString := map[int]string{-2: "=", -1: "-", 0: "0", 1: "1", 2: "2"}
	for !(-limit < num && num < limit) {
		digitCount++
		limit = int(math.Pow(5, float64(digitCount))) / 2
	}
	for i := digitCount - 1; i >= 0; i-- {
		// find closest
		valLeft := 999999999999999999
		valAtDigit := -3
		for _, val := range []int{-2,-1,0, 1, 2} {
			DeciValue := val * int(math.Pow(5, float64(i)))
			if utils.Abs(num-DeciValue) < valLeft {
				valLeft = num - DeciValue
				valAtDigit = val
			}
		}
		result += intToString[valAtDigit]
		num = valLeft
		// fmt.Println("valLeft:", valLeft)
		// fmt.Println("valAtDigit:", valAtDigit)
	}
	return result
}

func snafuDecode(sna string) (result int) {
	l := len(sna)
	for i := l - 1; i >= 0; i-- {
		digit := l - i - 1
		num := snafuDigitToInt(string(sna[i]))
		result += num * int(math.Pow(5, float64(digit)))
	}
	return result
}

func snafuDigitToInt(num string) int {
	switch num {
	case "2":
		return 2
	case "1":
		return 1
	case "0":
		return 0
	case "-":
		return -1
	case "=":
		return -2
	}
	panic("Unknow snafu digit")

}

func parseRaw(raw []string) {

}
