package utils

import "strconv"

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}


func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Contains[T int | rune](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
			if a == e {
					return true
			}
	}
	return false
}

func ReverseS(s string) (result string) {
	for _, r := range s {
		result = string(r) + result
	}
	return result
}

func SToInt(ss []string) []int {
	result := make([]int, len(ss))
	for i, s := range ss {
		r, err := strconv.Atoi(s)
		Check(err)
		result[i] = r
	}
	return result
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}
